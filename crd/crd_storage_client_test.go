package crd_test

import (
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/solo-io/glue-storage/crd"
	"github.com/solo-io/glue-storage/crd/solo.io/v1"
	. "github.com/solo-io/glue/test/helpers"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	apiexts "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

var _ = Describe("CrdStorageClient", func() {
	var (
		masterUrl, kubeconfigPath string
		mkb                       *MinikubeInstance
		namespace                 string
	)
	BeforeEach(func() {
		namespace = RandString(8)
		mkb = NewMinikube(false, namespace)
		err := mkb.Setup()
		Must(err)
		kubeconfigPath = filepath.Join(os.Getenv("HOME"), ".kube", "config")
		masterUrl, err = mkb.Addr()
		Must(err)
	})
	AfterEach(func() {
		mkb.Teardown()
	})
	Describe("New", func() {
		It("creates a new client without error", func() {
			cfg, err := clientcmd.BuildConfigFromFlags(masterUrl, kubeconfigPath)
			Expect(err).NotTo(HaveOccurred())
			_, err = NewClient(cfg, namespace)
			Expect(err).NotTo(HaveOccurred())
		})
	})
	Describe("Register", func() {
		It("registers the crds", func() {
			cfg, err := clientcmd.BuildConfigFromFlags(masterUrl, kubeconfigPath)
			Expect(err).NotTo(HaveOccurred())
			client, err := NewClient(cfg, namespace)
			Expect(err).NotTo(HaveOccurred())
			err = client.Register()
			Expect(err).NotTo(HaveOccurred())
			apiextClient, err := apiexts.NewForConfig(cfg)
			Expect(err).NotTo(HaveOccurred())
			crds, err := apiextClient.ApiextensionsV1beta1().CustomResourceDefinitions().List(metav1.ListOptions{})
			Expect(err).NotTo(HaveOccurred())
			for _, crdSchema := range v1.KnownCRDs {
				var foundCrd *v1beta1.CustomResourceDefinition
				for _, crd := range crds.Items {
					if crd.Spec.Names.Kind == crdSchema.Kind {
						foundCrd = &crd
						break
					}
				}
				// if crd wasnt found, err
				Expect(foundCrd).NotTo(BeNil())

				Expect(foundCrd.Spec.Version).To(Equal(crdSchema.Version))
				Expect(foundCrd.Spec.Group).To(Equal(crdSchema.Group))
			}
		})
	})
})
