//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Step) DeepCopyInto(out *Step) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Step.
func (in *Step) DeepCopy() *Step {
	if in == nil {
		return nil
	}
	out := new(Step)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VeleroExport) DeepCopyInto(out *VeleroExport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VeleroExport.
func (in *VeleroExport) DeepCopy() *VeleroExport {
	if in == nil {
		return nil
	}
	out := new(VeleroExport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VeleroExport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VeleroExportList) DeepCopyInto(out *VeleroExportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VeleroExport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VeleroExportList.
func (in *VeleroExportList) DeepCopy() *VeleroExportList {
	if in == nil {
		return nil
	}
	out := new(VeleroExportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VeleroExportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VeleroExportSpec) DeepCopyInto(out *VeleroExportSpec) {
	*out = *in
	if in.VeleroBackupRef != nil {
		in, out := &in.VeleroBackupRef, &out.VeleroBackupRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.DataSourceMapping != nil {
		in, out := &in.DataSourceMapping, &out.DataSourceMapping
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VeleroExportSpec.
func (in *VeleroExportSpec) DeepCopy() *VeleroExportSpec {
	if in == nil {
		return nil
	}
	out := new(VeleroExportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VeleroExportStatus) DeepCopyInto(out *VeleroExportStatus) {
	*out = *in
	if in.VeleroBackupRef != nil {
		in, out := &in.VeleroBackupRef, &out.VeleroBackupRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.StartTimestamp != nil {
		in, out := &in.StartTimestamp, &out.StartTimestamp
		*out = (*in).DeepCopy()
	}
	if in.CompletionTimestamp != nil {
		in, out := &in.CompletionTimestamp, &out.CompletionTimestamp
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VeleroExportStatus.
func (in *VeleroExportStatus) DeepCopy() *VeleroExportStatus {
	if in == nil {
		return nil
	}
	out := new(VeleroExportStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VeleroImport) DeepCopyInto(out *VeleroImport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VeleroImport.
func (in *VeleroImport) DeepCopy() *VeleroImport {
	if in == nil {
		return nil
	}
	out := new(VeleroImport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VeleroImport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VeleroImportList) DeepCopyInto(out *VeleroImportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VeleroImport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VeleroImportList.
func (in *VeleroImportList) DeepCopy() *VeleroImportList {
	if in == nil {
		return nil
	}
	out := new(VeleroImportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VeleroImportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VeleroImportSpec) DeepCopyInto(out *VeleroImportSpec) {
	*out = *in
	if in.VeleroBackupRef != nil {
		in, out := &in.VeleroBackupRef, &out.VeleroBackupRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.PvcNames != nil {
		in, out := &in.PvcNames, &out.PvcNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VeleroImportSpec.
func (in *VeleroImportSpec) DeepCopy() *VeleroImportSpec {
	if in == nil {
		return nil
	}
	out := new(VeleroImportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VeleroImportStatus) DeepCopyInto(out *VeleroImportStatus) {
	*out = *in
	if in.VeleroRestoreRef != nil {
		in, out := &in.VeleroRestoreRef, &out.VeleroRestoreRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.StartTimestamp != nil {
		in, out := &in.StartTimestamp, &out.StartTimestamp
		*out = (*in).DeepCopy()
	}
	if in.CompletionTimestamp != nil {
		in, out := &in.CompletionTimestamp, &out.CompletionTimestamp
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VeleroImportStatus.
func (in *VeleroImportStatus) DeepCopy() *VeleroImportStatus {
	if in == nil {
		return nil
	}
	out := new(VeleroImportStatus)
	in.DeepCopyInto(out)
	return out
}
