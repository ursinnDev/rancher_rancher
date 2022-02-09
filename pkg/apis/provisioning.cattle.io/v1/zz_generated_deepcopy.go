//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v1

import (
	rkecattleiov1 "github.com/rancher/rancher/pkg/apis/rke.cattle.io/v1"
	genericcondition "github.com/rancher/wrangler/pkg/genericcondition"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAPIConfig) DeepCopyInto(out *ClusterAPIConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAPIConfig.
func (in *ClusterAPIConfig) DeepCopy() *ClusterAPIConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterAPIConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	if in.ClusterAPIConfig != nil {
		in, out := &in.ClusterAPIConfig, &out.ClusterAPIConfig
		*out = new(ClusterAPIConfig)
		**out = **in
	}
	if in.RKEConfig != nil {
		in, out := &in.RKEConfig, &out.RKEConfig
		*out = new(RKEConfig)
		(*in).DeepCopyInto(*out)
	}
	out.LocalClusterAuthEndpoint = in.LocalClusterAuthEndpoint
	if in.AgentEnvVars != nil {
		in, out := &in.AgentEnvVars, &out.AgentEnvVars
		*out = make([]rkecattleiov1.EnvVar, len(*in))
		copy(*out, *in)
	}
	if in.EnableNetworkPolicy != nil {
		in, out := &in.EnableNetworkPolicy, &out.EnableNetworkPolicy
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]genericcondition.GenericCondition, len(*in))
		copy(*out, *in)
	}
	if in.ETCDSnapshots != nil {
		in, out := &in.ETCDSnapshots, &out.ETCDSnapshots
		*out = make([]rkecattleiov1.ETCDSnapshot, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImportedConfig) DeepCopyInto(out *ImportedConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImportedConfig.
func (in *ImportedConfig) DeepCopy() *ImportedConfig {
	if in == nil {
		return nil
	}
	out := new(ImportedConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEConfig) DeepCopyInto(out *RKEConfig) {
	*out = *in
	in.RKEClusterSpecCommon.DeepCopyInto(&out.RKEClusterSpecCommon)
	if in.ETCDSnapshotCreate != nil {
		in, out := &in.ETCDSnapshotCreate, &out.ETCDSnapshotCreate
		*out = new(rkecattleiov1.ETCDSnapshotCreate)
		(*in).DeepCopyInto(*out)
	}
	if in.ETCDSnapshotRestore != nil {
		in, out := &in.ETCDSnapshotRestore, &out.ETCDSnapshotRestore
		*out = new(rkecattleiov1.ETCDSnapshotRestore)
		(*in).DeepCopyInto(*out)
	}
	if in.RotateCertificates != nil {
		in, out := &in.RotateCertificates, &out.RotateCertificates
		*out = new(rkecattleiov1.RotateCertificates)
		(*in).DeepCopyInto(*out)
	}
	if in.MachinePools != nil {
		in, out := &in.MachinePools, &out.MachinePools
		*out = make([]RKEMachinePool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InfrastructureRef != nil {
		in, out := &in.InfrastructureRef, &out.InfrastructureRef
		*out = new(corev1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEConfig.
func (in *RKEConfig) DeepCopy() *RKEConfig {
	if in == nil {
		return nil
	}
	out := new(RKEConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEMachinePool) DeepCopyInto(out *RKEMachinePool) {
	*out = *in
	in.RKECommonNodeConfig.DeepCopyInto(&out.RKECommonNodeConfig)
	if in.NodeConfig != nil {
		in, out := &in.NodeConfig, &out.NodeConfig
		*out = new(corev1.ObjectReference)
		**out = **in
	}
	if in.Quantity != nil {
		in, out := &in.Quantity, &out.Quantity
		*out = new(int32)
		**out = **in
	}
	if in.RollingUpdate != nil {
		in, out := &in.RollingUpdate, &out.RollingUpdate
		*out = new(RKEMachinePoolRollingUpdate)
		(*in).DeepCopyInto(*out)
	}
	if in.MachineDeploymentLabels != nil {
		in, out := &in.MachineDeploymentLabels, &out.MachineDeploymentLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.MachineDeploymentAnnotations != nil {
		in, out := &in.MachineDeploymentAnnotations, &out.MachineDeploymentAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodeStartupTimeout != nil {
		in, out := &in.NodeStartupTimeout, &out.NodeStartupTimeout
		*out = new(metav1.Duration)
		**out = **in
	}
	if in.UnhealthyNodeTimeout != nil {
		in, out := &in.UnhealthyNodeTimeout, &out.UnhealthyNodeTimeout
		*out = new(metav1.Duration)
		**out = **in
	}
	if in.MaxUnhealthy != nil {
		in, out := &in.MaxUnhealthy, &out.MaxUnhealthy
		*out = new(intstr.IntOrString)
		**out = **in
	}
	if in.UnhealthyRange != nil {
		in, out := &in.UnhealthyRange, &out.UnhealthyRange
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEMachinePool.
func (in *RKEMachinePool) DeepCopy() *RKEMachinePool {
	if in == nil {
		return nil
	}
	out := new(RKEMachinePool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RKEMachinePoolRollingUpdate) DeepCopyInto(out *RKEMachinePoolRollingUpdate) {
	*out = *in
	if in.MaxUnavailable != nil {
		in, out := &in.MaxUnavailable, &out.MaxUnavailable
		*out = new(intstr.IntOrString)
		**out = **in
	}
	if in.MaxSurge != nil {
		in, out := &in.MaxSurge, &out.MaxSurge
		*out = new(intstr.IntOrString)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RKEMachinePoolRollingUpdate.
func (in *RKEMachinePoolRollingUpdate) DeepCopy() *RKEMachinePoolRollingUpdate {
	if in == nil {
		return nil
	}
	out := new(RKEMachinePoolRollingUpdate)
	in.DeepCopyInto(out)
	return out
}
