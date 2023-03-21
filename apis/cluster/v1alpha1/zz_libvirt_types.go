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

type AttachedDisksObservation struct {
}

type AttachedDisksParameters struct {

	// +kubebuilder:validation:Optional
	Managed *bool `json:"managed,omitempty" tf:"managed,omitempty"`

	// +kubebuilder:validation:Required
	SizeInGb *float64 `json:"sizeInGb" tf:"size_in_gb,omitempty"`
}

type GpuConfigObservation struct {
}

type GpuConfigParameters struct {

	// +kubebuilder:validation:Optional
	Addresses map[string]*string `json:"addresses,omitempty" tf:"addresses,omitempty"`

	// +kubebuilder:validation:Required
	DeviceModel *string `json:"deviceModel" tf:"device_model,omitempty"`

	// +kubebuilder:validation:Required
	NumGpus *float64 `json:"numGpus" tf:"num_gpus,omitempty"`

	// +kubebuilder:validation:Required
	Vendor *string `json:"vendor" tf:"vendor,omitempty"`
}

type LibvirtBackupPolicyObservation struct {
}

type LibvirtBackupPolicyParameters struct {

	// The ID of the backup location to use for the backup.
	// +kubebuilder:validation:Required
	BackupLocationID *string `json:"backupLocationId" tf:"backup_location_id,omitempty"`

	// The list of cluster UIDs to include in the backup. If `include_all_clusters` is set to `true`, then all clusters will be included.
	// +kubebuilder:validation:Optional
	ClusterUids []*string `json:"clusterUids,omitempty" tf:"cluster_uids,omitempty"`

	// The number of hours after which the backup will be deleted. For example, if the expiry is set to 24, the backup will be deleted after 24 hours.
	// +kubebuilder:validation:Required
	ExpiryInHour *float64 `json:"expiryInHour" tf:"expiry_in_hour,omitempty"`

	// Whether to include all clusters in the backup. If set to false, only the clusters specified in `cluster_uids` will be included.
	// +kubebuilder:validation:Optional
	IncludeAllClusters *bool `json:"includeAllClusters,omitempty" tf:"include_all_clusters,omitempty"`

	// Whether to include the cluster resources in the backup. If set to false, only the cluster configuration and disks will be backed up.
	// +kubebuilder:validation:Optional
	IncludeClusterResources *bool `json:"includeClusterResources,omitempty" tf:"include_cluster_resources,omitempty"`

	// Whether to include the disks in the backup. If set to false, only the cluster configuration will be backed up.
	// +kubebuilder:validation:Optional
	IncludeDisks *bool `json:"includeDisks,omitempty" tf:"include_disks,omitempty"`

	// The list of Kubernetes namespaces to include in the backup. If not specified, all namespaces will be included.
	// +kubebuilder:validation:Optional
	Namespaces []*string `json:"namespaces,omitempty" tf:"namespaces,omitempty"`

	// Prefix for the backup name. The backup name will be of the format <prefix>-<cluster-name>-<timestamp>.
	// +kubebuilder:validation:Required
	Prefix *string `json:"prefix" tf:"prefix,omitempty"`

	// The schedule for the backup. The schedule is specified in cron format. For example, to run the backup every day at 1:00 AM, the schedule should be set to `0 1 * * *`.
	// +kubebuilder:validation:Required
	Schedule *string `json:"schedule" tf:"schedule,omitempty"`
}

type LibvirtCloudConfigObservation struct {
}

type LibvirtCloudConfigParameters struct {

	// +kubebuilder:validation:Optional
	NtpServers []*string `json:"ntpServers,omitempty" tf:"ntp_servers,omitempty"`

	// +kubebuilder:validation:Required
	SSHKey *string `json:"sshKey" tf:"ssh_key,omitempty"`

	// +kubebuilder:validation:Required
	Vip *string `json:"vip" tf:"vip,omitempty"`
}

type LibvirtClusterProfileObservation struct {

	// +kubebuilder:validation:Optional
	Pack []LibvirtClusterProfilePackObservation `json:"pack,omitempty" tf:"pack,omitempty"`
}

