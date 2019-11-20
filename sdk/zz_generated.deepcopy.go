// +build !ignore_autogenerated

/*
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

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComposableGetValueFrom) DeepCopyInto(out *ComposableGetValueFrom) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.FormatTransformers != nil {
		in, out := &in.FormatTransformers, &out.FormatTransformers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComposableGetValueFrom.
func (in *ComposableGetValueFrom) DeepCopy() *ComposableGetValueFrom {
	if in == nil {
		return nil
	}
	out := new(ComposableGetValueFrom)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComposableGetValueFrom) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GetValueFromType) DeepCopyInto(out *GetValueFromType) {
	*out = *in
	in.GetValueFrom.DeepCopyInto(&out.GetValueFrom)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GetValueFromType.
func (in *GetValueFromType) DeepCopy() *GetValueFromType {
	if in == nil {
		return nil
	}
	out := new(GetValueFromType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GetValueFromType) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
