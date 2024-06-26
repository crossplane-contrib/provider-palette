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

type AddVolumeOptionsInitParameters struct {
	Disk []DiskInitParameters `json:"disk,omitempty" tf:"disk,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	VolumeSource []VolumeSourceInitParameters `json:"volumeSource,omitempty" tf:"volume_source,omitempty"`
}

type AddVolumeOptionsObservation struct {
	Disk []DiskObservation `json:"disk,omitempty" tf:"disk,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	VolumeSource []VolumeSourceObservation `json:"volumeSource,omitempty" tf:"volume_source,omitempty"`
}

type AddVolumeOptionsParameters struct {

	// +kubebuilder:validation:Optional
	Disk []DiskParameters `json:"disk" tf:"disk,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	VolumeSource []VolumeSourceParameters `json:"volumeSource" tf:"volume_source,omitempty"`
}

type BlankInitParameters struct {
}

type BlankObservation struct {
}

type BlankParameters struct {
}

type DataVolumeInitParameters struct {
	Hotpluggable *bool `json:"hotpluggable,omitempty" tf:"hotpluggable,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type DataVolumeObservation struct {
	Hotpluggable *bool `json:"hotpluggable,omitempty" tf:"hotpluggable,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type DataVolumeParameters struct {

	// +kubebuilder:validation:Optional
	Hotpluggable *bool `json:"hotpluggable,omitempty" tf:"hotpluggable,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type DatavolumeInitParameters struct {

	// DataVolumeSpec defines our specification for a DataVolume type
	AddVolumeOptions []AddVolumeOptionsInitParameters `json:"addVolumeOptions,omitempty" tf:"add_volume_options,omitempty"`

	ClusterContext *string `json:"clusterContext,omitempty" tf:"cluster_context,omitempty"`

	// The cluster UID to which the virtual machine belongs to.
	ClusterUID *string `json:"clusterUid,omitempty" tf:"cluster_uid,omitempty"`

	// Standard DataVolume's metadata. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata []MetadataInitParameters `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// DataVolumeSpec defines our specification for a DataVolume type
	Spec []SpecInitParameters `json:"spec,omitempty" tf:"spec,omitempty"`

	// DataVolumeStatus provides the parameters to store the phase of the Data Volume
	Status []DatavolumeStatusInitParameters `json:"status,omitempty" tf:"status,omitempty"`

	// The name of the virtual machine to which the data volume belongs to.
	VMName *string `json:"vmName,omitempty" tf:"vm_name,omitempty"`

	// The namespace of the virtual machine to which the data volume belongs to.
	VMNamespace *string `json:"vmNamespace,omitempty" tf:"vm_namespace,omitempty"`
}

type DatavolumeObservation struct {

	// DataVolumeSpec defines our specification for a DataVolume type
	AddVolumeOptions []AddVolumeOptionsObservation `json:"addVolumeOptions,omitempty" tf:"add_volume_options,omitempty"`

	ClusterContext *string `json:"clusterContext,omitempty" tf:"cluster_context,omitempty"`

	// The cluster UID to which the virtual machine belongs to.
	ClusterUID *string `json:"clusterUid,omitempty" tf:"cluster_uid,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Standard DataVolume's metadata. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata []MetadataObservation `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// DataVolumeSpec defines our specification for a DataVolume type
	Spec []SpecObservation `json:"spec,omitempty" tf:"spec,omitempty"`

	// DataVolumeStatus provides the parameters to store the phase of the Data Volume
	Status []DatavolumeStatusObservation `json:"status,omitempty" tf:"status,omitempty"`

	// The name of the virtual machine to which the data volume belongs to.
	VMName *string `json:"vmName,omitempty" tf:"vm_name,omitempty"`

	// The namespace of the virtual machine to which the data volume belongs to.
	VMNamespace *string `json:"vmNamespace,omitempty" tf:"vm_namespace,omitempty"`
}

