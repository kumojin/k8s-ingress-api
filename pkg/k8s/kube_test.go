package k8s

import (
  "context"
  "errors"
  "github.com/kumojin/k8s-ingress-api/api/config"
  "github.com/stretchr/testify/assert"
  metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
  "os"
  "testing"
  "time"
)

func skipCI(t *testing.T) {
  kc := config.GetKubernetesConfig()
  if _, err := os.Stat(kc.Kubeconfig); errors.Is(err, os.ErrNotExist) {
    t.Skip("Skipping testing in CI environment")
  }
}

func TestNewClientFromExternal(t *testing.T) {
  skipCI(t)

  kc := config.GetKubernetesConfig()
  kc.InCluster = false
  _, err := NewClient(kc, config.IngressConfig{})

  assert.Empty(t, err)
}

func TestLocalConnection(t *testing.T) {
  skipCI(t)

  kc := config.GetKubernetesConfig()
  kc.InCluster = false
  c, err := NewClient(kc, config.IngressConfig{})
  assert.Empty(t, err)

  ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
  defer cancel()

  coreApi := c.Client.CoreV1()
  nodes, err := coreApi.Nodes().List(ctx, metav1.ListOptions{FieldSelector: "metadata.name=docker-desktop"})
  assert.Empty(t, err)
  assert.Equal(t, 1, len(nodes.Items))
}
