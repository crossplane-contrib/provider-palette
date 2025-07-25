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

type ProjectInitParameters struct {

	// (String) The description of the project.
	// The description of the project.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The name of the project.
	// The name of the project.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Set of String) Assign tags to the project.
	// Assign tags to the project.
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ProjectObservation struct {

	// (String) The description of the project.
	// The description of the project.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The name of the project.
	// The name of the project.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Set of String) Assign tags to the project.
	// Assign tags to the project.
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ProjectParameters struct {

	// (String) The description of the project.
	// The description of the project.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The name of the project.
	// The name of the project.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Set of String) Assign tags to the project.
	// Assign tags to the project.
	// +kubebuilder:validation:Optional
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ProjectSpec defines the desired state of Project
type ProjectSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectParameters `json:"forProvider"`
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
	InitProvider ProjectInitParameters `json:"initProvider,omitempty"`
}

// ProjectStatus defines the observed state of Project.
type ProjectStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Project is the Schema for the Projects API. Create and manage projects in Palette.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,palette}
type Project struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ProjectSpec   `json:"spec"`
	Status ProjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectList contains a list of Projects
type ProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Project `json:"items"`
}

// Repository type metadata.
var (
	Project_Kind             = "Project"
	Project_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Project_Kind}.String()
	Project_KindAPIVersion   = Project_Kind + "." + CRDGroupVersion.String()
	Project_GroupVersionKind = CRDGroupVersion.WithKind(Project_Kind)
)

func init() {
	SchemeBuilder.Register(&Project{}, &ProjectList{})
}
