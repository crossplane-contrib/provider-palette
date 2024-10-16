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

type AlertInitParameters struct {

	// (Boolean) If set to true, the alert will be sent to all users. If false, it will target specific users or identifiers.
	// If set to `true`, the alert will be sent to all users. If `false`, it will target specific users or identifiers.
	AlertAllUsers *bool `json:"alertAllUsers,omitempty" tf:"alert_all_users,omitempty"`

	// (String) The component of the system that the alert is associated with. Currently, ClusterHealth is the only supported value.
	// The component of the system that the alert is associated with. Currently, `ClusterHealth` is the only supported value.
	Component *string `json:"component,omitempty" tf:"component,omitempty"`

	// (String) The user who created the alert.
	// The user who created the alert.
	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	// based alerts. This is used when the type is set to http. (see below for nested schema)
	// The configuration block for HTTP-based alerts. This is used when the `type` is set to `http`.
	HTTP []HTTPInitParameters `json:"http,omitempty" tf:"http,omitempty"`

	// (Set of String) A set of unique identifiers to which the alert will be sent. This is used to target specific users or groups.
	// A set of unique identifiers to which the alert will be sent. This is used to target specific users or groups.
	// +listType=set
	Identifiers []*string `json:"identifiers,omitempty" tf:"identifiers,omitempty"`

	// (Boolean) Indicates whether the alert is active. Set to true to activate the alert, or false to deactivate it.
	// Indicates whether the alert is active. Set to `true` to activate the alert, or `false` to deactivate it.
	IsActive *bool `json:"isActive,omitempty" tf:"is_active,omitempty"`

	// (String) The project to which the alert belongs to.
	// The project to which the alert belongs to.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// (Block List) A status block representing the internal status of the alert. This is primarily for internal use and not utilized directly. (see below for nested schema)
	// A status block representing the internal status of the alert. This is primarily for internal use and not utilized directly.
	Status []StatusInitParameters `json:"status,omitempty" tf:"status,omitempty"`

	// (String) The type of alert mechanism to use. Can be either email for email alerts or http for sending HTTP requests.
	// The type of alert mechanism to use. Can be either `email` for email alerts or `http` for sending HTTP requests.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AlertObservation struct {

	// (Boolean) If set to true, the alert will be sent to all users. If false, it will target specific users or identifiers.
	// If set to `true`, the alert will be sent to all users. If `false`, it will target specific users or identifiers.
	AlertAllUsers *bool `json:"alertAllUsers,omitempty" tf:"alert_all_users,omitempty"`

	// (String) The component of the system that the alert is associated with. Currently, ClusterHealth is the only supported value.
	// The component of the system that the alert is associated with. Currently, `ClusterHealth` is the only supported value.
	Component *string `json:"component,omitempty" tf:"component,omitempty"`

	// (String) The user who created the alert.
	// The user who created the alert.
	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	// based alerts. This is used when the type is set to http. (see below for nested schema)
	// The configuration block for HTTP-based alerts. This is used when the `type` is set to `http`.
	HTTP []HTTPObservation `json:"http,omitempty" tf:"http,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Set of String) A set of unique identifiers to which the alert will be sent. This is used to target specific users or groups.
	// A set of unique identifiers to which the alert will be sent. This is used to target specific users or groups.
	// +listType=set
	Identifiers []*string `json:"identifiers,omitempty" tf:"identifiers,omitempty"`

	// (Boolean) Indicates whether the alert is active. Set to true to activate the alert, or false to deactivate it.
	// Indicates whether the alert is active. Set to `true` to activate the alert, or `false` to deactivate it.
	IsActive *bool `json:"isActive,omitempty" tf:"is_active,omitempty"`

	// (String) The project to which the alert belongs to.
	// The project to which the alert belongs to.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// (Block List) A status block representing the internal status of the alert. This is primarily for internal use and not utilized directly. (see below for nested schema)
	// A status block representing the internal status of the alert. This is primarily for internal use and not utilized directly.
	Status []StatusObservation `json:"status,omitempty" tf:"status,omitempty"`

	// (String) The type of alert mechanism to use. Can be either email for email alerts or http for sending HTTP requests.
	// The type of alert mechanism to use. Can be either `email` for email alerts or `http` for sending HTTP requests.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AlertParameters struct {

	// (Boolean) If set to true, the alert will be sent to all users. If false, it will target specific users or identifiers.
	// If set to `true`, the alert will be sent to all users. If `false`, it will target specific users or identifiers.
	// +kubebuilder:validation:Optional
	AlertAllUsers *bool `json:"alertAllUsers,omitempty" tf:"alert_all_users,omitempty"`

	// (String) The component of the system that the alert is associated with. Currently, ClusterHealth is the only supported value.
	// The component of the system that the alert is associated with. Currently, `ClusterHealth` is the only supported value.
	// +kubebuilder:validation:Optional
	Component *string `json:"component,omitempty" tf:"component,omitempty"`

	// (String) The user who created the alert.
	// The user who created the alert.
	// +kubebuilder:validation:Optional
	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	// based alerts. This is used when the type is set to http. (see below for nested schema)
	// The configuration block for HTTP-based alerts. This is used when the `type` is set to `http`.
	// +kubebuilder:validation:Optional
	HTTP []HTTPParameters `json:"http,omitempty" tf:"http,omitempty"`

	// (Set of String) A set of unique identifiers to which the alert will be sent. This is used to target specific users or groups.
	// A set of unique identifiers to which the alert will be sent. This is used to target specific users or groups.
	// +kubebuilder:validation:Optional
	// +listType=set
	Identifiers []*string `json:"identifiers,omitempty" tf:"identifiers,omitempty"`

	// (Boolean) Indicates whether the alert is active. Set to true to activate the alert, or false to deactivate it.
	// Indicates whether the alert is active. Set to `true` to activate the alert, or `false` to deactivate it.
	// +kubebuilder:validation:Optional
	IsActive *bool `json:"isActive,omitempty" tf:"is_active,omitempty"`

	// (String) The project to which the alert belongs to.
	// The project to which the alert belongs to.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// (Block List) A status block representing the internal status of the alert. This is primarily for internal use and not utilized directly. (see below for nested schema)
	// A status block representing the internal status of the alert. This is primarily for internal use and not utilized directly.
	// +kubebuilder:validation:Optional
	Status []StatusParameters `json:"status,omitempty" tf:"status,omitempty"`

	// (String) The type of alert mechanism to use. Can be either email for email alerts or http for sending HTTP requests.
	// The type of alert mechanism to use. Can be either `email` for email alerts or `http` for sending HTTP requests.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type HTTPInitParameters struct {

	// (String) The payload to include in the HTTP request body when the alert is triggered.
	// The payload to include in the HTTP request body when the alert is triggered.
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// value pair.
	// Optional HTTP headers to include in the request. Each header should be specified as a key-value pair.
	// +mapType=granular
	Headers map[string]*string `json:"headers,omitempty" tf:"headers,omitempty"`

	// (String) The HTTP method to use for the alert. Supported values are POST, GET, and PUT.
	// The HTTP method to use for the alert. Supported values are `POST`, `GET`, and `PUT`.
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// (String) The target URL to send the HTTP request to when the alert is triggered.
	// The target URL to send the HTTP request to when the alert is triggered.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type HTTPObservation struct {

	// (String) The payload to include in the HTTP request body when the alert is triggered.
	// The payload to include in the HTTP request body when the alert is triggered.
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// value pair.
	// Optional HTTP headers to include in the request. Each header should be specified as a key-value pair.
	// +mapType=granular
	Headers map[string]*string `json:"headers,omitempty" tf:"headers,omitempty"`

	// (String) The HTTP method to use for the alert. Supported values are POST, GET, and PUT.
	// The HTTP method to use for the alert. Supported values are `POST`, `GET`, and `PUT`.
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// (String) The target URL to send the HTTP request to when the alert is triggered.
	// The target URL to send the HTTP request to when the alert is triggered.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type HTTPParameters struct {

	// (String) The payload to include in the HTTP request body when the alert is triggered.
	// The payload to include in the HTTP request body when the alert is triggered.
	// +kubebuilder:validation:Optional
	Body *string `json:"body" tf:"body,omitempty"`

	// value pair.
	// Optional HTTP headers to include in the request. Each header should be specified as a key-value pair.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Headers map[string]*string `json:"headers,omitempty" tf:"headers,omitempty"`

	// (String) The HTTP method to use for the alert. Supported values are POST, GET, and PUT.
	// The HTTP method to use for the alert. Supported values are `POST`, `GET`, and `PUT`.
	// +kubebuilder:validation:Optional
	Method *string `json:"method" tf:"method,omitempty"`

	// (String) The target URL to send the HTTP request to when the alert is triggered.
	// The target URL to send the HTTP request to when the alert is triggered.
	// +kubebuilder:validation:Optional
	URL *string `json:"url" tf:"url,omitempty"`
}

type StatusInitParameters struct {

	// (Boolean)
	IsSucceeded *bool `json:"isSucceeded,omitempty" tf:"is_succeeded,omitempty"`

	// (String)
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// (String)
	Time *string `json:"time,omitempty" tf:"time,omitempty"`
}

type StatusObservation struct {

	// (Boolean)
	IsSucceeded *bool `json:"isSucceeded,omitempty" tf:"is_succeeded,omitempty"`

	// (String)
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// (String)
	Time *string `json:"time,omitempty" tf:"time,omitempty"`
}

type StatusParameters struct {

	// (Boolean)
	// +kubebuilder:validation:Optional
	IsSucceeded *bool `json:"isSucceeded,omitempty" tf:"is_succeeded,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Time *string `json:"time,omitempty" tf:"time,omitempty"`
}

// AlertSpec defines the desired state of Alert
type AlertSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AlertParameters `json:"forProvider"`
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
	InitProvider AlertInitParameters `json:"initProvider,omitempty"`
}

// AlertStatus defines the observed state of Alert.
type AlertStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AlertObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Alert is the Schema for the Alerts API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,palette}
type Alert struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.component) || (has(self.initProvider) && has(self.initProvider.component))",message="spec.forProvider.component is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.isActive) || (has(self.initProvider) && has(self.initProvider.isActive))",message="spec.forProvider.isActive is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.project) || (has(self.initProvider) && has(self.initProvider.project))",message="spec.forProvider.project is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   AlertSpec   `json:"spec"`
	Status AlertStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlertList contains a list of Alerts
type AlertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Alert `json:"items"`
}

// Repository type metadata.
var (
	Alert_Kind             = "Alert"
	Alert_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Alert_Kind}.String()
	Alert_KindAPIVersion   = Alert_Kind + "." + CRDGroupVersion.String()
	Alert_GroupVersionKind = CRDGroupVersion.WithKind(Alert_Kind)
)

func init() {
	SchemeBuilder.Register(&Alert{}, &AlertList{})
}
