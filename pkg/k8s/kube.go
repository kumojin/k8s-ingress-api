package k8s

import (
	"github.com/kumojin/k8s-ingress-api/api/config"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Client struct {
	config config.IngressConfig
	Client *kubernetes.Clientset
}

func NewClient(kc config.KubernetesConfig, ig config.IngressConfig) (*Client, error) {
	var client *kubernetes.Clientset
	var err error
	if kc.InCluster {
		client, err = getInClusterClient()
	} else {
		client, err = getExternalClusterClient(kc)
	}

	if err != nil {
		return nil, err
	}

	c := Client{config: ig, Client: client}
	return &c, nil
}

// getInClusterClient creates client from in cluster config
func getInClusterClient() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	return kubernetes.NewForConfig(config)
}

// getExternalClusterClient creates client from kubeconfig file
func getExternalClusterClient(kc config.KubernetesConfig) (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", kc.Kubeconfig)
	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(config)
}
