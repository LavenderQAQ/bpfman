//go:build linux
// +build linux

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	bpfmanHelpers "github.com/bpfman/bpfman-operator/pkg/helpers"
	gobpfman "github.com/bpfman/bpfman/clients/gobpfman/v1"
	configMgmt "github.com/bpfman/bpfman/examples/pkg/config-mgmt"
	"github.com/cilium/ebpf"
)

const (
	UretprobeProgramName = "go-uretprobe-counter-example"
	BpfProgramMapIndex   = "uretprobe_stats_map"
	DefaultByteCodeFile  = "bpf_x86_bpfel.o"

	// MapsMountPoint is the "go-uretprobe-counter-maps" volumeMount "mountPath" from "deployment.yaml"
	MapsMountPoint = "/run/uretprobe/maps"
)

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -cc clang -no-strip -cflags "-O2 -g -Wall" -target amd64,arm64,ppc64le,s390x bpf ./bpf/uretprobe_counter.c -- -I.:/usr/include/bpf:/usr/include/linux

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// pull the BPFMAN config management data to determine if we're running on a
	// system with BPFMAN available.
	paramData, err := configMgmt.ParseParamData(configMgmt.ProgTypeUprobe, DefaultByteCodeFile)
	if err != nil {
		log.Printf("error processing parameters: %v\n", err)
		return
	}

	// determine the path to the uprobe_stats_map, whether provided via CRD
	// or BPFMAN or otherwise.
	var mapPath string
	// If running in a Kubernetes deployment, the eBPF program is already loaded.
	// Only need the map path, which is at a known location in the pod using VolumeMounts
	// and the CSI Driver.
	if paramData.CrdFlag {
		// 3. Get access to our map
		mapPath = fmt.Sprintf("%s/%s", MapsMountPoint, BpfProgramMapIndex)
	} else { // if not on k8s, find the map path from the system
		ctx := context.Background()

		// connect to the BPFMAN server
		conn, err := configMgmt.CreateConnection(ctx)
		if err != nil {
			log.Printf("failed to create client connection: %v", err)
			return
		}

		c := gobpfman.NewBpfmanClient(conn)

		fnName := "malloc"

		// If the bytecode src is a Program ID, skip the loading and unloading of the bytecode.
		if paramData.BytecodeSrc != configMgmt.SrcProgId {
			var loadRequest *gobpfman.LoadRequest
			if paramData.MapOwnerId != 0 {
				mapOwnerId := uint32(paramData.MapOwnerId)
				loadRequest = &gobpfman.LoadRequest{
					Bytecode:    paramData.BytecodeSource,
					Name:        "uretprobe_counter",
					ProgramType: *bpfmanHelpers.Kprobe.Uint32(),
					Attach: &gobpfman.AttachInfo{
						Info: &gobpfman.AttachInfo_UprobeAttachInfo{
							UprobeAttachInfo: &gobpfman.UprobeAttachInfo{
								FnName:   &fnName,
								Target:   "libc",
								Retprobe: true,
							},
						},
					},
					MapOwnerId: &mapOwnerId,
				}
			} else {
				loadRequest = &gobpfman.LoadRequest{
					Bytecode:    paramData.BytecodeSource,
					Name:        "uretprobe_counter",
					ProgramType: *bpfmanHelpers.Kprobe.Uint32(),
					Attach: &gobpfman.AttachInfo{
						Info: &gobpfman.AttachInfo_UprobeAttachInfo{
							UprobeAttachInfo: &gobpfman.UprobeAttachInfo{
								FnName:   &fnName,
								Target:   "libc",
								Retprobe: true,
							},
						},
					},
				}
			}

			// 1. Load Program using bpfman
			var res *gobpfman.LoadResponse
			res, err = c.Load(ctx, loadRequest)
			if err != nil {
				conn.Close()
				log.Print(err)
				return
			}

			kernelInfo := res.GetKernelInfo()
			if kernelInfo != nil {
				paramData.ProgId = uint(kernelInfo.GetId())
			} else {
				conn.Close()
				log.Printf("kernelInfo not returned in LoadResponse")
				return
			}
			log.Printf("Program registered with id %d\n", paramData.ProgId)

			// 2. Set up defer to unload program when this is closed
			defer func(id uint) {
				log.Printf("unloading program: %d\n", id)
				_, err = c.Unload(ctx, &gobpfman.UnloadRequest{Id: uint32(id)})
				if err != nil {
					conn.Close()
					log.Print(err)
					return
				}
				conn.Close()
			}(paramData.ProgId)

			// 3. Get access to our map
			mapPath, err = configMgmt.CalcMapPinPath(res.GetInfo(), "uretprobe_stats_map")
			if err != nil {
				log.Print(err)
				return
			}
		} else {
			// 2. Set up defer to close connection
			defer func(id uint) {
				log.Printf("Closing Connection for Program: %d\n", id)
				conn.Close()
			}(paramData.ProgId)

			// 3. Get access to our map
			mapPath, err = configMgmt.RetrieveMapPinPath(ctx, c, paramData.ProgId, "uretprobe_stats_map")
			if err != nil {
				log.Print(err)
				return
			}
		}
	}

	// load the pinned stats map which is keeping count of uprobe hits
	opts := &ebpf.LoadPinOptions{
		ReadOnly:  false,
		WriteOnly: false,
		Flags:     0,
	}
	statsMap, err := ebpf.LoadPinnedMap(mapPath, opts)
	if err != nil {
		log.Printf("Failed to load pinned Map: %s\n", mapPath)
		log.Print(err)
		return
	}

	// retrieve and report on the number of times the uprobe is executed.
	index := uint32(0)
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for range ticker.C {
			var stats []uint64
			var totalCount uint64

			if err := statsMap.Lookup(&index, &stats); err != nil {
				log.Printf("map lookup failed: %v", err)
				return
			}

			for _, stat := range stats {
				totalCount += stat
			}

			log.Printf("Uretprobe count: %d\n", totalCount)
		}
	}()

	<-stop

	log.Printf("Exiting...\n")
}
