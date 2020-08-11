package k8s

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestCreateRuleConfig(t *testing.T) {
	rule := IngressBuildCreateRuleConfig(&IngressCreateRuleOptions{
		Host:        "gl.com",
		ServiceName: "some-service-name",
		ServicePort: "80",
	})

	ruleYml, err := yaml.Marshal(&rule)
	assert.Empty(t, err)

	expectedYml, _ := ioutil.ReadFile("./test_data/TestCreateRuleConfig.yml")
	assert.Equal(t, string(expectedYml), string(ruleYml))
}

func TestCreateConfig(t *testing.T) {
	rule := IngressBuildCreateConfig(&IngressCreateOptions{
		Name:        "some-name",
		Host:        "gl.com",
		TargetHost:  "google.com",
		ServiceName: "some-service-name",
		ServicePort: "80",
	})

	ruleYml, err := yaml.Marshal(&rule)
	assert.Empty(t, err)

	expectedYml, _ := ioutil.ReadFile("./test_data/TestCreateConfig.yml")
	assert.Equal(t, string(expectedYml), string(ruleYml))
}
