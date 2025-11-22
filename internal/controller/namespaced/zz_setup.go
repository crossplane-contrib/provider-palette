/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	deployment "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/addon/deployment"
	profile "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/application/profile"
	storagelocation "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/backup/storagelocation"
	aws "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/cloudaccount/aws"
	azure "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/cloudaccount/azure"
	custom "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/cloudaccount/custom"
	gcp "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/cloudaccount/gcp"
	maas "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/cloudaccount/maas"
	openstack "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/cloudaccount/openstack"
	vsphere "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/cloudaccount/vsphere"
	aks "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/cluster/aks"
	awscluster "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/cluster/aws"
	azurecluster "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/cluster/azure"
	configpolicy "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/cluster/configpolicy"
	configtemplate "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/cluster/configtemplate"
	customcloud "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/cluster/customcloud"
	edgenative "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/cluster/edgenative"
	edgevsphere "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/cluster/edgevsphere"
	eks "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/cluster/eks"
	gcpcluster "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/cluster/gcp"
	gke "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/cluster/gke"
	group "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/cluster/group"
	maascluster "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/cluster/maas"
	openstackcluster "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/cluster/openstack"
	profilecluster "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/cluster/profile"
	vspherecluster "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/cluster/vsphere"
	setting "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/developer/setting"
	policy "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/password/policy"
	settingplatform "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/platform/setting"
	dnsmap "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/privatecloudgateway/dnsmap"
	ippool "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/privatecloudgateway/ippool"
	providerconfig "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/providerconfig"
	token "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/registration/token"
	helm "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/registry/helm"
	oci "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/registry/oci"
	limit "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/resource/limit"
	alert "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/spectrocloud/alert"
	appliance "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/spectrocloud/appliance"
	application "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/spectrocloud/application"
	datavolume "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/spectrocloud/datavolume"
	filter "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/spectrocloud/filter"
	macros "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/spectrocloud/macros"
	project "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/spectrocloud/project"
	role "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/spectrocloud/role"
	sso "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/spectrocloud/sso"
	team "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/spectrocloud/team"
	user "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/spectrocloud/user"
	workspace "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/spectrocloud/workspace"
	key "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/ssh/key"
	cluster "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/virtual/cluster"
	machine "github.com/crossplane-contrib/provider-palette/internal/controller/namespaced/virtual/machine"
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
