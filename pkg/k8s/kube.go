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
	if kc.InCluster {
		client, err := getInClusterClient()
		if err != nil {
			return nil, err
		}
		c := Client{config: ig, Client: client}
		return &c, nil
	}

	config, err := clientcmd.BuildConfigFromFlags("", kc.Kubeconfig)
	if err != nil {
		return nil, err
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	c := Client{config: ig, Client: client}
	return &c, nil
}

func getInClusterClient() (*kubernetes.Clientset, error) {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	return kubernetes.NewForConfig(config)
}