type LibvirtClusterProfilePackManifestObservation struct {
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`
}

type LibvirtClusterProfilePackManifestParameters struct {

	// The content of the manifest. The content is the YAML content of the manifest.
	// +kubebuilder:validation:Required
	Content *string `json:"content" tf:"content,omitempty"`

	// The name of the manifest. The name must be unique within the pack.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type LibvirtClusterProfilePackObservation struct {

	// +kubebuilder:validation:Optional
	Manifest []LibvirtClusterProfilePackManifestObservation `json:"manifest,omitempty" tf:"manifest,omitempty"`
}

type LibvirtClusterProfilePackParameters struct {

	// +kubebuilder:validation:Optional
	Manifest []LibvirtClusterProfilePackManifestParameters `json:"manifest,omitempty" tf:"manifest,omitempty"`

	// The name of the pack. The name must be unique within the cluster profile.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The registry UID of the pack. The registry UID is the unique identifier of the registry.
	// +kubebuilder:validation:Optional
	RegistryUID *string `json:"registryUid,omitempty" tf:"registry_uid,omitempty"`

	// The tag of the pack. The tag is the version of the pack.
	// +kubebuilder:validation:Optional
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

	// The type of the pack. The default value is `spectro`.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`

	// The values of the pack. The values are the configuration values of the pack. The values are specified in YAML format.
	// +kubebuilder:validation:Required
	Values *string `json:"values" tf:"values,omitempty"`
}

type LibvirtClusterProfileParameters struct {

	// The ID of the cluster profile.
	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// +kubebuilder:validation:Optional
	Pack []LibvirtClusterProfilePackParameters `json:"pack,omitempty" tf:"pack,omitempty"`
}

type LibvirtClusterRbacBindingObservation struct {
}

type LibvirtClusterRbacBindingParameters struct {

	// The Kubernetes namespace of the RBAC binding. Required if 'type' is set to 'RoleBinding'.
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The role of the RBAC binding. Required if 'type' is set to 'RoleBinding'.
	// +kubebuilder:validation:Optional
	Role map[string]*string `json:"role,omitempty" tf:"role,omitempty"`

	// +kubebuilder:validation:Optional
	Subjects []LibvirtClusterRbacBindingSubjectsParameters `json:"subjects,omitempty" tf:"subjects,omitempty"`

	// The type of the RBAC binding. Can be one of the following values: `RoleBinding`, or `ClusterRoleBinding`.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type LibvirtClusterRbacBindingSubjectsObservation struct {
}

type LibvirtClusterRbacBindingSubjectsParameters struct {

	// The name of the subject. Required if 'type' is set to 'User' or 'Group'.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The Kubernetes namespace of the subject. Required if 'type' is set to 'ServiceAccount'.
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The type of the subject. Can be one of the following values: `User`, `Group`, or `ServiceAccount`.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type LibvirtHostConfigObservation struct {
}

type LibvirtHostConfigParameters struct {

	// The external traffic policy for the cluster.
	// +kubebuilder:validation:Optional
	ExternalTrafficPolicy *string `json:"externalTrafficPolicy,omitempty" tf:"external_traffic_policy,omitempty"`

	// The type of endpoint for the cluster. Can be either 'Ingress' or 'LoadBalancer'. The default is 'Ingress'.
	// +kubebuilder:validation:Optional
	HostEndpointType *string `json:"hostEndpointType,omitempty" tf:"host_endpoint_type,omitempty"`

	// The host for the Ingress endpoint. Required if 'host_endpoint_type' is set to 'Ingress'.
	// +kubebuilder:validation:Optional
	IngressHost *string `json:"ingressHost,omitempty" tf:"ingress_host,omitempty"`

	// The source ranges for the load balancer. Required if 'host_endpoint_type' is set to 'LoadBalancer'.
	// +kubebuilder:validation:Optional
	LoadBalancerSourceRanges *string `json:"loadBalancerSourceRanges,omitempty" tf:"load_balancer_source_ranges,omitempty"`
}

type LibvirtLocationConfigObservation struct {
}

type LibvirtLocationConfigParameters struct {

	// The country code of the country the cluster is located in.
	// +kubebuilder:validation:Optional
	CountryCode *string `json:"countryCode,omitempty" tf:"country_code,omitempty"`

	// The name of the country.
	// +kubebuilder:validation:Optional
	CountryName *string `json:"countryName,omitempty" tf:"country_name,omitempty"`

	// The latitude coordinates value.
	// +kubebuilder:validation:Required
	Latitude *float64 `json:"latitude" tf:"latitude,omitempty"`

	// The longitude coordinates value.
	// +kubebuilder:validation:Required
	Longitude *float64 `json:"longitude" tf:"longitude,omitempty"`

	// The region code of where the cluster is located in.
	// +kubebuilder:validation:Optional
	RegionCode *string `json:"regionCode,omitempty" tf:"region_code,omitempty"`

	// The name of the region.
	// +kubebuilder:validation:Optional
	RegionName *string `json:"regionName,omitempty" tf:"region_name,omitempty"`
}

type LibvirtMachinePoolObservation struct {
}

type LibvirtMachinePoolParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalLabels map[string]*string `json:"additionalLabels,omitempty" tf:"additional_labels,omitempty"`

	// Whether this machine pool is a control plane. Defaults to `false`.
	// +kubebuilder:validation:Optional
	ControlPlane *bool `json:"controlPlane,omitempty" tf:"control_plane,omitempty"`

	// Whether this machine pool is a control plane and a worker. Defaults to `false`.
	// +kubebuilder:validation:Optional
	ControlPlaneAsWorker *bool `json:"controlPlaneAsWorker,omitempty" tf:"control_plane_as_worker,omitempty"`

	// Number of nodes in the machine pool.
	// +kubebuilder:validation:Required
	Count *float64 `json:"count" tf:"count,omitempty"`

	// +kubebuilder:validation:Required
	InstanceType []MachinePoolInstanceTypeParameters `json:"instanceType" tf:"instance_type,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Placements []MachinePoolPlacementsParameters `json:"placements" tf:"placements,omitempty"`

	// +kubebuilder:validation:Optional
	Taints []LibvirtMachinePoolTaintsParameters `json:"taints,omitempty" tf:"taints,omitempty"`

	// Update strategy for the machine pool. Valid values are `RollingUpdateScaleOut` and `RollingUpdateScaleIn`.
	// +kubebuilder:validation:Optional
	UpdateStrategy *string `json:"updateStrategy,omitempty" tf:"update_strategy,omitempty"`
}

type LibvirtMachinePoolTaintsObservation struct {
}

type LibvirtMachinePoolTaintsParameters struct {

	// The effect of the taint. Allowed values are: `NoSchedule`, `PreferNoSchedule` or `NoExecute`.
	// +kubebuilder:validation:Required
	Effect *string `json:"effect" tf:"effect,omitempty"`

	// The key of the taint.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// The value of the taint.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type LibvirtNamespacesObservation struct {
}

type LibvirtNamespacesParameters struct {

	// List of images to disallow for the namespace. For example, `['nginx:latest', 'redis:latest']`
	// +kubebuilder:validation:Optional
	ImagesBlacklist []*string `json:"imagesBlacklist,omitempty" tf:"images_blacklist,omitempty"`

	// Name of the namespace. This is the name of the Kubernetes namespace in the cluster.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Resource allocation for the namespace. This is a map containing the resource type and the resource value. For example, `{cpu_cores: '2', memory_MiB: '2048'}`
	// +kubebuilder:validation:Required
	ResourceAllocation map[string]*string `json:"resourceAllocation" tf:"resource_allocation,omitempty"`
}

type LibvirtObservation struct {

	// ID of the cloud config used for the cluster. This cloud config must be of type `azure`.
	CloudConfigID *string `json:"cloudConfigId,omitempty" tf:"cloud_config_id,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterProfile []LibvirtClusterProfileObservation `json:"clusterProfile,omitempty" tf:"cluster_profile,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Kubeconfig for the cluster. This can be used to connect to the cluster using `kubectl`.
	Kubeconfig *string `json:"kubeconfig,omitempty" tf:"kubeconfig,omitempty"`
}

