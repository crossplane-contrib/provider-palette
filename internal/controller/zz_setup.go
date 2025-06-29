/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	role "github.com/crossplane-contrib/provider-palette/internal/controller/accesscontrol/role"
	team "github.com/crossplane-contrib/provider-palette/internal/controller/accesscontrol/team"
	user "github.com/crossplane-contrib/provider-palette/internal/controller/accesscontrol/user"
	application "github.com/crossplane-contrib/provider-palette/internal/controller/application/application"
	applicationprofile "github.com/crossplane-contrib/provider-palette/internal/controller/application/applicationprofile"
	awscloudaccount "github.com/crossplane-contrib/provider-palette/internal/controller/aws/awscloudaccount"
	awscluster "github.com/crossplane-contrib/provider-palette/internal/controller/aws/awscluster"
	ekscluster "github.com/crossplane-contrib/provider-palette/internal/controller/aws/ekscluster"
	akscluster "github.com/crossplane-contrib/provider-palette/internal/controller/azure/akscluster"
	azurecloudaccount "github.com/crossplane-contrib/provider-palette/internal/controller/azure/azurecloudaccount"
	azurecluster "github.com/crossplane-contrib/provider-palette/internal/controller/azure/azurecluster"
	backupstoragelocation "github.com/crossplane-contrib/provider-palette/internal/controller/backup/backupstoragelocation"
	clustergroup "github.com/crossplane-contrib/provider-palette/internal/controller/clustergroup/clustergroup"
	clusterprofile "github.com/crossplane-contrib/provider-palette/internal/controller/clusterprofile/clusterprofile"
	customcloudaccount "github.com/crossplane-contrib/provider-palette/internal/controller/customcloud/customcloudaccount"
	customcluster "github.com/crossplane-contrib/provider-palette/internal/controller/customcloud/customcluster"
	appliance "github.com/crossplane-contrib/provider-palette/internal/controller/edge/appliance"
	edgecluster "github.com/crossplane-contrib/provider-palette/internal/controller/edge/edgecluster"
	gcpcloudaccount "github.com/crossplane-contrib/provider-palette/internal/controller/gcp/gcpcloudaccount"
	gcpcluster "github.com/crossplane-contrib/provider-palette/internal/controller/gcp/gcpcluster"
	gkecluster "github.com/crossplane-contrib/provider-palette/internal/controller/gcp/gkecluster"
	maascloudaccount "github.com/crossplane-contrib/provider-palette/internal/controller/maas/maascloudaccount"
	maascluster "github.com/crossplane-contrib/provider-palette/internal/controller/maas/maascluster"
	alert "github.com/crossplane-contrib/provider-palette/internal/controller/monitoring/alert"
	filter "github.com/crossplane-contrib/provider-palette/internal/controller/monitoring/filter"
	openstackcloudaccount "github.com/crossplane-contrib/provider-palette/internal/controller/openstack/openstackcloudaccount"
	openstackcluster "github.com/crossplane-contrib/provider-palette/internal/controller/openstack/openstackcluster"
	platformsetting "github.com/crossplane-contrib/provider-palette/internal/controller/platform/platformsetting"
	resourcelimit "github.com/crossplane-contrib/provider-palette/internal/controller/platform/resourcelimit"
	dnsmap "github.com/crossplane-contrib/provider-palette/internal/controller/privatecloudgateway/dnsmap"
	ippool "github.com/crossplane-contrib/provider-palette/internal/controller/privatecloudgateway/ippool"
	providerconfig "github.com/crossplane-contrib/provider-palette/internal/controller/providerconfig"
	helmregistry "github.com/crossplane-contrib/provider-palette/internal/controller/registry/helmregistry"
	ociregistry "github.com/crossplane-contrib/provider-palette/internal/controller/registry/ociregistry"
	developersetting "github.com/crossplane-contrib/provider-palette/internal/controller/security/developersetting"
	passwordpolicy "github.com/crossplane-contrib/provider-palette/internal/controller/security/passwordpolicy"
	registrationtoken "github.com/crossplane-contrib/provider-palette/internal/controller/security/registrationtoken"
	sshkey "github.com/crossplane-contrib/provider-palette/internal/controller/security/sshkey"
	sso "github.com/crossplane-contrib/provider-palette/internal/controller/security/sso"
	addondeployment "github.com/crossplane-contrib/provider-palette/internal/controller/spectrocloud/addondeployment"
	macro "github.com/crossplane-contrib/provider-palette/internal/controller/spectrocloud/macro"
	macros "github.com/crossplane-contrib/provider-palette/internal/controller/spectrocloud/macros"
	project "github.com/crossplane-contrib/provider-palette/internal/controller/spectrocloud/project"
	workspace "github.com/crossplane-contrib/provider-palette/internal/controller/spectrocloud/workspace"
	virtualcluster "github.com/crossplane-contrib/provider-palette/internal/controller/virtualcluster/virtualcluster"
	datavolume "github.com/crossplane-contrib/provider-palette/internal/controller/virtualmachine/datavolume"
	virtualmachine "github.com/crossplane-contrib/provider-palette/internal/controller/virtualmachine/virtualmachine"
	edgevspherecluster "github.com/crossplane-contrib/provider-palette/internal/controller/vsphere/edgevspherecluster"
	vspherecloudaccount "github.com/crossplane-contrib/provider-palette/internal/controller/vsphere/vspherecloudaccount"
	vspherecluster "github.com/crossplane-contrib/provider-palette/internal/controller/vsphere/vspherecluster"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		role.Setup,
		team.Setup,
		user.Setup,
		application.Setup,
		applicationprofile.Setup,
		awscloudaccount.Setup,
		awscluster.Setup,
		ekscluster.Setup,
		akscluster.Setup,
		azurecloudaccount.Setup,
		azurecluster.Setup,
		backupstoragelocation.Setup,
		clustergroup.Setup,
		clusterprofile.Setup,
		customcloudaccount.Setup,
		customcluster.Setup,
		appliance.Setup,
		edgecluster.Setup,
		gcpcloudaccount.Setup,
		gcpcluster.Setup,
		gkecluster.Setup,
		maascloudaccount.Setup,
		maascluster.Setup,
		alert.Setup,
		filter.Setup,
		openstackcloudaccount.Setup,
		openstackcluster.Setup,
		platformsetting.Setup,
		resourcelimit.Setup,
		dnsmap.Setup,
		ippool.Setup,
		providerconfig.Setup,
		helmregistry.Setup,
		ociregistry.Setup,
		developersetting.Setup,
		passwordpolicy.Setup,
		registrationtoken.Setup,
		sshkey.Setup,
		sso.Setup,
		addondeployment.Setup,
		macro.Setup,
		macros.Setup,
		project.Setup,
		workspace.Setup,
		virtualcluster.Setup,
		datavolume.Setup,
		virtualmachine.Setup,
		edgevspherecluster.Setup,
		vspherecloudaccount.Setup,
		vspherecluster.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
