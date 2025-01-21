/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	deployment "github.com/crossplane-contrib/provider-palette/internal/controller/addon/deployment"
	profile "github.com/crossplane-contrib/provider-palette/internal/controller/application/profile"
	storagelocation "github.com/crossplane-contrib/provider-palette/internal/controller/backup/storagelocation"
	aws "github.com/crossplane-contrib/provider-palette/internal/controller/cloudaccount/aws"
	azure "github.com/crossplane-contrib/provider-palette/internal/controller/cloudaccount/azure"
	custom "github.com/crossplane-contrib/provider-palette/internal/controller/cloudaccount/custom"
	gcp "github.com/crossplane-contrib/provider-palette/internal/controller/cloudaccount/gcp"
	maas "github.com/crossplane-contrib/provider-palette/internal/controller/cloudaccount/maas"
	openstack "github.com/crossplane-contrib/provider-palette/internal/controller/cloudaccount/openstack"
	vsphere "github.com/crossplane-contrib/provider-palette/internal/controller/cloudaccount/vsphere"
	aks "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/aks"
	awscluster "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/aws"
	azurecluster "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/azure"
	customcloud "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/customcloud"
	edgenative "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/edgenative"
	edgevsphere "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/edgevsphere"
	eks "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/eks"
	gcpcluster "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/gcp"
	gke "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/gke"
	group "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/group"
	maascluster "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/maas"
	openstackcluster "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/openstack"
	profilecluster "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/profile"
	vspherecluster "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/vsphere"
	policy "github.com/crossplane-contrib/provider-palette/internal/controller/password/policy"
	dnsmap "github.com/crossplane-contrib/provider-palette/internal/controller/privatecloudgateway/dnsmap"
	ippool "github.com/crossplane-contrib/provider-palette/internal/controller/privatecloudgateway/ippool"
	providerconfig "github.com/crossplane-contrib/provider-palette/internal/controller/providerconfig"
	helm "github.com/crossplane-contrib/provider-palette/internal/controller/registry/helm"
	oci "github.com/crossplane-contrib/provider-palette/internal/controller/registry/oci"
	alert "github.com/crossplane-contrib/provider-palette/internal/controller/spectrocloud/alert"
	appliance "github.com/crossplane-contrib/provider-palette/internal/controller/spectrocloud/appliance"
	application "github.com/crossplane-contrib/provider-palette/internal/controller/spectrocloud/application"
	datavolume "github.com/crossplane-contrib/provider-palette/internal/controller/spectrocloud/datavolume"
	filter "github.com/crossplane-contrib/provider-palette/internal/controller/spectrocloud/filter"
	macro "github.com/crossplane-contrib/provider-palette/internal/controller/spectrocloud/macro"
	macros "github.com/crossplane-contrib/provider-palette/internal/controller/spectrocloud/macros"
	project "github.com/crossplane-contrib/provider-palette/internal/controller/spectrocloud/project"
	role "github.com/crossplane-contrib/provider-palette/internal/controller/spectrocloud/role"
	team "github.com/crossplane-contrib/provider-palette/internal/controller/spectrocloud/team"
	user "github.com/crossplane-contrib/provider-palette/internal/controller/spectrocloud/user"
	workspace "github.com/crossplane-contrib/provider-palette/internal/controller/spectrocloud/workspace"
	key "github.com/crossplane-contrib/provider-palette/internal/controller/ssh/key"
	cluster "github.com/crossplane-contrib/provider-palette/internal/controller/virtual/cluster"
	machine "github.com/crossplane-contrib/provider-palette/internal/controller/virtual/machine"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		deployment.Setup,
		profile.Setup,
		storagelocation.Setup,
		aws.Setup,
		azure.Setup,
		custom.Setup,
		gcp.Setup,
		maas.Setup,
		openstack.Setup,
		vsphere.Setup,
		aks.Setup,
		awscluster.Setup,
		azurecluster.Setup,
		customcloud.Setup,
		edgenative.Setup,
		edgevsphere.Setup,
		eks.Setup,
		gcpcluster.Setup,
		gke.Setup,
		group.Setup,
		maascluster.Setup,
		openstackcluster.Setup,
		profilecluster.Setup,
		vspherecluster.Setup,
		policy.Setup,
		dnsmap.Setup,
		ippool.Setup,
		providerconfig.Setup,
		helm.Setup,
		oci.Setup,
		alert.Setup,
		appliance.Setup,
		application.Setup,
		datavolume.Setup,
		filter.Setup,
		macro.Setup,
		macros.Setup,
		project.Setup,
		role.Setup,
		team.Setup,
		user.Setup,
		workspace.Setup,
		key.Setup,
		cluster.Setup,
		machine.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
