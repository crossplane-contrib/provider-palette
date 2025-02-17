//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialsInitParameters) DeepCopyInto(out *CredentialsInitParameters) {
	*out = *in
	if in.CredentialType != nil {
		in, out := &in.CredentialType, &out.CredentialType
		*out = new(string)
		**out = **in
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(string)
		**out = **in
	}
	if in.Token != nil {
		in, out := &in.Token, &out.Token
		*out = new(string)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialsInitParameters.
func (in *CredentialsInitParameters) DeepCopy() *CredentialsInitParameters {
	if in == nil {
		return nil
	}
	out := new(CredentialsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialsObservation) DeepCopyInto(out *CredentialsObservation) {
	*out = *in
	if in.CredentialType != nil {
		in, out := &in.CredentialType, &out.CredentialType
		*out = new(string)
		**out = **in
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(string)
		**out = **in
	}
	if in.Token != nil {
		in, out := &in.Token, &out.Token
		*out = new(string)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialsObservation.
func (in *CredentialsObservation) DeepCopy() *CredentialsObservation {
	if in == nil {
		return nil
	}
	out := new(CredentialsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialsParameters) DeepCopyInto(out *CredentialsParameters) {
	*out = *in
	if in.CredentialType != nil {
		in, out := &in.CredentialType, &out.CredentialType
		*out = new(string)
		**out = **in
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(string)
		**out = **in
	}
	if in.Token != nil {
		in, out := &in.Token, &out.Token
		*out = new(string)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialsParameters.
func (in *CredentialsParameters) DeepCopy() *CredentialsParameters {
	if in == nil {
		return nil
	}
	out := new(CredentialsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Helm) DeepCopyInto(out *Helm) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Helm.
func (in *Helm) DeepCopy() *Helm {
	if in == nil {
		return nil
	}
	out := new(Helm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Helm) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmInitParameters) DeepCopyInto(out *HelmInitParameters) {
	*out = *in
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = make([]CredentialsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	if in.IsPrivate != nil {
		in, out := &in.IsPrivate, &out.IsPrivate
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmInitParameters.
func (in *HelmInitParameters) DeepCopy() *HelmInitParameters {
	if in == nil {
		return nil
	}
	out := new(HelmInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmList) DeepCopyInto(out *HelmList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Helm, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmList.
func (in *HelmList) DeepCopy() *HelmList {
	if in == nil {
		return nil
	}
	out := new(HelmList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HelmList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmObservation) DeepCopyInto(out *HelmObservation) {
	*out = *in
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = make([]CredentialsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IsPrivate != nil {
		in, out := &in.IsPrivate, &out.IsPrivate
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmObservation.
func (in *HelmObservation) DeepCopy() *HelmObservation {
	if in == nil {
		return nil
	}
	out := new(HelmObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmParameters) DeepCopyInto(out *HelmParameters) {
	*out = *in
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = make([]CredentialsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	if in.IsPrivate != nil {
		in, out := &in.IsPrivate, &out.IsPrivate
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmParameters.
func (in *HelmParameters) DeepCopy() *HelmParameters {
	if in == nil {
		return nil
	}
	out := new(HelmParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmSpec) DeepCopyInto(out *HelmSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmSpec.
func (in *HelmSpec) DeepCopy() *HelmSpec {
	if in == nil {
		return nil
	}
	out := new(HelmSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmStatus) DeepCopyInto(out *HelmStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmStatus.
func (in *HelmStatus) DeepCopy() *HelmStatus {
	if in == nil {
		return nil
	}
	out := new(HelmStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Oci) DeepCopyInto(out *Oci) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Oci.
func (in *Oci) DeepCopy() *Oci {
	if in == nil {
		return nil
	}
	out := new(Oci)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Oci) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OciCredentialsInitParameters) DeepCopyInto(out *OciCredentialsInitParameters) {
	*out = *in
	if in.AccessKey != nil {
		in, out := &in.AccessKey, &out.AccessKey
		*out = new(string)
		**out = **in
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.CredentialType != nil {
		in, out := &in.CredentialType, &out.CredentialType
		*out = new(string)
		**out = **in
	}
	if in.ExternalID != nil {
		in, out := &in.ExternalID, &out.ExternalID
		*out = new(string)
		**out = **in
	}
	if in.TLSConfig != nil {
		in, out := &in.TLSConfig, &out.TLSConfig
		*out = make([]TLSConfigInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OciCredentialsInitParameters.
func (in *OciCredentialsInitParameters) DeepCopy() *OciCredentialsInitParameters {
	if in == nil {
		return nil
	}
	out := new(OciCredentialsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OciCredentialsObservation) DeepCopyInto(out *OciCredentialsObservation) {
	*out = *in
	if in.AccessKey != nil {
		in, out := &in.AccessKey, &out.AccessKey
		*out = new(string)
		**out = **in
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.CredentialType != nil {
		in, out := &in.CredentialType, &out.CredentialType
		*out = new(string)
		**out = **in
	}
	if in.ExternalID != nil {
		in, out := &in.ExternalID, &out.ExternalID
		*out = new(string)
		**out = **in
	}
	if in.TLSConfig != nil {
		in, out := &in.TLSConfig, &out.TLSConfig
		*out = make([]TLSConfigObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OciCredentialsObservation.
func (in *OciCredentialsObservation) DeepCopy() *OciCredentialsObservation {
	if in == nil {
		return nil
	}
	out := new(OciCredentialsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OciCredentialsParameters) DeepCopyInto(out *OciCredentialsParameters) {
	*out = *in
	if in.AccessKey != nil {
		in, out := &in.AccessKey, &out.AccessKey
		*out = new(string)
		**out = **in
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.CredentialType != nil {
		in, out := &in.CredentialType, &out.CredentialType
		*out = new(string)
		**out = **in
	}
	if in.ExternalID != nil {
		in, out := &in.ExternalID, &out.ExternalID
		*out = new(string)
		**out = **in
	}
	if in.PasswordSecretRef != nil {
		in, out := &in.PasswordSecretRef, &out.PasswordSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.SecretKeySecretRef != nil {
		in, out := &in.SecretKeySecretRef, &out.SecretKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.TLSConfig != nil {
		in, out := &in.TLSConfig, &out.TLSConfig
		*out = make([]TLSConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OciCredentialsParameters.
func (in *OciCredentialsParameters) DeepCopy() *OciCredentialsParameters {
	if in == nil {
		return nil
	}
	out := new(OciCredentialsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OciInitParameters) DeepCopyInto(out *OciInitParameters) {
	*out = *in
	if in.BaseContentPath != nil {
		in, out := &in.BaseContentPath, &out.BaseContentPath
		*out = new(string)
		**out = **in
	}
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = make([]OciCredentialsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	if in.EndpointSuffix != nil {
		in, out := &in.EndpointSuffix, &out.EndpointSuffix
		*out = new(string)
		**out = **in
	}
	if in.IsPrivate != nil {
		in, out := &in.IsPrivate, &out.IsPrivate
		*out = new(bool)
		**out = **in
	}
	if in.IsSynchronization != nil {
		in, out := &in.IsSynchronization, &out.IsSynchronization
		*out = new(bool)
		**out = **in
	}
	if in.ProviderType != nil {
		in, out := &in.ProviderType, &out.ProviderType
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OciInitParameters.
func (in *OciInitParameters) DeepCopy() *OciInitParameters {
	if in == nil {
		return nil
	}
	out := new(OciInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OciList) DeepCopyInto(out *OciList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Oci, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OciList.
func (in *OciList) DeepCopy() *OciList {
	if in == nil {
		return nil
	}
	out := new(OciList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OciList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OciObservation) DeepCopyInto(out *OciObservation) {
	*out = *in
	if in.BaseContentPath != nil {
		in, out := &in.BaseContentPath, &out.BaseContentPath
		*out = new(string)
		**out = **in
	}
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = make([]OciCredentialsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	if in.EndpointSuffix != nil {
		in, out := &in.EndpointSuffix, &out.EndpointSuffix
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IsPrivate != nil {
		in, out := &in.IsPrivate, &out.IsPrivate
		*out = new(bool)
		**out = **in
	}
	if in.IsSynchronization != nil {
		in, out := &in.IsSynchronization, &out.IsSynchronization
		*out = new(bool)
		**out = **in
	}
	if in.ProviderType != nil {
		in, out := &in.ProviderType, &out.ProviderType
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OciObservation.
func (in *OciObservation) DeepCopy() *OciObservation {
	if in == nil {
		return nil
	}
	out := new(OciObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OciParameters) DeepCopyInto(out *OciParameters) {
	*out = *in
	if in.BaseContentPath != nil {
		in, out := &in.BaseContentPath, &out.BaseContentPath
		*out = new(string)
		**out = **in
	}
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = make([]OciCredentialsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	if in.EndpointSuffix != nil {
		in, out := &in.EndpointSuffix, &out.EndpointSuffix
		*out = new(string)
		**out = **in
	}
	if in.IsPrivate != nil {
		in, out := &in.IsPrivate, &out.IsPrivate
		*out = new(bool)
		**out = **in
	}
	if in.IsSynchronization != nil {
		in, out := &in.IsSynchronization, &out.IsSynchronization
		*out = new(bool)
		**out = **in
	}
	if in.ProviderType != nil {
		in, out := &in.ProviderType, &out.ProviderType
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OciParameters.
func (in *OciParameters) DeepCopy() *OciParameters {
	if in == nil {
		return nil
	}
	out := new(OciParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OciSpec) DeepCopyInto(out *OciSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OciSpec.
func (in *OciSpec) DeepCopy() *OciSpec {
	if in == nil {
		return nil
	}
	out := new(OciSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OciStatus) DeepCopyInto(out *OciStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OciStatus.
func (in *OciStatus) DeepCopy() *OciStatus {
	if in == nil {
		return nil
	}
	out := new(OciStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSConfigInitParameters) DeepCopyInto(out *TLSConfigInitParameters) {
	*out = *in
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = new(string)
		**out = **in
	}
	if in.InsecureSkipVerify != nil {
		in, out := &in.InsecureSkipVerify, &out.InsecureSkipVerify
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSConfigInitParameters.
func (in *TLSConfigInitParameters) DeepCopy() *TLSConfigInitParameters {
	if in == nil {
		return nil
	}
	out := new(TLSConfigInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSConfigObservation) DeepCopyInto(out *TLSConfigObservation) {
	*out = *in
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = new(string)
		**out = **in
	}
	if in.InsecureSkipVerify != nil {
		in, out := &in.InsecureSkipVerify, &out.InsecureSkipVerify
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSConfigObservation.
func (in *TLSConfigObservation) DeepCopy() *TLSConfigObservation {
	if in == nil {
		return nil
	}
	out := new(TLSConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSConfigParameters) DeepCopyInto(out *TLSConfigParameters) {
	*out = *in
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = new(string)
		**out = **in
	}
	if in.InsecureSkipVerify != nil {
		in, out := &in.InsecureSkipVerify, &out.InsecureSkipVerify
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSConfigParameters.
func (in *TLSConfigParameters) DeepCopy() *TLSConfigParameters {
	if in == nil {
		return nil
	}
	out := new(TLSConfigParameters)
	in.DeepCopyInto(out)
	return out
}
