package config

import "github.com/spf13/viper"

type portConfig struct {
	Name   *string
	Number *int32
}
type serviceConfig struct {
	Name string
	Port portConfig
}
type IngressConfig struct {
	Namespace     string
	IngressClass  string
	ClusterIssuer string
	CustomMeta    map[string]string
	Service       serviceConfig
}

func GetIngressConfig() IngressConfig {
	viper.SetDefault("namespace", "default")
	viper.SetDefault("clusterIssuer", "letsencrypt")
	viper.SetDefault("ingress.class", "nginx")

	c := IngressConfig{
		Namespace:     viper.GetString("namespace"),
		IngressClass:  viper.GetString("ingress.class"),
		ClusterIssuer: viper.GetString("clusterIssuer"),
		CustomMeta:    viper.GetStringMapString("ingress.customMeta"),
		Service: serviceConfig{
			Name: viper.GetString("ingress.service.name"),
			Port: portConfig{},
		},
	}

	if viper.IsSet("ingress.service.port.number") {
		v := viper.GetInt32("ingress.service.port.number")
		c.Service.Port.Number = &v
	}

	if viper.IsSet("ingress.service.port.name") {
		v := viper.GetString("ingress.service.port.name")
		c.Service.Port.Name = &v
	}

	return c
}
