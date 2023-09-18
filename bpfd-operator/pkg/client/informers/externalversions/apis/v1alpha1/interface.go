/*
Copyright 2023 The bpfd Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/bpfd-dev/bpfd/bpfd-operator/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// BpfPrograms returns a BpfProgramInformer.
	BpfPrograms() BpfProgramInformer
	// KprobePrograms returns a KprobeProgramInformer.
	KprobePrograms() KprobeProgramInformer
	// TcPrograms returns a TcProgramInformer.
	TcPrograms() TcProgramInformer
	// TracepointPrograms returns a TracepointProgramInformer.
	TracepointPrograms() TracepointProgramInformer
	// UprobePrograms returns a UprobeProgramInformer.
	UprobePrograms() UprobeProgramInformer
	// XdpPrograms returns a XdpProgramInformer.
	XdpPrograms() XdpProgramInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// BpfPrograms returns a BpfProgramInformer.
func (v *version) BpfPrograms() BpfProgramInformer {
	return &bpfProgramInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// KprobePrograms returns a KprobeProgramInformer.
func (v *version) KprobePrograms() KprobeProgramInformer {
	return &kprobeProgramInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// TcPrograms returns a TcProgramInformer.
func (v *version) TcPrograms() TcProgramInformer {
	return &tcProgramInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// TracepointPrograms returns a TracepointProgramInformer.
func (v *version) TracepointPrograms() TracepointProgramInformer {
	return &tracepointProgramInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// UprobePrograms returns a UprobeProgramInformer.
func (v *version) UprobePrograms() UprobeProgramInformer {
	return &uprobeProgramInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// XdpPrograms returns a XdpProgramInformer.
func (v *version) XdpPrograms() XdpProgramInformer {
	return &xdpProgramInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
