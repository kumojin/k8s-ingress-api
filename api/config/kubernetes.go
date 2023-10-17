package config

import (
	"github.com/spf13/viper"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

type KubernetesConfig struct {
	InCluster  bool
	Kubeconfig string
}

func GetKubernetesConfig() KubernetesConfig {
	viper.SetDefault("kubernetes.inCluster", true)
	viper.SetDefault("kubeconfig", filepath.Join(homedir.HomeDir(), ".kube", "config"))

	c := KubernetesConfig{
		InCluster:  viper.GetBool("kubernetes.inCluster"),
		Kubeconfig: viper.GetString("kubeconfig"),
	}

	return c
}
