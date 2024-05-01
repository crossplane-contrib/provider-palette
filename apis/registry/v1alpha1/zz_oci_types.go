// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type OciCredentialsInitParameters struct {

	// (String)
	AccessKey *string `json:"accessKey,omitempty" tf:"access_key,omitempty"`

	// (String)
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// (String)
	CredentialType *string `json:"credentialType,omitempty" tf:"credential_type,omitempty"`

	// (String)
	ExternalID *string `json:"externalId,omitempty" tf:"external_id,omitempty"`
}

type OciCredentialsObservation struct {

	// (String)
	AccessKey *string `json:"accessKey,omitempty" tf:"access_key,omitempty"`

	// (String)
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// (String)
	CredentialType *string `json:"credentialType,omitempty" tf:"credential_type,omitempty"`

	// (String)
	ExternalID *string `json:"externalId,omitempty" tf:"external_id,omitempty"`
}

type OciCredentialsParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	AccessKey *string `json:"accessKey,omitempty" tf:"access_key,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	CredentialType *string `json:"credentialType" tf:"credential_type,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	ExternalID *string `json:"externalId,omitempty" tf:"external_id,omitempty"`

	// (String, Sensitive)
	// +kubebuilder:validation:Optional
	SecretKeySecretRef *v1.SecretKeySelector `json:"secretKeySecretRef,omitempty" tf:"-"`
}

type OciInitParameters struct {

	// (Block List, Min: 1, Max: 1) (see below for nested schema)
	Credentials []OciCredentialsInitParameters `json:"credentials,omitempty" tf:"credentials,omitempty"`

	// (String)
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// (Boolean)
	IsPrivate *bool `json:"isPrivate,omitempty" tf:"is_private,omitempty"`

	// (String)
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type OciObservation struct {

	// (Block List, Min: 1, Max: 1) (see below for nested schema)
	Credentials []OciCredentialsObservation `json:"credentials,omitempty" tf:"credentials,omitempty"`

	// (String)
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Boolean)
	IsPrivate *bool `json:"isPrivate,omitempty" tf:"is_private,omitempty"`

	// (String)
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type OciParameters struct {

	// (Block List, Min: 1, Max: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Credentials []OciCredentialsParameters `json:"credentials,omitempty" tf:"credentials,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// (Boolean)
	// +kubebuilder:validation:Optional
	IsPrivate *bool `json:"isPrivate,omitempty" tf:"is_private,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// OciSpec defines the desired state of Oci
type OciSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OciParameters `json:"forProvider"`
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
	InitProvider OciInitParameters `json:"initProvider,omitempty"`
}

// OciStatus defines the observed state of Oci.
type OciStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OciObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Oci is the Schema for the Ocis API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,palette}
type Oci struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.credentials) || (has(self.initProvider) && has(self.initProvider.credentials))",message="spec.forProvider.credentials is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.endpoint) || (has(self.initProvider) && has(self.initProvider.endpoint))",message="spec.forProvider.endpoint is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.isPrivate) || (has(self.initProvider) && has(self.initProvider.isPrivate))",message="spec.forProvider.isPrivate is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   OciSpec   `json:"spec"`
	Status OciStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OciList contains a list of Ocis
type OciList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Oci `json:"items"`
}

// Repository type metadata.
var (
	Oci_Kind             = "Oci"
	Oci_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Oci_Kind}.String()
	Oci_KindAPIVersion   = Oci_Kind + "." + CRDGroupVersion.String()
	Oci_GroupVersionKind = CRDGroupVersion.WithKind(Oci_Kind)
)

func init() {
	SchemeBuilder.Register(&Oci{}, &OciList{})
}