type LibvirtParameters struct {

	// +kubebuilder:validation:Optional
	ApplySetting *string `json:"applySetting,omitempty" tf:"apply_setting,omitempty"`

	// +kubebuilder:validation:Optional
	BackupPolicy []LibvirtBackupPolicyParameters `json:"backupPolicy,omitempty" tf:"backup_policy,omitempty"`

	// +kubebuilder:validation:Optional
	CloudAccountID *string `json:"cloudAccountId,omitempty" tf:"cloud_account_id,omitempty"`

	// +kubebuilder:validation:Required
	CloudConfig []LibvirtCloudConfigParameters `json:"cloudConfig" tf:"cloud_config,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterProfile []LibvirtClusterProfileParameters `json:"clusterProfile,omitempty" tf:"cluster_profile,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterRbacBinding []LibvirtClusterRbacBindingParameters `json:"clusterRbacBinding,omitempty" tf:"cluster_rbac_binding,omitempty"`

	// +kubebuilder:validation:Optional
	HostConfig []LibvirtHostConfigParameters `json:"hostConfig,omitempty" tf:"host_config,omitempty"`

	// +kubebuilder:validation:Optional
	LocationConfig []LibvirtLocationConfigParameters `json:"locationConfig,omitempty" tf:"location_config,omitempty"`

	// +kubebuilder:validation:Required
	MachinePool []LibvirtMachinePoolParameters `json:"machinePool" tf:"machine_pool,omitempty"`

	// +kubebuilder:validation:Optional
	Namespaces []LibvirtNamespacesParameters `json:"namespaces,omitempty" tf:"namespaces,omitempty"`

	// Date and time after which to patch cluster `RFC3339: 2006-01-02T15:04:05Z07:00`
	// +kubebuilder:validation:Optional
	OsPatchAfter *string `json:"osPatchAfter,omitempty" tf:"os_patch_after,omitempty"`

	// Whether to apply OS patch on boot. Default is `false`.
	// +kubebuilder:validation:Optional
	OsPatchOnBoot *bool `json:"osPatchOnBoot,omitempty" tf:"os_patch_on_boot,omitempty"`

	// Cron schedule for OS patching. This must be in the form of `0 0 * * *`.
	// +kubebuilder:validation:Optional
	OsPatchSchedule *string `json:"osPatchSchedule,omitempty" tf:"os_patch_schedule,omitempty"`

	// +kubebuilder:validation:Optional
	ScanPolicy []LibvirtScanPolicyParameters `json:"scanPolicy,omitempty" tf:"scan_policy,omitempty"`

	// If `true`, the cluster will be created asynchronously. Default value is `false`.
	// +kubebuilder:validation:Optional
	SkipCompletion *bool `json:"skipCompletion,omitempty" tf:"skip_completion,omitempty"`

	// A list of tags to be applied to the cluster. Tags must be in the form of `key:value`.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type LibvirtScanPolicyObservation struct {
}

