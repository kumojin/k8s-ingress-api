package config

import "github.com/spf13/viper"

type PortConfig struct {
	Name   *string
	Number *int32
}

type ServiceConfig struct {
	Name string
	Port PortConfig
}

type IngressConfig struct {
	Namespace     string
	IngressClass  string
	ClusterIssuer string
	CustomMeta    map[string]string
	Service       ServiceConfig
}

func GetIngressConfig() IngressConfig {
	viper.SetDefault("namespace", "default")
	viper.SetDefault("clusterIssuer", "letsencrypt")
	viper.SetDefault("ingress.class", "nginx")

	if !viper.IsSet("ingress.service.name") {
		panic("ingress.service.name config is missing")
	}

	if !viper.IsSet("ingress.service.port.number") && !viper.IsSet("ingress.service.port.name") {
		panic("at least ingress.service.port.number or ingress.service.port.name config should be set")
	}

	c := IngressConfig{
		Namespace:     viper.GetString("namespace"),
		IngressClass:  viper.GetString("ingress.class"),
		ClusterIssuer: viper.GetString("clusterIssuer"),
		CustomMeta:    viper.GetStringMapString("ingress.customMeta"),
		Service: ServiceConfig{
			Name: viper.GetString("ingress.service.name"),
			Port: PortConfig{},
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
