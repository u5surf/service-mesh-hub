package cert_signer_test

import (
	"context"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rotisserie/eris"
	. "github.com/solo-io/go-utils/testutils"
	core_types "github.com/solo-io/mesh-projects/pkg/api/core.zephyr.solo.io/v1alpha1/types"
	"github.com/solo-io/mesh-projects/pkg/api/networking.zephyr.solo.io/v1alpha1"
	v1alpha1_types "github.com/solo-io/mesh-projects/pkg/api/networking.zephyr.solo.io/v1alpha1/types"
	mock_kubernetes_core "github.com/solo-io/mesh-projects/pkg/clients/kubernetes/core/mocks"
	mock_zephyr_networking "github.com/solo-io/mesh-projects/pkg/clients/zephyr/networking/mocks"
	"github.com/solo-io/mesh-projects/pkg/env"
	cert_secrets "github.com/solo-io/mesh-projects/pkg/security/secrets"
	cert_signer "github.com/solo-io/mesh-projects/services/mesh-networking/pkg/security/cert-signer"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("mesh group cert client", func() {
	var (
		ctrl                *gomock.Controller
		ctx                 context.Context
		secretClient        *mock_kubernetes_core.MockSecretsClient
		meshGroupClient     *mock_zephyr_networking.MockMeshGroupClient
		meshGroupCertCLient cert_signer.MeshGroupCertClient
		testErr             = eris.New("hello")
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		ctx = context.TODO()
		secretClient = mock_kubernetes_core.NewMockSecretsClient(ctrl)
		meshGroupClient = mock_zephyr_networking.NewMockMeshGroupClient(ctrl)
		meshGroupCertCLient = cert_signer.NewMeshGroupCertClient(secretClient, meshGroupClient)
	})

	It("will fail if meshgroup cannot be found", func() {
		meshRef := &core_types.ResourceRef{
			Name:      "name",
			Namespace: "namespace",
		}
		meshGroupClient.EXPECT().Get(ctx, meshRef.Name, meshRef.Namespace).Return(nil, testErr)
		_, err := meshGroupCertCLient.GetRootCaBundle(ctx, meshRef)
		Expect(err).To(HaveOccurred())
		Expect(err).To(HaveInErrorChain(testErr))
	})

	It("will use default trust bundle if mg one not set", func() {
		meshRef := &core_types.ResourceRef{
			Name:      "name",
			Namespace: "namespace",
		}
		mg := &v1alpha1.MeshGroup{
			ObjectMeta: metav1.ObjectMeta{
				Name:      meshRef.Name,
				Namespace: meshRef.Namespace,
			},
			Spec: v1alpha1_types.MeshGroupSpec{},
		}
		meshGroupClient.EXPECT().Get(ctx, meshRef.Name, meshRef.Namespace).Return(mg, nil)
		secretClient.EXPECT().Get(ctx, cert_signer.DefaultRootCaName(mg), env.DefaultWriteNamespace).
			Return(nil, testErr)
		_, err := meshGroupCertCLient.GetRootCaBundle(ctx, meshRef)
		Expect(err).To(HaveOccurred())
		Expect(err).To(HaveInErrorChain(testErr))
	})

	It("will use trust bundle in mg if set", func() {
		meshRef := &core_types.ResourceRef{
			Name:      "name",
			Namespace: "namespace",
		}
		mg := &v1alpha1.MeshGroup{
			ObjectMeta: metav1.ObjectMeta{
				Name:      meshRef.Name,
				Namespace: meshRef.Namespace,
			},
			Spec: v1alpha1_types.MeshGroupSpec{
				TrustBundleRef: &core_types.ResourceRef{
					Name:      "tb_name",
					Namespace: "tb_namespace",
				},
			},
		}
		meshGroupClient.EXPECT().Get(ctx, meshRef.Name, meshRef.Namespace).Return(mg, nil)
		secretClient.EXPECT().Get(ctx, mg.Spec.TrustBundleRef.GetName(), mg.Spec.TrustBundleRef.GetNamespace()).
			Return(nil, testErr)
		_, err := meshGroupCertCLient.GetRootCaBundle(ctx, meshRef)
		Expect(err).To(HaveOccurred())
		Expect(err).To(HaveInErrorChain(testErr))
	})

	It("will return proper CA data", func() {
		meshRef := &core_types.ResourceRef{
			Name:      "name",
			Namespace: "namespace",
		}
		mg := &v1alpha1.MeshGroup{
			ObjectMeta: metav1.ObjectMeta{
				Name:      meshRef.Name,
				Namespace: meshRef.Namespace,
			},
			Spec: v1alpha1_types.MeshGroupSpec{
				TrustBundleRef: &core_types.ResourceRef{
					Name:      "tb_name",
					Namespace: "tb_namespace",
				},
			},
		}
		matchData := &cert_secrets.RootCaData{
			CertAndKeyData: cert_secrets.CertAndKeyData{
				CertChain:  []byte("cert_chain"),
				PrivateKey: []byte("private_key"),
				RootCert:   []byte("root_cert"),
			},
			CaCert:       []byte("ca_cert"),
			CaPrivateKey: []byte("ca_key"),
		}
		matchSecret := &corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      mg.Spec.TrustBundleRef.GetName(),
				Namespace: mg.Spec.TrustBundleRef.GetNamespace(),
			},
			Data: map[string][]byte{
				cert_secrets.RootCertID:     matchData.RootCert,
				cert_secrets.PrivateKeyID:   matchData.PrivateKey,
				cert_secrets.CertChainID:    matchData.CertChain,
				cert_secrets.CaPrivateKeyID: matchData.CaPrivateKey,
				cert_secrets.CaCertID:       matchData.CaCert,
			},
			Type: cert_secrets.RootCertSecretType,
		}
		meshGroupClient.EXPECT().Get(ctx, meshRef.Name, meshRef.Namespace).Return(mg, nil)
		secretClient.EXPECT().Get(ctx, mg.Spec.TrustBundleRef.GetName(), mg.Spec.TrustBundleRef.GetNamespace()).
			Return(matchSecret, nil)
		data, err := meshGroupCertCLient.GetRootCaBundle(ctx, meshRef)
		Expect(err).NotTo(HaveOccurred())
		Expect(data).To(Equal(matchData))
	})
})