type LibvirtScanPolicyParameters struct {

	// The schedule for configuration scan.
	// +kubebuilder:validation:Required
	ConfigurationScanSchedule *string `json:"configurationScanSchedule" tf:"configuration_scan_schedule,omitempty"`

	// The schedule for conformance scan.
	// +kubebuilder:validation:Required
	ConformanceScanSchedule *string `json:"conformanceScanSchedule" tf:"conformance_scan_schedule,omitempty"`

	// The schedule for penetration scan.
	// +kubebuilder:validation:Required
	PenetrationScanSchedule *string `json:"penetrationScanSchedule" tf:"penetration_scan_schedule,omitempty"`
}

type MachinePoolInstanceTypeObservation struct {
}

type MachinePoolInstanceTypeParameters struct {

	// +kubebuilder:validation:Optional
	AttachedDisks []AttachedDisksParameters `json:"attachedDisks,omitempty" tf:"attached_disks,omitempty"`

	// +kubebuilder:validation:Required
	CPU *float64 `json:"cpu" tf:"cpu,omitempty"`

	// +kubebuilder:validation:Optional
	CachePassthrough *bool `json:"cachePassthrough,omitempty" tf:"cache_passthrough,omitempty"`

	// +kubebuilder:validation:Optional
	CpusSets *string `json:"cpusSets,omitempty" tf:"cpus_sets,omitempty"`

	// +kubebuilder:validation:Required
	DiskSizeGb *float64 `json:"diskSizeGb" tf:"disk_size_gb,omitempty"`

	// +kubebuilder:validation:Optional
	GpuConfig []GpuConfigParameters `json:"gpuConfig,omitempty" tf:"gpu_config,omitempty"`

	// +kubebuilder:validation:Required
	MemoryMb *float64 `json:"memoryMb" tf:"memory_mb,omitempty"`
}

type MachinePoolPlacementsObservation struct {
}

type MachinePoolPlacementsParameters struct {

	// +kubebuilder:validation:Required
	ApplianceID *string `json:"applianceId" tf:"appliance_id,omitempty"`

	// +kubebuilder:validation:Required
	DataStoragePool *string `json:"dataStoragePool" tf:"data_storage_pool,omitempty"`

	// +kubebuilder:validation:Required
	ImageStoragePool *string `json:"imageStoragePool" tf:"image_storage_pool,omitempty"`

	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// +kubebuilder:validation:Required
	NetworkNames *string `json:"networkNames" tf:"network_names,omitempty"`

	// +kubebuilder:validation:Required
	NetworkType *string `json:"networkType" tf:"network_type,omitempty"`

	// +kubebuilder:validation:Required
	TargetStoragePool *string `json:"targetStoragePool" tf:"target_storage_pool,omitempty"`
}

// LibvirtSpec defines the desired state of Libvirt
type LibvirtSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LibvirtParameters `json:"forProvider"`
}

// LibvirtStatus defines the observed state of Libvirt.
type LibvirtStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LibvirtObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Libvirt is the Schema for the Libvirts API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,palette}
type Libvirt struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LibvirtSpec   `json:"spec"`
	Status            LibvirtStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LibvirtList contains a list of Libvirts
type LibvirtList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Libvirt `json:"items"`
}

// Repository type metadata.
var (
	Libvirt_Kind             = "Libvirt"
	Libvirt_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Libvirt_Kind}.String()
	Libvirt_KindAPIVersion   = Libvirt_Kind + "." + CRDGroupVersion.String()
	Libvirt_GroupVersionKind = CRDGroupVersion.WithKind(Libvirt_Kind)
)

func init() {
	SchemeBuilder.Register(&Libvirt{}, &LibvirtList{})
}
