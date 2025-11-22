/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	deployment "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/addon/deployment"
	profile "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/application/profile"
	storagelocation "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/backup/storagelocation"
	aws "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/cloudaccount/aws"
	azure "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/cloudaccount/azure"
	custom "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/cloudaccount/custom"
	gcp "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/cloudaccount/gcp"
	maas "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/cloudaccount/maas"
	openstack "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/cloudaccount/openstack"
	vsphere "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/cloudaccount/vsphere"
	aks "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/cluster/aks"
	awscluster "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/cluster/aws"
	azurecluster "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/cluster/azure"
	configpolicy "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/cluster/configpolicy"
	configtemplate "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/cluster/configtemplate"
	customcloud "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/cluster/customcloud"
	edgenative "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/cluster/edgenative"
	edgevsphere "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/cluster/edgevsphere"
	eks "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/cluster/eks"
	gcpcluster "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/cluster/gcp"
	gke "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/cluster/gke"
	group "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/cluster/group"
	maascluster "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/cluster/maas"
	openstackcluster "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/cluster/openstack"
	profilecluster "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/cluster/profile"
	vspherecluster "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/cluster/vsphere"
	setting "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/developer/setting"
	policy "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/password/policy"
	settingplatform "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/platform/setting"
	dnsmap "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/privatecloudgateway/dnsmap"
	ippool "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/privatecloudgateway/ippool"
	providerconfig "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/providerconfig"
	token "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/registration/token"
	helm "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/registry/helm"
	oci "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/registry/oci"
	limit "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/resource/limit"
	alert "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/spectrocloud/alert"
	appliance "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/spectrocloud/appliance"
	application "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/spectrocloud/application"
	datavolume "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/spectrocloud/datavolume"
	filter "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/spectrocloud/filter"
	macros "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/spectrocloud/macros"
	project "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/spectrocloud/project"
	role "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/spectrocloud/role"
	sso "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/spectrocloud/sso"
	team "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/spectrocloud/team"
	user "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/spectrocloud/user"
	workspace "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/spectrocloud/workspace"
	key "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/ssh/key"
	cluster "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/virtual/cluster"
	machine "github.com/crossplane-contrib/provider-palette/internal/controller/cluster/virtual/machine"
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
		configpolicy.Setup,
		configtemplate.Setup,
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
		setting.Setup,
		policy.Setup,
		settingplatform.Setup,
		dnsmap.Setup,
		ippool.Setup,
		providerconfig.Setup,
		token.Setup,
		helm.Setup,
		oci.Setup,
		limit.Setup,
		alert.Setup,
		appliance.Setup,
		application.Setup,
		datavolume.Setup,
		filter.Setup,
		macros.Setup,
		project.Setup,
		role.Setup,
		sso.Setup,
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

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		deployment.SetupGated,
		profile.SetupGated,
		storagelocation.SetupGated,
		aws.SetupGated,
		azure.SetupGated,
		custom.SetupGated,
		gcp.SetupGated,
		maas.SetupGated,
		openstack.SetupGated,
		vsphere.SetupGated,
		aks.SetupGated,
		awscluster.SetupGated,
		azurecluster.SetupGated,
		configpolicy.SetupGated,
		configtemplate.SetupGated,
		customcloud.SetupGated,
		edgenative.SetupGated,
		edgevsphere.SetupGated,
		eks.SetupGated,
		gcpcluster.SetupGated,
		gke.SetupGated,
		group.SetupGated,
		maascluster.SetupGated,
		openstackcluster.SetupGated,
		profilecluster.SetupGated,
		vspherecluster.SetupGated,
		setting.SetupGated,
		policy.SetupGated,
		settingplatform.SetupGated,
		dnsmap.SetupGated,
		ippool.SetupGated,
		providerconfig.SetupGated,
		token.SetupGated,
		helm.SetupGated,
		oci.SetupGated,
		limit.SetupGated,
		alert.SetupGated,
		appliance.SetupGated,
		application.SetupGated,
		datavolume.SetupGated,
		filter.SetupGated,
		macros.SetupGated,
		project.SetupGated,
		role.SetupGated,
		sso.SetupGated,
		team.SetupGated,
		user.SetupGated,
		workspace.SetupGated,
		key.SetupGated,
		cluster.SetupGated,
		machine.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
