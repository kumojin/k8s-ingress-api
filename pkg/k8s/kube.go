package k8s

import (
	"path/filepath"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func DefaultKubeConfigPath() string {
	return filepath.Join(homedir.HomeDir(), ".kube", "config")
}

func BuildDefaultKubeConfig() (*rest.Config, error) {
	return BuildKubeConfig(DefaultKubeConfigPath())
}

func BuildKubeConfig(kubeConfig string) (*rest.Config, error) {
	return clientcmd.BuildConfigFromFlags("", kubeConfig)
}