type DatavolumeParameters struct {

	// DataVolumeSpec defines our specification for a DataVolume type
	// +kubebuilder:validation:Optional
	AddVolumeOptions []AddVolumeOptionsParameters `json:"addVolumeOptions,omitempty" tf:"add_volume_options,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterContext *string `json:"clusterContext,omitempty" tf:"cluster_context,omitempty"`

	// The cluster UID to which the virtual machine belongs to.
	// +kubebuilder:validation:Optional
	ClusterUID *string `json:"clusterUid,omitempty" tf:"cluster_uid,omitempty"`

	// Standard DataVolume's metadata. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +kubebuilder:validation:Optional
	Metadata []MetadataParameters `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// DataVolumeSpec defines our specification for a DataVolume type
	// +kubebuilder:validation:Optional
	Spec []SpecParameters `json:"spec,omitempty" tf:"spec,omitempty"`

	// DataVolumeStatus provides the parameters to store the phase of the Data Volume
	// +kubebuilder:validation:Optional
	Status []DatavolumeStatusParameters `json:"status,omitempty" tf:"status,omitempty"`

	// The name of the virtual machine to which the data volume belongs to.
	// +kubebuilder:validation:Optional
	VMName *string `json:"vmName,omitempty" tf:"vm_name,omitempty"`

	// The namespace of the virtual machine to which the data volume belongs to.
	// +kubebuilder:validation:Optional
	VMNamespace *string `json:"vmNamespace,omitempty" tf:"vm_namespace,omitempty"`
}

type DatavolumeStatusInitParameters struct {

	// DataVolumePhase is the current phase of the DataVolume.
	Phase *string `json:"phase,omitempty" tf:"phase,omitempty"`

	// DataVolumePhase is the current phase of the DataVolume.
	Progress *string `json:"progress,omitempty" tf:"progress,omitempty"`
}

type DatavolumeStatusObservation struct {

	// DataVolumePhase is the current phase of the DataVolume.
	Phase *string `json:"phase,omitempty" tf:"phase,omitempty"`

	// DataVolumePhase is the current phase of the DataVolume.
	Progress *string `json:"progress,omitempty" tf:"progress,omitempty"`
}

type DatavolumeStatusParameters struct {

	// DataVolumePhase is the current phase of the DataVolume.
	// +kubebuilder:validation:Optional
	Phase *string `json:"phase,omitempty" tf:"phase,omitempty"`

	// DataVolumePhase is the current phase of the DataVolume.
	// +kubebuilder:validation:Optional
	Progress *string `json:"progress,omitempty" tf:"progress,omitempty"`
}

type DiskInitParameters struct {
	Bus *string `json:"bus,omitempty" tf:"bus,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type DiskObservation struct {
	Bus *string `json:"bus,omitempty" tf:"bus,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type DiskParameters struct {

	// +kubebuilder:validation:Optional
	Bus *string `json:"bus" tf:"bus,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type MatchExpressionsInitParameters struct {

	// The label key that the selector applies to.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// A key's relationship to a set of values. Valid operators ard `In`, `NotIn`, `Exists` and `DoesNotExist`.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// An array of string values. If the operator is `In` or `NotIn`, the values array must be non-empty. If the operator is `Exists` or `DoesNotExist`, the values array must be empty. This array is replaced during a strategic merge patch.
	// +listType=set
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type MatchExpressionsObservation struct {

	// The label key that the selector applies to.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// A key's relationship to a set of values. Valid operators ard `In`, `NotIn`, `Exists` and `DoesNotExist`.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// An array of string values. If the operator is `In` or `NotIn`, the values array must be non-empty. If the operator is `Exists` or `DoesNotExist`, the values array must be empty. This array is replaced during a strategic merge patch.
	// +listType=set
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type MatchExpressionsParameters struct {

	// The label key that the selector applies to.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// A key's relationship to a set of values. Valid operators ard `In`, `NotIn`, `Exists` and `DoesNotExist`.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// An array of string values. If the operator is `In` or `NotIn`, the values array must be non-empty. If the operator is `Exists` or `DoesNotExist`, the values array must be empty. This array is replaced during a strategic merge patch.
	// +kubebuilder:validation:Optional
	// +listType=set
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type MetadataInitParameters struct {

	// An unstructured key value map stored with the DataVolume that may be used to store arbitrary metadata. More info: http://kubernetes.io/docs/user-guide/annotations
	// +mapType=granular
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// Map of string keys and values that can be used to organize and categorize (scope and select) the DataVolume. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the DataVolume, must be unique. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Namespace defines the space within which name of the DataVolume must be unique.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type MetadataObservation struct {

	// An unstructured key value map stored with the DataVolume that may be used to store arbitrary metadata. More info: http://kubernetes.io/docs/user-guide/annotations
	// +mapType=granular
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// A sequence number representing a specific generation of the desired state.
	Generation *float64 `json:"generation,omitempty" tf:"generation,omitempty"`

	// Map of string keys and values that can be used to organize and categorize (scope and select) the DataVolume. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the DataVolume, must be unique. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Namespace defines the space within which name of the DataVolume must be unique.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// An opaque value that represents the internal version of this DataVolume that can be used by clients to determine when DataVolume has changed. Read more: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
	ResourceVersion *string `json:"resourceVersion,omitempty" tf:"resource_version,omitempty"`

	// A URL representing this DataVolume.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// The unique in time and space value for this DataVolume. More info: http://kubernetes.io/docs/user-guide/identifiers#uids
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`
}

