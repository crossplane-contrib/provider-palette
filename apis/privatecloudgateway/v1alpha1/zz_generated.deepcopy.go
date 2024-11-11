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
func (in *DNSMap) DeepCopyInto(out *DNSMap) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSMap.
func (in *DNSMap) DeepCopy() *DNSMap {
	if in == nil {
		return nil
	}
	out := new(DNSMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DNSMap) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSMapInitParameters) DeepCopyInto(out *DNSMapInitParameters) {
	*out = *in
	if in.DataCenter != nil {
		in, out := &in.DataCenter, &out.DataCenter
		*out = new(string)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(string)
		**out = **in
	}
	if in.PrivateCloudGatewayID != nil {
		in, out := &in.PrivateCloudGatewayID, &out.PrivateCloudGatewayID
		*out = new(string)
		**out = **in
	}
	if in.SearchDomainName != nil {
		in, out := &in.SearchDomainName, &out.SearchDomainName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSMapInitParameters.
func (in *DNSMapInitParameters) DeepCopy() *DNSMapInitParameters {
	if in == nil {
		return nil
	}
	out := new(DNSMapInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSMapList) DeepCopyInto(out *DNSMapList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DNSMap, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSMapList.
func (in *DNSMapList) DeepCopy() *DNSMapList {
	if in == nil {
		return nil
	}
	out := new(DNSMapList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DNSMapList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSMapObservation) DeepCopyInto(out *DNSMapObservation) {
	*out = *in
	if in.DataCenter != nil {
		in, out := &in.DataCenter, &out.DataCenter
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(string)
		**out = **in
	}
	if in.PrivateCloudGatewayID != nil {
		in, out := &in.PrivateCloudGatewayID, &out.PrivateCloudGatewayID
		*out = new(string)
		**out = **in
	}
	if in.SearchDomainName != nil {
		in, out := &in.SearchDomainName, &out.SearchDomainName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSMapObservation.
func (in *DNSMapObservation) DeepCopy() *DNSMapObservation {
	if in == nil {
		return nil
	}
	out := new(DNSMapObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSMapParameters) DeepCopyInto(out *DNSMapParameters) {
	*out = *in
	if in.DataCenter != nil {
		in, out := &in.DataCenter, &out.DataCenter
		*out = new(string)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(string)
		**out = **in
	}
	if in.PrivateCloudGatewayID != nil {
		in, out := &in.PrivateCloudGatewayID, &out.PrivateCloudGatewayID
		*out = new(string)
		**out = **in
	}
	if in.SearchDomainName != nil {
		in, out := &in.SearchDomainName, &out.SearchDomainName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSMapParameters.
func (in *DNSMapParameters) DeepCopy() *DNSMapParameters {
	if in == nil {
		return nil
	}
	out := new(DNSMapParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSMapSpec) DeepCopyInto(out *DNSMapSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSMapSpec.
func (in *DNSMapSpec) DeepCopy() *DNSMapSpec {
	if in == nil {
		return nil
	}
	out := new(DNSMapSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSMapStatus) DeepCopyInto(out *DNSMapStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSMapStatus.
func (in *DNSMapStatus) DeepCopy() *DNSMapStatus {
	if in == nil {
		return nil
	}
	out := new(DNSMapStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ippool) DeepCopyInto(out *Ippool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ippool.
func (in *Ippool) DeepCopy() *Ippool {
	if in == nil {
		return nil
	}
	out := new(Ippool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Ippool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IppoolInitParameters) DeepCopyInto(out *IppoolInitParameters) {
	*out = *in
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = new(string)
		**out = **in
	}
	if in.IPEndRange != nil {
		in, out := &in.IPEndRange, &out.IPEndRange
		*out = new(string)
		**out = **in
	}
	if in.IPStartRange != nil {
		in, out := &in.IPStartRange, &out.IPStartRange
		*out = new(string)
		**out = **in
	}
	if in.NameserverAddresses != nil {
		in, out := &in.NameserverAddresses, &out.NameserverAddresses
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.NameserverSearchSuffix != nil {
		in, out := &in.NameserverSearchSuffix, &out.NameserverSearchSuffix
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.NetworkType != nil {
		in, out := &in.NetworkType, &out.NetworkType
		*out = new(string)
		**out = **in
	}
	if in.Prefix != nil {
		in, out := &in.Prefix, &out.Prefix
		*out = new(float64)
		**out = **in
	}
	if in.PrivateCloudGatewayID != nil {
		in, out := &in.PrivateCloudGatewayID, &out.PrivateCloudGatewayID
		*out = new(string)
		**out = **in
	}
	if in.RestrictToSingleCluster != nil {
		in, out := &in.RestrictToSingleCluster, &out.RestrictToSingleCluster
		*out = new(bool)
		**out = **in
	}
	if in.SubnetCidr != nil {
		in, out := &in.SubnetCidr, &out.SubnetCidr
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IppoolInitParameters.
func (in *IppoolInitParameters) DeepCopy() *IppoolInitParameters {
	if in == nil {
		return nil
	}
	out := new(IppoolInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IppoolList) DeepCopyInto(out *IppoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Ippool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IppoolList.
func (in *IppoolList) DeepCopy() *IppoolList {
	if in == nil {
		return nil
	}
	out := new(IppoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IppoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IppoolObservation) DeepCopyInto(out *IppoolObservation) {
	*out = *in
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IPEndRange != nil {
		in, out := &in.IPEndRange, &out.IPEndRange
		*out = new(string)
		**out = **in
	}
	if in.IPStartRange != nil {
		in, out := &in.IPStartRange, &out.IPStartRange
		*out = new(string)
		**out = **in
	}
	if in.NameserverAddresses != nil {
		in, out := &in.NameserverAddresses, &out.NameserverAddresses
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.NameserverSearchSuffix != nil {
		in, out := &in.NameserverSearchSuffix, &out.NameserverSearchSuffix
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.NetworkType != nil {
		in, out := &in.NetworkType, &out.NetworkType
		*out = new(string)
		**out = **in
	}
	if in.Prefix != nil {
		in, out := &in.Prefix, &out.Prefix
		*out = new(float64)
		**out = **in
	}
	if in.PrivateCloudGatewayID != nil {
		in, out := &in.PrivateCloudGatewayID, &out.PrivateCloudGatewayID
		*out = new(string)
		**out = **in
	}
	if in.RestrictToSingleCluster != nil {
		in, out := &in.RestrictToSingleCluster, &out.RestrictToSingleCluster
		*out = new(bool)
		**out = **in
	}
	if in.SubnetCidr != nil {
		in, out := &in.SubnetCidr, &out.SubnetCidr
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IppoolObservation.
func (in *IppoolObservation) DeepCopy() *IppoolObservation {
	if in == nil {
		return nil
	}
	out := new(IppoolObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IppoolParameters) DeepCopyInto(out *IppoolParameters) {
	*out = *in
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = new(string)
		**out = **in
	}
	if in.IPEndRange != nil {
		in, out := &in.IPEndRange, &out.IPEndRange
		*out = new(string)
		**out = **in
	}
	if in.IPStartRange != nil {
		in, out := &in.IPStartRange, &out.IPStartRange
		*out = new(string)
		**out = **in
	}
	if in.NameserverAddresses != nil {
		in, out := &in.NameserverAddresses, &out.NameserverAddresses
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.NameserverSearchSuffix != nil {
		in, out := &in.NameserverSearchSuffix, &out.NameserverSearchSuffix
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.NetworkType != nil {
		in, out := &in.NetworkType, &out.NetworkType
		*out = new(string)
		**out = **in
	}
	if in.Prefix != nil {
		in, out := &in.Prefix, &out.Prefix
		*out = new(float64)
		**out = **in
	}
	if in.PrivateCloudGatewayID != nil {
		in, out := &in.PrivateCloudGatewayID, &out.PrivateCloudGatewayID
		*out = new(string)
		**out = **in
	}
	if in.RestrictToSingleCluster != nil {
		in, out := &in.RestrictToSingleCluster, &out.RestrictToSingleCluster
		*out = new(bool)
		**out = **in
	}
	if in.SubnetCidr != nil {
		in, out := &in.SubnetCidr, &out.SubnetCidr
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IppoolParameters.
func (in *IppoolParameters) DeepCopy() *IppoolParameters {
	if in == nil {
		return nil
	}
	out := new(IppoolParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IppoolSpec) DeepCopyInto(out *IppoolSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IppoolSpec.
func (in *IppoolSpec) DeepCopy() *IppoolSpec {
	if in == nil {
		return nil
	}
	out := new(IppoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IppoolStatus) DeepCopyInto(out *IppoolStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IppoolStatus.
func (in *IppoolStatus) DeepCopy() *IppoolStatus {
	if in == nil {
		return nil
	}
	out := new(IppoolStatus)
	in.DeepCopyInto(out)
	return out
}
