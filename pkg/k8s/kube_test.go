package k8s

//
//import (
//	"context"
//	"os"
//	"testing"
//	"time"
//
//	"github.com/stretchr/testify/assert"
//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"k8s.io/client-go/kubernetes"
//)
//
//func TestDefaultKubeConfigPath(t *testing.T) {
//	path := DefaultKubeConfigPath()
//	assert.Equal(t, os.Getenv("HOME")+"/.kube/config", path)
//}
//
//func TestBuildKubeConfig(t *testing.T) {
//	config, err := BuildKubeConfig(DefaultKubeConfigPath())
//	assert.Empty(t, err)
//	assert.NotEmpty(t, config)
//	assert.Equal(t, "https://127.0.0.1:32772", config.Host)
//}
//
//func TestBuildDefaultKubeConfig(t *testing.T) {
//	config, err := BuildDefaultKubeConfig()
//	assert.Empty(t, err)
//	assert.NotEmpty(t, config)
//	assert.Equal(t, "https://127.0.0.1:32772", config.Host)
//}
//
//func TestLocalConnection(t *testing.T) {
//	config, err := BuildDefaultKubeConfig()
//	assert.Empty(t, err)
//
//	k8sClient, err := kubernetes.NewForConfig(config)
//	assert.Empty(t, err)
//	assert.NotEmpty(t, k8sClient)
//
//	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
//	defer cancel()
//
//	coreApi := k8sClient.CoreV1()
//
//	nodes, err := coreApi.Nodes().List(ctx, metav1.ListOptions{FieldSelector: "metadata.name=minikube"})
//	assert.Empty(t, err)
//	assert.Equal(t, 1, len(nodes.Items))
//	assert.Equal(t, "minikube", nodes.Items[0].ObjectMeta.Name)
//}
