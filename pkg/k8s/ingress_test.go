package k8s

import (
	"github.com/kumojin/k8s-ingress-api/api/config"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/networking/v1"
	"testing"
)

func TestBuildIngressSpec(t *testing.T) {
	kc := config.GetKubernetesConfig()
	kc.InCluster = false

	p := int32(80)
	ic := config.IngressConfig{
		Namespace:     "my-namespace",
		IngressClass:  "ingress-class",
		ClusterIssuer: "cluster-issuer",
		CustomMeta:    nil,
		Service: config.ServiceConfig{
			Name: "my-service",
			Port: config.PortConfig{
				Number: &p,
			},
		},
	}

	c, _ := NewClient(kc, ic)

	ingress := c.BuildIngressSpec("my.host")

	assert.Equal(t, "my.host-ingress", ingress.Name)
	assert.Equal(t, map[string]string{
		"kumojin.com/managed-by":         "k8s-ingress-api",
		"kubernetes.io/ingress.class":    ic.IngressClass,
		"cert-manager.io/cluster-issuer": ic.ClusterIssuer,
	}, ingress.Annotations)

	assert.Equal(t, 1, len(ingress.Spec.Rules))
	assert.Equal(t, "my.host", ingress.Spec.Rules[0].Host)
	assert.Equal(t, 1, len(ingress.Spec.Rules[0].IngressRuleValue.HTTP.Paths))
	assert.Equal(t, "/", ingress.Spec.Rules[0].IngressRuleValue.HTTP.Paths[0].Path)
	assert.Equal(t, v1.PathTypeImplementationSpecific, *ingress.Spec.Rules[0].IngressRuleValue.HTTP.Paths[0].PathType)
	assert.Equal(t, p, ingress.Spec.Rules[0].IngressRuleValue.HTTP.Paths[0].Backend.Service.Port.Number)

}

func TestBuildIngressSpecWithPortName(t *testing.T) {
	kc := config.GetKubernetesConfig()
	kc.InCluster = false

	p := "myPort"
	ic := config.IngressConfig{
		Namespace:     "my-namespace",
		IngressClass:  "ingress-class",
		ClusterIssuer: "cluster-issuer",
		CustomMeta:    nil,
		Service: config.ServiceConfig{
			Name: "my-service",
			Port: config.PortConfig{
				Name: &p,
			},
		},
	}
	c, _ := NewClient(kc, ic)

	ingress := c.BuildIngressSpec("my.host")

	assert.Equal(t, 1, len(ingress.Spec.Rules))
	assert.Equal(t, 1, len(ingress.Spec.Rules[0].IngressRuleValue.HTTP.Paths))
	assert.Equal(t, p, ingress.Spec.Rules[0].IngressRuleValue.HTTP.Paths[0].Backend.Service.Port.Name)
}

func TestBuildIngressSpecWithCustomMeta(t *testing.T) {
	kc := config.GetKubernetesConfig()
	kc.InCluster = false

	p := "myPort"
	ic := config.IngressConfig{
		Namespace:     "my-namespace",
		IngressClass:  "ingress-class",
		ClusterIssuer: "cluster-issuer",
		CustomMeta:    map[string]string{"annotation1": "value1", "annotation2": "value2"},
		Service: config.ServiceConfig{
			Name: "my-service",
			Port: config.PortConfig{
				Name: &p,
			},
		},
	}
	c, _ := NewClient(kc, ic)

	ingress := c.BuildIngressSpec("my.host")

	assert.Equal(t, map[string]string{
		"kumojin.com/managed-by":         "k8s-ingress-api",
		"kubernetes.io/ingress.class":    ic.IngressClass,
		"cert-manager.io/cluster-issuer": ic.ClusterIssuer,
		"annotation1":                    "value1",
		"annotation2":                    "value2",
	}, ingress.Annotations)
}
