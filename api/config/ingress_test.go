package config

import (
	"bytes"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetIngressConfig(t *testing.T) {
	viper.SetConfigType("yaml")
	testConfig := []byte(`
clusterIssuer: letsencrypt
namespace: test
ingress:
  class: nginx
  customMeta:
      "nginx.ingress.kubernetes.io/force-ssl-redirect": "true"
  service:
    name: example-svc
    port:
      number: 80
  `)
	viper.ReadConfig(bytes.NewBuffer(testConfig))
	ig := GetIngressConfig()

	expectedPortNumber := int32(80)
	assert.EqualValues(t, ig, IngressConfig{
		Namespace:     "test",
		IngressClass:  "nginx",
		ClusterIssuer: "letsencrypt",
		CustomMeta:    map[string]string{"nginx.ingress.kubernetes.io/force-ssl-redirect": "true"},
		Service: ServiceConfig{
			Name: "example-svc",
			Port: PortConfig{
				Number: &expectedPortNumber,
			},
		},
	})
}
