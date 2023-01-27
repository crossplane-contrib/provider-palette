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

type MacroObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type MacroParameters struct {

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// MacroSpec defines the desired state of Macro
type MacroSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MacroParameters `json:"forProvider"`
}

// MacroStatus defines the observed state of Macro.
type MacroStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MacroObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Macro is the Schema for the Macros API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,jet-palette}
type Macro struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MacroSpec   `json:"spec"`
	Status            MacroStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MacroList contains a list of Macros
type MacroList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Macro `json:"items"`
}

// Repository type metadata.
var (
	Macro_Kind             = "Macro"
	Macro_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Macro_Kind}.String()
	Macro_KindAPIVersion   = Macro_Kind + "." + CRDGroupVersion.String()
	Macro_GroupVersionKind = CRDGroupVersion.WithKind(Macro_Kind)
)

func init() {
	SchemeBuilder.Register(&Macro{}, &MacroList{})
}
