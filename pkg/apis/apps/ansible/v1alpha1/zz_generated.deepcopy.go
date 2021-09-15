// +build !ignore_autogenerated

// Copyright 2019 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"encoding/json"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnsibleJob) DeepCopyInto(out *AnsibleJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnsibleJob.
func (in *AnsibleJob) DeepCopy() *AnsibleJob {
	if in == nil {
		return nil
	}
	out := new(AnsibleJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnsibleJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnsibleJobList) DeepCopyInto(out *AnsibleJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AnsibleJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnsibleJobList.
func (in *AnsibleJobList) DeepCopy() *AnsibleJobList {
	if in == nil {
		return nil
	}
	out := new(AnsibleJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnsibleJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnsibleJobResult) DeepCopyInto(out *AnsibleJobResult) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnsibleJobResult.
func (in *AnsibleJobResult) DeepCopy() *AnsibleJobResult {
	if in == nil {
		return nil
	}
	out := new(AnsibleJobResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnsibleJobSpec) DeepCopyInto(out *AnsibleJobSpec) {
	*out = *in
	if in.ExtraVars != nil {
		in, out := &in.ExtraVars, &out.ExtraVars
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnsibleJobSpec.
func (in *AnsibleJobSpec) DeepCopy() *AnsibleJobSpec {
	if in == nil {
		return nil
	}
	out := new(AnsibleJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnsibleJobStatus) DeepCopyInto(out *AnsibleJobStatus) {
	*out = *in
	out.AnsibleJobResult = in.AnsibleJobResult
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.K8sJob = in.K8sJob
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnsibleJobStatus.
func (in *AnsibleJobStatus) DeepCopy() *AnsibleJobStatus {
	if in == nil {
		return nil
	}
	out := new(AnsibleJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnsibleResult) DeepCopyInto(out *AnsibleResult) {
	*out = *in
	in.TimeOfCompletion.DeepCopyInto(&out.TimeOfCompletion)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnsibleResult.
func (in *AnsibleResult) DeepCopy() *AnsibleResult {
	if in == nil {
		return nil
	}
	out := new(AnsibleResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	if in.AnsibleResult != nil {
		in, out := &in.AnsibleResult, &out.AnsibleResult
		*out = new(AnsibleResult)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Env) DeepCopyInto(out *Env) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Env.
func (in *Env) DeepCopy() *Env {
	if in == nil {
		return nil
	}
	out := new(Env)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventTime) DeepCopyInto(out *EventTime) {
	*out = *in
	in.Time.DeepCopyInto(&out.Time)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventTime.
func (in *EventTime) DeepCopy() *EventTime {
	if in == nil {
		return nil
	}
	out := new(EventTime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8sJob) DeepCopyInto(out *K8sJob) {
	*out = *in
	out.Env = in.Env
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8sJob.
func (in *K8sJob) DeepCopy() *K8sJob {
	if in == nil {
		return nil
	}
	out := new(K8sJob)
	in.DeepCopyInto(out)
	return out
}
