/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CredentialsInitParameters struct {

	// based authentication.
	// The type of authentication used for the Helm registry. Supported values are 'noAuth' for no authentication, 'basic' for username/password, and 'token' for token-based authentication.
	CredentialType *string `json:"credentialType,omitempty" tf:"credential_type,omitempty"`

	// (String) The password for basic authentication. Required if 'credential_type' is set to 'basic'.
	// The password for basic authentication. Required if 'credential_type' is set to 'basic'.
	Password *string `json:"password,omitempty" tf:"password,omitempty"`

	// (String) The authentication token. Required if 'credential_type' is set to 'token'.
	// The authentication token. Required if 'credential_type' is set to 'token'.
	Token *string `json:"token,omitempty" tf:"token,omitempty"`

	// (String) The username for basic authentication. Required if 'credential_type' is set to 'basic'.
	// The username for basic authentication. Required if 'credential_type' is set to 'basic'.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type CredentialsObservation struct {

	// based authentication.
	// The type of authentication used for the Helm registry. Supported values are 'noAuth' for no authentication, 'basic' for username/password, and 'token' for token-based authentication.
	CredentialType *string `json:"credentialType,omitempty" tf:"credential_type,omitempty"`

	// (String) The password for basic authentication. Required if 'credential_type' is set to 'basic'.
	// The password for basic authentication. Required if 'credential_type' is set to 'basic'.
	Password *string `json:"password,omitempty" tf:"password,omitempty"`

	// (String) The authentication token. Required if 'credential_type' is set to 'token'.
	// The authentication token. Required if 'credential_type' is set to 'token'.
	Token *string `json:"token,omitempty" tf:"token,omitempty"`

	// (String) The username for basic authentication. Required if 'credential_type' is set to 'basic'.
	// The username for basic authentication. Required if 'credential_type' is set to 'basic'.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type CredentialsParameters struct {

	// based authentication.
	// The type of authentication used for the Helm registry. Supported values are 'noAuth' for no authentication, 'basic' for username/password, and 'token' for token-based authentication.
	// +kubebuilder:validation:Optional
	CredentialType *string `json:"credentialType" tf:"credential_type,omitempty"`

	// (String) The password for basic authentication. Required if 'credential_type' is set to 'basic'.
	// The password for basic authentication. Required if 'credential_type' is set to 'basic'.
	// +kubebuilder:validation:Optional
	Password *string `json:"password,omitempty" tf:"password,omitempty"`

	// (String) The authentication token. Required if 'credential_type' is set to 'token'.
	// The authentication token. Required if 'credential_type' is set to 'token'.
	// +kubebuilder:validation:Optional
	Token *string `json:"token,omitempty" tf:"token,omitempty"`

	// (String) The username for basic authentication. Required if 'credential_type' is set to 'basic'.
	// The username for basic authentication. Required if 'credential_type' is set to 'basic'.
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type HelmInitParameters struct {

	// (Block List, Min: 1, Max: 1) Authentication credentials for accessing the Helm registry. (see below for nested schema)
	// Authentication credentials for accessing the Helm registry.
	Credentials []CredentialsInitParameters `json:"credentials,omitempty" tf:"credentials,omitempty"`

	// (String) The URL endpoint of the Helm registry where the charts are hosted.
	// The URL endpoint of the Helm registry where the charts are hosted.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// (Boolean) Specifies whether the Helm registry is private or public.
	// Specifies whether the Helm registry is private or public.
	IsPrivate *bool `json:"isPrivate,omitempty" tf:"is_private,omitempty"`

	// (String) The name of the Helm registry. This must be unique
	// The name of the Helm registry. This must be unique
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type HelmObservation struct {

	// (Block List, Min: 1, Max: 1) Authentication credentials for accessing the Helm registry. (see below for nested schema)
	// Authentication credentials for accessing the Helm registry.
	Credentials []CredentialsObservation `json:"credentials,omitempty" tf:"credentials,omitempty"`

	// (String) The URL endpoint of the Helm registry where the charts are hosted.
	// The URL endpoint of the Helm registry where the charts are hosted.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Boolean) Specifies whether the Helm registry is private or public.
	// Specifies whether the Helm registry is private or public.
	IsPrivate *bool `json:"isPrivate,omitempty" tf:"is_private,omitempty"`

	// (String) The name of the Helm registry. This must be unique
	// The name of the Helm registry. This must be unique
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type HelmParameters struct {

	// (Block List, Min: 1, Max: 1) Authentication credentials for accessing the Helm registry. (see below for nested schema)
	// Authentication credentials for accessing the Helm registry.
	// +kubebuilder:validation:Optional
	Credentials []CredentialsParameters `json:"credentials,omitempty" tf:"credentials,omitempty"`

	// (String) The URL endpoint of the Helm registry where the charts are hosted.
	// The URL endpoint of the Helm registry where the charts are hosted.
	// +kubebuilder:validation:Optional
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// (Boolean) Specifies whether the Helm registry is private or public.
	// Specifies whether the Helm registry is private or public.
	// +kubebuilder:validation:Optional
	IsPrivate *bool `json:"isPrivate,omitempty" tf:"is_private,omitempty"`

	// (String) The name of the Helm registry. This must be unique
	// The name of the Helm registry. This must be unique
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// HelmSpec defines the desired state of Helm
type HelmSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HelmParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider HelmInitParameters `json:"initProvider,omitempty"`
}

// HelmStatus defines the observed state of Helm.
type HelmStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HelmObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Helm is the Schema for the Helms API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,palette}
type Helm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.credentials) || (has(self.initProvider) && has(self.initProvider.credentials))",message="spec.forProvider.credentials is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.endpoint) || (has(self.initProvider) && has(self.initProvider.endpoint))",message="spec.forProvider.endpoint is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.isPrivate) || (has(self.initProvider) && has(self.initProvider.isPrivate))",message="spec.forProvider.isPrivate is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   HelmSpec   `json:"spec"`
	Status HelmStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HelmList contains a list of Helms
type HelmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Helm `json:"items"`
}

// Repository type metadata.
var (
	Helm_Kind             = "Helm"
	Helm_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Helm_Kind}.String()
	Helm_KindAPIVersion   = Helm_Kind + "." + CRDGroupVersion.String()
	Helm_GroupVersionKind = CRDGroupVersion.WithKind(Helm_Kind)
)

func init() {
	SchemeBuilder.Register(&Helm{}, &HelmList{})
}
