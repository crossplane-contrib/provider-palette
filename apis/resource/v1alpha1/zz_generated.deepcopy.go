//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Limit) DeepCopyInto(out *Limit) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Limit.
func (in *Limit) DeepCopy() *Limit {
	if in == nil {
		return nil
	}
	out := new(Limit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Limit) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LimitInitParameters) DeepCopyInto(out *LimitInitParameters) {
	*out = *in
	if in.APIKeys != nil {
		in, out := &in.APIKeys, &out.APIKeys
		*out = new(float64)
		**out = **in
	}
	if in.Alert != nil {
		in, out := &in.Alert, &out.Alert
		*out = new(float64)
		**out = **in
	}
	if in.Appliance != nil {
		in, out := &in.Appliance, &out.Appliance
		*out = new(float64)
		**out = **in
	}
	if in.ApplianceToken != nil {
		in, out := &in.ApplianceToken, &out.ApplianceToken
		*out = new(float64)
		**out = **in
	}
	if in.ApplicationDeployment != nil {
		in, out := &in.ApplicationDeployment, &out.ApplicationDeployment
		*out = new(float64)
		**out = **in
	}
	if in.ApplicationProfile != nil {
		in, out := &in.ApplicationProfile, &out.ApplicationProfile
		*out = new(float64)
		**out = **in
	}
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = new(float64)
		**out = **in
	}
	if in.CloudAccount != nil {
		in, out := &in.CloudAccount, &out.CloudAccount
		*out = new(float64)
		**out = **in
	}
	if in.Cluster != nil {
		in, out := &in.Cluster, &out.Cluster
		*out = new(float64)
		**out = **in
	}
	if in.ClusterGroup != nil {
		in, out := &in.ClusterGroup, &out.ClusterGroup
		*out = new(float64)
		**out = **in
	}
	if in.ClusterProfile != nil {
		in, out := &in.ClusterProfile, &out.ClusterProfile
		*out = new(float64)
		**out = **in
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(float64)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(float64)
		**out = **in
	}
	if in.Macro != nil {
		in, out := &in.Macro, &out.Macro
		*out = new(float64)
		**out = **in
	}
	if in.PrivateGateway != nil {
		in, out := &in.PrivateGateway, &out.PrivateGateway
		*out = new(float64)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(float64)
		**out = **in
	}
	if in.Registry != nil {
		in, out := &in.Registry, &out.Registry
		*out = new(float64)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(float64)
		**out = **in
	}
	if in.SSHKey != nil {
		in, out := &in.SSHKey, &out.SSHKey
		*out = new(float64)
		**out = **in
	}
	if in.Team != nil {
		in, out := &in.Team, &out.Team
		*out = new(float64)
		**out = **in
	}
	if in.User != nil {
		in, out := &in.User, &out.User
		*out = new(float64)
		**out = **in
	}
	if in.Workspace != nil {
		in, out := &in.Workspace, &out.Workspace
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LimitInitParameters.
func (in *LimitInitParameters) DeepCopy() *LimitInitParameters {
	if in == nil {
		return nil
	}
	out := new(LimitInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LimitList) DeepCopyInto(out *LimitList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Limit, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LimitList.
func (in *LimitList) DeepCopy() *LimitList {
	if in == nil {
		return nil
	}
	out := new(LimitList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LimitList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LimitObservation) DeepCopyInto(out *LimitObservation) {
	*out = *in
	if in.APIKeys != nil {
		in, out := &in.APIKeys, &out.APIKeys
		*out = new(float64)
		**out = **in
	}
	if in.Alert != nil {
		in, out := &in.Alert, &out.Alert
		*out = new(float64)
		**out = **in
	}
	if in.Appliance != nil {
		in, out := &in.Appliance, &out.Appliance
		*out = new(float64)
		**out = **in
	}
	if in.ApplianceToken != nil {
		in, out := &in.ApplianceToken, &out.ApplianceToken
		*out = new(float64)
		**out = **in
	}
	if in.ApplicationDeployment != nil {
		in, out := &in.ApplicationDeployment, &out.ApplicationDeployment
		*out = new(float64)
		**out = **in
	}
	if in.ApplicationProfile != nil {
		in, out := &in.ApplicationProfile, &out.ApplicationProfile
		*out = new(float64)
		**out = **in
	}
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = new(float64)
		**out = **in
	}
	if in.CloudAccount != nil {
		in, out := &in.CloudAccount, &out.CloudAccount
		*out = new(float64)
		**out = **in
	}
	if in.Cluster != nil {
		in, out := &in.Cluster, &out.Cluster
		*out = new(float64)
		**out = **in
	}
	if in.ClusterGroup != nil {
		in, out := &in.ClusterGroup, &out.ClusterGroup
		*out = new(float64)
		**out = **in
	}
	if in.ClusterProfile != nil {
		in, out := &in.ClusterProfile, &out.ClusterProfile
		*out = new(float64)
		**out = **in
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(float64)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(float64)
		**out = **in
	}
	if in.Macro != nil {
		in, out := &in.Macro, &out.Macro
		*out = new(float64)
		**out = **in
	}
	if in.PrivateGateway != nil {
		in, out := &in.PrivateGateway, &out.PrivateGateway
		*out = new(float64)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(float64)
		**out = **in
	}
	if in.Registry != nil {
		in, out := &in.Registry, &out.Registry
		*out = new(float64)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(float64)
		**out = **in
	}
	if in.SSHKey != nil {
		in, out := &in.SSHKey, &out.SSHKey
		*out = new(float64)
		**out = **in
	}
	if in.Team != nil {
		in, out := &in.Team, &out.Team
		*out = new(float64)
		**out = **in
	}
	if in.User != nil {
		in, out := &in.User, &out.User
		*out = new(float64)
		**out = **in
	}
	if in.Workspace != nil {
		in, out := &in.Workspace, &out.Workspace
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LimitObservation.
func (in *LimitObservation) DeepCopy() *LimitObservation {
	if in == nil {
		return nil
	}
	out := new(LimitObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LimitParameters) DeepCopyInto(out *LimitParameters) {
	*out = *in
	if in.APIKeys != nil {
		in, out := &in.APIKeys, &out.APIKeys
		*out = new(float64)
		**out = **in
	}
	if in.Alert != nil {
		in, out := &in.Alert, &out.Alert
		*out = new(float64)
		**out = **in
	}
	if in.Appliance != nil {
		in, out := &in.Appliance, &out.Appliance
		*out = new(float64)
		**out = **in
	}
	if in.ApplianceToken != nil {
		in, out := &in.ApplianceToken, &out.ApplianceToken
		*out = new(float64)
		**out = **in
	}
	if in.ApplicationDeployment != nil {
		in, out := &in.ApplicationDeployment, &out.ApplicationDeployment
		*out = new(float64)
		**out = **in
	}
	if in.ApplicationProfile != nil {
		in, out := &in.ApplicationProfile, &out.ApplicationProfile
		*out = new(float64)
		**out = **in
	}
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = new(float64)
		**out = **in
	}
	if in.CloudAccount != nil {
		in, out := &in.CloudAccount, &out.CloudAccount
		*out = new(float64)
		**out = **in
	}
	if in.Cluster != nil {
		in, out := &in.Cluster, &out.Cluster
		*out = new(float64)
		**out = **in
	}
	if in.ClusterGroup != nil {
		in, out := &in.ClusterGroup, &out.ClusterGroup
		*out = new(float64)
		**out = **in
	}
	if in.ClusterProfile != nil {
		in, out := &in.ClusterProfile, &out.ClusterProfile
		*out = new(float64)
		**out = **in
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(float64)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(float64)
		**out = **in
	}
	if in.Macro != nil {
		in, out := &in.Macro, &out.Macro
		*out = new(float64)
		**out = **in
	}
	if in.PrivateGateway != nil {
		in, out := &in.PrivateGateway, &out.PrivateGateway
		*out = new(float64)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(float64)
		**out = **in
	}
	if in.Registry != nil {
		in, out := &in.Registry, &out.Registry
		*out = new(float64)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(float64)
		**out = **in
	}
	if in.SSHKey != nil {
		in, out := &in.SSHKey, &out.SSHKey
		*out = new(float64)
		**out = **in
	}
	if in.Team != nil {
		in, out := &in.Team, &out.Team
		*out = new(float64)
		**out = **in
	}
	if in.User != nil {
		in, out := &in.User, &out.User
		*out = new(float64)
		**out = **in
	}
	if in.Workspace != nil {
		in, out := &in.Workspace, &out.Workspace
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LimitParameters.
func (in *LimitParameters) DeepCopy() *LimitParameters {
	if in == nil {
		return nil
	}
	out := new(LimitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LimitSpec) DeepCopyInto(out *LimitSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LimitSpec.
func (in *LimitSpec) DeepCopy() *LimitSpec {
	if in == nil {
		return nil
	}
	out := new(LimitSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LimitStatus) DeepCopyInto(out *LimitStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LimitStatus.
func (in *LimitStatus) DeepCopy() *LimitStatus {
	if in == nil {
		return nil
	}
	out := new(LimitStatus)
	in.DeepCopyInto(out)
	return out
}
