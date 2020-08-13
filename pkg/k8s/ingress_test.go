package k8s

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func TestBuildIngressCreateConfig(t *testing.T) {
	createOpts := &IngressCreateTrimOptions{
		Name:        "stent-urlshortener",
		Host:        "linkedin.cofomo.com",
		TargetHost:  "stnt.co",
		ServicePort: 80,
	}
	create := BuildIngressCreateConfig(createOpts)
	assert.Equal(t, "Ingress", create.TypeMeta.Kind)
	assert.Equal(t, createOpts.Name, create.ObjectMeta.Name)
	//TODO continue
}

func TestCreateIngress(t *testing.T) {
	ingressCreate := BuildIngressCreateConfig(&IngressCreateTrimOptions{
		Name:        "stent-urlshortener",
		Host:        "linkedin.cofomo.com",
		TargetHost:  "stnt.co",
		ServicePort: 80,
	})

	config, err := BuildDefaultKubeConfig()
	assert.Empty(t, err)

	k8sClient, err := kubernetes.NewForConfig(config)
	assert.Empty(t, err)

	ingresses := k8sClient.ExtensionsV1beta1().Ingresses(testNS)

	// create the ingress
	ingress, err := ingresses.Create(context.Background(), ingressCreate, metav1.CreateOptions{})
	assert.Empty(t, err)
	assert.NotEmpty(t, ingress)

	// then, delete it
	err = ingresses.Delete(context.Background(), ingressCreate.Name, metav1.DeleteOptions{})
	assert.Empty(t, err)
}