type MetadataParameters struct {

	// An unstructured key value map stored with the DataVolume that may be used to store arbitrary metadata. More info: http://kubernetes.io/docs/user-guide/annotations
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// Map of string keys and values that can be used to organize and categorize (scope and select) the DataVolume. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the DataVolume, must be unique. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Namespace defines the space within which name of the DataVolume must be unique.
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type PvcInitParameters struct {

	// A set of the desired access modes the volume should have. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#access-modes-1
	// +listType=set
	AccessModes []*string `json:"accessModes,omitempty" tf:"access_modes,omitempty"`

	// A list of the minimum resources the volume should have. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#resources
	Resources []ResourcesInitParameters `json:"resources,omitempty" tf:"resources,omitempty"`

	// A label query over volumes to consider for binding.
	Selector []SelectorInitParameters `json:"selector,omitempty" tf:"selector,omitempty"`

	// Name of the storage class requested by the claim
	StorageClassName *string `json:"storageClassName,omitempty" tf:"storage_class_name,omitempty"`

	// volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.
	VolumeMode *string `json:"volumeMode,omitempty" tf:"volume_mode,omitempty"`

	// The binding reference to the PersistentVolume backing this claim.
	VolumeName *string `json:"volumeName,omitempty" tf:"volume_name,omitempty"`
}

type PvcObservation struct {

	// A set of the desired access modes the volume should have. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#access-modes-1
	// +listType=set
	AccessModes []*string `json:"accessModes,omitempty" tf:"access_modes,omitempty"`

	// A list of the minimum resources the volume should have. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#resources
	Resources []ResourcesObservation `json:"resources,omitempty" tf:"resources,omitempty"`

	// A label query over volumes to consider for binding.
	Selector []SelectorObservation `json:"selector,omitempty" tf:"selector,omitempty"`

	// Name of the storage class requested by the claim
	StorageClassName *string `json:"storageClassName,omitempty" tf:"storage_class_name,omitempty"`

	// volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.
	VolumeMode *string `json:"volumeMode,omitempty" tf:"volume_mode,omitempty"`

	// The binding reference to the PersistentVolume backing this claim.
	VolumeName *string `json:"volumeName,omitempty" tf:"volume_name,omitempty"`
}

type PvcParameters struct {

	// A set of the desired access modes the volume should have. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#access-modes-1
	// +kubebuilder:validation:Optional
	// +listType=set
	AccessModes []*string `json:"accessModes" tf:"access_modes,omitempty"`

	// A list of the minimum resources the volume should have. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#resources
	// +kubebuilder:validation:Optional
	Resources []ResourcesParameters `json:"resources" tf:"resources,omitempty"`

	// A label query over volumes to consider for binding.
	// +kubebuilder:validation:Optional
	Selector []SelectorParameters `json:"selector,omitempty" tf:"selector,omitempty"`

	// Name of the storage class requested by the claim
	// +kubebuilder:validation:Optional
	StorageClassName *string `json:"storageClassName,omitempty" tf:"storage_class_name,omitempty"`

	// volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.
	// +kubebuilder:validation:Optional
	VolumeMode *string `json:"volumeMode,omitempty" tf:"volume_mode,omitempty"`

	// The binding reference to the PersistentVolume backing this claim.
	// +kubebuilder:validation:Optional
	VolumeName *string `json:"volumeName,omitempty" tf:"volume_name,omitempty"`
}

type RegistryInitParameters struct {

	// The registry URL of the image to download.
	ImageURL *string `json:"imageUrl,omitempty" tf:"image_url,omitempty"`
}

type RegistryObservation struct {

	// The registry URL of the image to download.
	ImageURL *string `json:"imageUrl,omitempty" tf:"image_url,omitempty"`
}

type RegistryParameters struct {

	// The registry URL of the image to download.
	// +kubebuilder:validation:Optional
	ImageURL *string `json:"imageUrl,omitempty" tf:"image_url,omitempty"`
}

type ResourcesInitParameters struct {

	// Map describing the maximum amount of compute resources allowed. More info: http://kubernetes.io/docs/user-guide/compute-resources/
	// +mapType=granular
	Limits map[string]*string `json:"limits,omitempty" tf:"limits,omitempty"`

	// Map describing the minimum amount of compute resources required. If this is omitted for a container, it defaults to `limits` if that is explicitly specified, otherwise to an implementation-defined value. More info: http://kubernetes.io/docs/user-guide/compute-resources/
	// +mapType=granular
	Requests map[string]*string `json:"requests,omitempty" tf:"requests,omitempty"`
}

type ResourcesObservation struct {

	// Map describing the maximum amount of compute resources allowed. More info: http://kubernetes.io/docs/user-guide/compute-resources/
	// +mapType=granular
	Limits map[string]*string `json:"limits,omitempty" tf:"limits,omitempty"`

	// Map describing the minimum amount of compute resources required. If this is omitted for a container, it defaults to `limits` if that is explicitly specified, otherwise to an implementation-defined value. More info: http://kubernetes.io/docs/user-guide/compute-resources/
	// +mapType=granular
	Requests map[string]*string `json:"requests,omitempty" tf:"requests,omitempty"`
}

type ResourcesParameters struct {

	// Map describing the maximum amount of compute resources allowed. More info: http://kubernetes.io/docs/user-guide/compute-resources/
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Limits map[string]*string `json:"limits,omitempty" tf:"limits,omitempty"`

	// Map describing the minimum amount of compute resources required. If this is omitted for a container, it defaults to `limits` if that is explicitly specified, otherwise to an implementation-defined value. More info: http://kubernetes.io/docs/user-guide/compute-resources/
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Requests map[string]*string `json:"requests,omitempty" tf:"requests,omitempty"`
}

type SelectorInitParameters struct {

	// A list of label selector requirements. The requirements are ANDed.
	MatchExpressions []MatchExpressionsInitParameters `json:"matchExpressions,omitempty" tf:"match_expressions,omitempty"`

	// A map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of `match_expressions`, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed.
	// +mapType=granular
	MatchLabels map[string]*string `json:"matchLabels,omitempty" tf:"match_labels,omitempty"`
}

type SelectorObservation struct {

	// A list of label selector requirements. The requirements are ANDed.
	MatchExpressions []MatchExpressionsObservation `json:"matchExpressions,omitempty" tf:"match_expressions,omitempty"`

	// A map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of `match_expressions`, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed.
	// +mapType=granular
	MatchLabels map[string]*string `json:"matchLabels,omitempty" tf:"match_labels,omitempty"`
}

type SelectorParameters struct {

	// A list of label selector requirements. The requirements are ANDed.
	// +kubebuilder:validation:Optional
	MatchExpressions []MatchExpressionsParameters `json:"matchExpressions,omitempty" tf:"match_expressions,omitempty"`

	// A map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of `match_expressions`, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	MatchLabels map[string]*string `json:"matchLabels,omitempty" tf:"match_labels,omitempty"`
}

type SourceHTTPInitParameters struct {

	// Cert_config_map provides a reference to the Registry certs.
	CertConfigMap *string `json:"certConfigMap,omitempty" tf:"cert_config_map,omitempty"`

	// Secret_ref provides the secret reference needed to access the HTTP source.
	SecretRef *string `json:"secretRef,omitempty" tf:"secret_ref,omitempty"`

	// url is the URL of the http source.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type SourceHTTPObservation struct {

	// Cert_config_map provides a reference to the Registry certs.
	CertConfigMap *string `json:"certConfigMap,omitempty" tf:"cert_config_map,omitempty"`

	// Secret_ref provides the secret reference needed to access the HTTP source.
	SecretRef *string `json:"secretRef,omitempty" tf:"secret_ref,omitempty"`

	// url is the URL of the http source.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type SourceHTTPParameters struct {

	// Cert_config_map provides a reference to the Registry certs.
	// +kubebuilder:validation:Optional
	CertConfigMap *string `json:"certConfigMap,omitempty" tf:"cert_config_map,omitempty"`

	// Secret_ref provides the secret reference needed to access the HTTP source.
	// +kubebuilder:validation:Optional
	SecretRef *string `json:"secretRef,omitempty" tf:"secret_ref,omitempty"`

	// url is the URL of the http source.
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type SourceInitParameters struct {

	// DataVolumeSourceBlank provides the parameters to create a Data Volume from an empty source.
	Blank []BlankInitParameters `json:"blank,omitempty" tf:"blank,omitempty"`

	// DataVolumeSourceHTTP provides the parameters to create a Data Volume from an HTTP source.
	HTTP []SourceHTTPInitParameters `json:"http,omitempty" tf:"http,omitempty"`

	// DataVolumeSourcePVC provides the parameters to create a Data Volume from an existing PVC.
	Pvc []SourcePvcInitParameters `json:"pvc,omitempty" tf:"pvc,omitempty"`

	// DataVolumeSourceRegistry provides the parameters to create a Data Volume from an existing PVC.
	Registry []RegistryInitParameters `json:"registry,omitempty" tf:"registry,omitempty"`
}

type SourceObservation struct {

	// DataVolumeSourceBlank provides the parameters to create a Data Volume from an empty source.
	Blank []BlankParameters `json:"blank,omitempty" tf:"blank,omitempty"`

	// DataVolumeSourceHTTP provides the parameters to create a Data Volume from an HTTP source.
	HTTP []SourceHTTPObservation `json:"http,omitempty" tf:"http,omitempty"`

	// DataVolumeSourcePVC provides the parameters to create a Data Volume from an existing PVC.
	Pvc []SourcePvcObservation `json:"pvc,omitempty" tf:"pvc,omitempty"`

	// DataVolumeSourceRegistry provides the parameters to create a Data Volume from an existing PVC.
	Registry []RegistryObservation `json:"registry,omitempty" tf:"registry,omitempty"`
}

type SourceParameters struct {

	// DataVolumeSourceBlank provides the parameters to create a Data Volume from an empty source.
	// +kubebuilder:validation:Optional
	Blank []BlankParameters `json:"blank,omitempty" tf:"blank,omitempty"`

	// DataVolumeSourceHTTP provides the parameters to create a Data Volume from an HTTP source.
	// +kubebuilder:validation:Optional
	HTTP []SourceHTTPParameters `json:"http,omitempty" tf:"http,omitempty"`

	// DataVolumeSourcePVC provides the parameters to create a Data Volume from an existing PVC.
	// +kubebuilder:validation:Optional
	Pvc []SourcePvcParameters `json:"pvc,omitempty" tf:"pvc,omitempty"`

	// DataVolumeSourceRegistry provides the parameters to create a Data Volume from an existing PVC.
	// +kubebuilder:validation:Optional
	Registry []RegistryParameters `json:"registry,omitempty" tf:"registry,omitempty"`
}

type SourcePvcInitParameters struct {

	// The name of the PVC.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace which the PVC located in.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type SourcePvcObservation struct {

	// The name of the PVC.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace which the PVC located in.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type SourcePvcParameters struct {

	// The name of the PVC.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace which the PVC located in.
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type SpecInitParameters struct {

	// ContentType options: "kubevirt", "archive".
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// PVC is a pointer to the PVC Spec we want to use.
	Pvc []PvcInitParameters `json:"pvc,omitempty" tf:"pvc,omitempty"`

	// Source is the src of the data for the requested DataVolume.
	Source []SourceInitParameters `json:"source,omitempty" tf:"source,omitempty"`
}

type SpecObservation struct {

	// ContentType options: "kubevirt", "archive".
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// PVC is a pointer to the PVC Spec we want to use.
	Pvc []PvcObservation `json:"pvc,omitempty" tf:"pvc,omitempty"`

	// Source is the src of the data for the requested DataVolume.
	Source []SourceObservation `json:"source,omitempty" tf:"source,omitempty"`
}

type SpecParameters struct {

	// ContentType options: "kubevirt", "archive".
	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// PVC is a pointer to the PVC Spec we want to use.
	// +kubebuilder:validation:Optional
	Pvc []PvcParameters `json:"pvc" tf:"pvc,omitempty"`

	// Source is the src of the data for the requested DataVolume.
	// +kubebuilder:validation:Optional
	Source []SourceParameters `json:"source,omitempty" tf:"source,omitempty"`
}

type VolumeSourceInitParameters struct {
	DataVolume []DataVolumeInitParameters `json:"dataVolume,omitempty" tf:"data_volume,omitempty"`
}

type VolumeSourceObservation struct {
	DataVolume []DataVolumeObservation `json:"dataVolume,omitempty" tf:"data_volume,omitempty"`
}

type VolumeSourceParameters struct {

	// +kubebuilder:validation:Optional
	DataVolume []DataVolumeParameters `json:"dataVolume" tf:"data_volume,omitempty"`
}

// DatavolumeSpec defines the desired state of Datavolume
type DatavolumeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatavolumeParameters `json:"forProvider"`
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
	InitProvider DatavolumeInitParameters `json:"initProvider,omitempty"`
}

// DatavolumeStatus defines the observed state of Datavolume.
type DatavolumeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatavolumeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Datavolume is the Schema for the Datavolumes API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,palette}
type Datavolume struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.addVolumeOptions) || (has(self.initProvider) && has(self.initProvider.addVolumeOptions))",message="spec.forProvider.addVolumeOptions is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clusterContext) || (has(self.initProvider) && has(self.initProvider.clusterContext))",message="spec.forProvider.clusterContext is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.metadata) || (has(self.initProvider) && has(self.initProvider.metadata))",message="spec.forProvider.metadata is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.spec) || (has(self.initProvider) && has(self.initProvider.spec))",message="spec.forProvider.spec is a required parameter"
	Spec   DatavolumeSpec   `json:"spec"`
	Status DatavolumeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatavolumeList contains a list of Datavolumes
type DatavolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Datavolume `json:"items"`
}

// Repository type metadata.
var (
	Datavolume_Kind             = "Datavolume"
	Datavolume_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Datavolume_Kind}.String()
	Datavolume_KindAPIVersion   = Datavolume_Kind + "." + CRDGroupVersion.String()
	Datavolume_GroupVersionKind = CRDGroupVersion.WithKind(Datavolume_Kind)
)

func init() {
	SchemeBuilder.Register(&Datavolume{}, &DatavolumeList{})
}
