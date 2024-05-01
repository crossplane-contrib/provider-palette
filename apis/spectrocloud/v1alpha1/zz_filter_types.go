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

type FilterGroupInitParameters struct {

	// (String) Conjunction operation of the filter group. Valid values are 'and' and 'or'.
	// Conjunction operation of the filter group. Valid values are 'and' and 'or'.
	Conjunction *string `json:"conjunction,omitempty" tf:"conjunction,omitempty"`

	// (Block List, Min: 1) List of filters in the filter group. (see below for nested schema)
	// List of filters in the filter group.
	Filters []FiltersInitParameters `json:"filters,omitempty" tf:"filters,omitempty"`
}

type FilterGroupObservation struct {

	// (String) Conjunction operation of the filter group. Valid values are 'and' and 'or'.
	// Conjunction operation of the filter group. Valid values are 'and' and 'or'.
	Conjunction *string `json:"conjunction,omitempty" tf:"conjunction,omitempty"`

	// (Block List, Min: 1) List of filters in the filter group. (see below for nested schema)
	// List of filters in the filter group.
	Filters []FiltersObservation `json:"filters,omitempty" tf:"filters,omitempty"`
}

type FilterGroupParameters struct {

	// (String) Conjunction operation of the filter group. Valid values are 'and' and 'or'.
	// Conjunction operation of the filter group. Valid values are 'and' and 'or'.
	// +kubebuilder:validation:Optional
	Conjunction *string `json:"conjunction" tf:"conjunction,omitempty"`

	// (Block List, Min: 1) List of filters in the filter group. (see below for nested schema)
	// List of filters in the filter group.
	// +kubebuilder:validation:Optional
	Filters []FiltersParameters `json:"filters" tf:"filters,omitempty"`
}

type FilterInitParameters struct {

	// (Block List, Min: 1, Max: 1) Metadata of the filter. (see below for nested schema)
	// Metadata of the filter.
	Metadata []FilterMetadataInitParameters `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// (Block List, Min: 1, Max: 1) Specification of the filter. (see below for nested schema)
	// Specification of the filter.
	Spec []FilterSpecInitParameters `json:"spec,omitempty" tf:"spec,omitempty"`
}

type FilterMetadataInitParameters struct {

	// (String) The name of the filter.
	// The name of the filter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type FilterMetadataObservation struct {

	// (String) The name of the filter.
	// The name of the filter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type FilterMetadataParameters struct {

	// (String) The name of the filter.
	// The name of the filter.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type FilterObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Block List, Min: 1, Max: 1) Metadata of the filter. (see below for nested schema)
	// Metadata of the filter.
	Metadata []FilterMetadataObservation `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// (Block List, Min: 1, Max: 1) Specification of the filter. (see below for nested schema)
	// Specification of the filter.
	Spec []FilterSpecObservation `json:"spec,omitempty" tf:"spec,omitempty"`
}

type FilterParameters struct {

	// (Block List, Min: 1, Max: 1) Metadata of the filter. (see below for nested schema)
	// Metadata of the filter.
	// +kubebuilder:validation:Optional
	Metadata []FilterMetadataParameters `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// (Block List, Min: 1, Max: 1) Specification of the filter. (see below for nested schema)
	// Specification of the filter.
	// +kubebuilder:validation:Optional
	Spec []FilterSpecParameters `json:"spec,omitempty" tf:"spec,omitempty"`
}

type FilterSpecInitParameters struct {

	// (Block List, Min: 1, Max: 1) Filter group of the filter. (see below for nested schema)
	// Filter group of the filter.
	FilterGroup []FilterGroupInitParameters `json:"filterGroup,omitempty" tf:"filter_group,omitempty"`
}

type FilterSpecObservation struct {

	// (Block List, Min: 1, Max: 1) Filter group of the filter. (see below for nested schema)
	// Filter group of the filter.
	FilterGroup []FilterGroupObservation `json:"filterGroup,omitempty" tf:"filter_group,omitempty"`
}

type FilterSpecParameters struct {

	// (Block List, Min: 1, Max: 1) Filter group of the filter. (see below for nested schema)
	// Filter group of the filter.
	// +kubebuilder:validation:Optional
	FilterGroup []FilterGroupParameters `json:"filterGroup" tf:"filter_group,omitempty"`
}

type FiltersInitParameters struct {

	// (String) Key of the filter.
	// Key of the filter.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// (Boolean) Negation flag of the filter condition.
	// Negation flag of the filter condition.
	Negation *bool `json:"negation,omitempty" tf:"negation,omitempty"`

	// (String) Operator of the filter. Valid values are 'eq'.
	// Operator of the filter. Valid values are 'eq'.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// (List of String) Values of the filter.
	// Values of the filter.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type FiltersObservation struct {

	// (String) Key of the filter.
	// Key of the filter.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// (Boolean) Negation flag of the filter condition.
	// Negation flag of the filter condition.
	Negation *bool `json:"negation,omitempty" tf:"negation,omitempty"`

	// (String) Operator of the filter. Valid values are 'eq'.
	// Operator of the filter. Valid values are 'eq'.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// (List of String) Values of the filter.
	// Values of the filter.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type FiltersParameters struct {

	// (String) Key of the filter.
	// Key of the filter.
	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// (Boolean) Negation flag of the filter condition.
	// Negation flag of the filter condition.
	// +kubebuilder:validation:Optional
	Negation *bool `json:"negation,omitempty" tf:"negation,omitempty"`

	// (String) Operator of the filter. Valid values are 'eq'.
	// Operator of the filter. Valid values are 'eq'.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// (List of String) Values of the filter.
	// Values of the filter.
	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

// FilterSpec defines the desired state of Filter
type FilterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FilterParameters `json:"forProvider"`
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
	InitProvider FilterInitParameters `json:"initProvider,omitempty"`
}

// FilterStatus defines the observed state of Filter.
type FilterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FilterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Filter is the Schema for the Filters API. A resource for creating and managing filters.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,palette}
type Filter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.metadata) || (has(self.initProvider) && has(self.initProvider.metadata))",message="spec.forProvider.metadata is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.spec) || (has(self.initProvider) && has(self.initProvider.spec))",message="spec.forProvider.spec is a required parameter"
	Spec   FilterSpec   `json:"spec"`
	Status FilterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FilterList contains a list of Filters
type FilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Filter `json:"items"`
}

// Repository type metadata.
var (
	Filter_Kind             = "Filter"
	Filter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Filter_Kind}.String()
	Filter_KindAPIVersion   = Filter_Kind + "." + CRDGroupVersion.String()
	Filter_GroupVersionKind = CRDGroupVersion.WithKind(Filter_Kind)
)

func init() {
	SchemeBuilder.Register(&Filter{}, &FilterList{})
}
