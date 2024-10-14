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

type ApplicationInitParameters struct {

	// (String)
	// The unique identifier (UID) of the application profile to use for this application.
	ApplicationProfileUID *string `json:"applicationProfileUid,omitempty" tf:"application_profile_uid,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	// The configuration block for specifying cluster and resource limits for the application.
	Config []ConfigInitParameters `json:"config,omitempty" tf:"config,omitempty"`

	// (Set of String)
	// A set of tags to associate with the application for easier identification and categorization.
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ApplicationObservation struct {

	// (String)
	// The unique identifier (UID) of the application profile to use for this application.
	ApplicationProfileUID *string `json:"applicationProfileUid,omitempty" tf:"application_profile_uid,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	// The configuration block for specifying cluster and resource limits for the application.
	Config []ConfigObservation `json:"config,omitempty" tf:"config,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Set of String)
	// A set of tags to associate with the application for easier identification and categorization.
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ApplicationParameters struct {

	// (String)
	// The unique identifier (UID) of the application profile to use for this application.
	// +kubebuilder:validation:Optional
	ApplicationProfileUID *string `json:"applicationProfileUid,omitempty" tf:"application_profile_uid,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	// The configuration block for specifying cluster and resource limits for the application.
	// +kubebuilder:validation:Optional
	Config []ConfigParameters `json:"config,omitempty" tf:"config,omitempty"`

	// (Set of String)
	// A set of tags to associate with the application for easier identification and categorization.
	// +kubebuilder:validation:Optional
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ConfigInitParameters struct {

	// (String)
	// The context for the cluster,  Either `tenant` or `project` can be provided.
	ClusterContext *string `json:"clusterContext,omitempty" tf:"cluster_context,omitempty"`

	// (String)
	// The unique identifier (UID) of the cluster group. Either `cluster_uid` or `cluster_group_uid` can be provided.
	ClusterGroupUID *string `json:"clusterGroupUid,omitempty" tf:"cluster_group_uid,omitempty"`

	// (String)
	// An optional name for the target cluster.
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// (String)
	// The unique identifier (UID) of the target cluster. Either `cluster_uid` or `cluster_group_uid` can be provided.
	ClusterUID *string `json:"clusterUid,omitempty" tf:"cluster_uid,omitempty"`

	// (Block List) (see below for nested schema)
	// Optional resource limits for the application, including CPU, memory, and storage constraints.
	Limits []LimitsInitParameters `json:"limits,omitempty" tf:"limits,omitempty"`
}

type ConfigObservation struct {

	// (String)
	// The context for the cluster,  Either `tenant` or `project` can be provided.
	ClusterContext *string `json:"clusterContext,omitempty" tf:"cluster_context,omitempty"`

	// (String)
	// The unique identifier (UID) of the cluster group. Either `cluster_uid` or `cluster_group_uid` can be provided.
	ClusterGroupUID *string `json:"clusterGroupUid,omitempty" tf:"cluster_group_uid,omitempty"`

	// (String)
	// An optional name for the target cluster.
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// (String)
	// The unique identifier (UID) of the target cluster. Either `cluster_uid` or `cluster_group_uid` can be provided.
	ClusterUID *string `json:"clusterUid,omitempty" tf:"cluster_uid,omitempty"`

	// (Block List) (see below for nested schema)
	// Optional resource limits for the application, including CPU, memory, and storage constraints.
	Limits []LimitsObservation `json:"limits,omitempty" tf:"limits,omitempty"`
}

type ConfigParameters struct {

	// (String)
	// The context for the cluster,  Either `tenant` or `project` can be provided.
	// +kubebuilder:validation:Optional
	ClusterContext *string `json:"clusterContext" tf:"cluster_context,omitempty"`

	// (String)
	// The unique identifier (UID) of the cluster group. Either `cluster_uid` or `cluster_group_uid` can be provided.
	// +kubebuilder:validation:Optional
	ClusterGroupUID *string `json:"clusterGroupUid,omitempty" tf:"cluster_group_uid,omitempty"`

	// (String)
	// An optional name for the target cluster.
	// +kubebuilder:validation:Optional
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// (String)
	// The unique identifier (UID) of the target cluster. Either `cluster_uid` or `cluster_group_uid` can be provided.
	// +kubebuilder:validation:Optional
	ClusterUID *string `json:"clusterUid,omitempty" tf:"cluster_uid,omitempty"`

	// (Block List) (see below for nested schema)
	// Optional resource limits for the application, including CPU, memory, and storage constraints.
	// +kubebuilder:validation:Optional
	Limits []LimitsParameters `json:"limits,omitempty" tf:"limits,omitempty"`
}

type LimitsInitParameters struct {

	// (Number)
	// The CPU allocation for the application, specified in integer values.
	CPU *float64 `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// (Number)
	// The memory allocation for the application, specified in megabytes.
	Memory *float64 `json:"memory,omitempty" tf:"memory,omitempty"`

	// (Number)
	// The storage allocation for the application, specified in gigabytes.
	Storage *float64 `json:"storage,omitempty" tf:"storage,omitempty"`
}

type LimitsObservation struct {

	// (Number)
	// The CPU allocation for the application, specified in integer values.
	CPU *float64 `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// (Number)
	// The memory allocation for the application, specified in megabytes.
	Memory *float64 `json:"memory,omitempty" tf:"memory,omitempty"`

	// (Number)
	// The storage allocation for the application, specified in gigabytes.
	Storage *float64 `json:"storage,omitempty" tf:"storage,omitempty"`
}

type LimitsParameters struct {

	// (Number)
	// The CPU allocation for the application, specified in integer values.
	// +kubebuilder:validation:Optional
	CPU *float64 `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// (Number)
	// The memory allocation for the application, specified in megabytes.
	// +kubebuilder:validation:Optional
	Memory *float64 `json:"memory,omitempty" tf:"memory,omitempty"`

	// (Number)
	// The storage allocation for the application, specified in gigabytes.
	// +kubebuilder:validation:Optional
	Storage *float64 `json:"storage,omitempty" tf:"storage,omitempty"`
}

// ApplicationSpec defines the desired state of Application
type ApplicationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationParameters `json:"forProvider"`
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
	InitProvider ApplicationInitParameters `json:"initProvider,omitempty"`
}

// ApplicationStatus defines the observed state of Application.
type ApplicationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Application is the Schema for the Applications API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,palette}
type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.applicationProfileUid) || (has(self.initProvider) && has(self.initProvider.applicationProfileUid))",message="spec.forProvider.applicationProfileUid is a required parameter"
	Spec   ApplicationSpec   `json:"spec"`
	Status ApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationList contains a list of Applications
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Application `json:"items"`
}

// Repository type metadata.
var (
	Application_Kind             = "Application"
	Application_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Application_Kind}.String()
	Application_KindAPIVersion   = Application_Kind + "." + CRDGroupVersion.String()
	Application_GroupVersionKind = CRDGroupVersion.WithKind(Application_Kind)
)

func init() {
	SchemeBuilder.Register(&Application{}, &ApplicationList{})
}
