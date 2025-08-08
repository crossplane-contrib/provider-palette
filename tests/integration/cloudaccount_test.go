package integration

import (
	v1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	// cloudaccountv1alpha1 "github.com/crossplane-contrib/provider-palette/apis/cloudaccount/v1alpha1"
	"github.com/crossplane-contrib/provider-palette/internal/utils"
)

var _ = Describe("Cloud Account Tests", Ordered, func() {

	Describe("Creating an AWS cloud account", func() {
		It("should succeed", func() {
			awsCloudAccount := &cloudaccountv1alpha1.Aws{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "aws-cloud-account",
					Namespace: "default",
				},
				Spec: cloudaccountv1alpha1.AwsSpec{
					ResourceSpec: v1.ResourceSpec{
						ProviderConfigReference: &v1.Reference{
							Name: "provider-palette-config",
						},
					},
					ForProvider: cloudaccountv1alpha1.AwsParameters{
						Name: utils.Ptr("aws-cloud-account"),
						Type: utils.Ptr("secret"),
					},
				},
			}
			err := kclient.Create(tc, awsCloudAccount)
			if err != nil {
				_ = kclient.Delete(tc, awsCloudAccount)
			}
			Expect(err).ShouldNot(HaveOccurred())

			Eventually(func() bool {
				if err := kclient.Get(tc, keyFromObj(awsCloudAccount), awsCloudAccount); err != nil {
					return false
				}
				if len(awsCloudAccount.Status.Conditions) == 0 {
					return false
				}
				for _, c := range awsCloudAccount.Status.Conditions {
					logf.Log.Info("AWSCloudAccount Condition", "Type", c.Type, "Status", c.Status, "Reason", c.Reason)
					if c.Status != corev1.ConditionTrue {
						return false
					}
				}
				logf.Log.Info("AWSCloudAccount created successfully")
				return true
			}, timeout, interval).Should(BeTrue(), "failed to create AWS cloud account")
		})
	})
})
