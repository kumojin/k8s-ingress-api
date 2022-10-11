package k8s

import (
	"context"
	v1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

type IngressCreateTrimOptions struct {
	Name              string `json:"name"`
	Host              string `json:"host"`
	TargetServiceName string `json:"targetServiceName"`
	TargetServicePort int    `json:"targetServicePort"`
}

func (c *Client) CreateIngress(host string, dryRun bool) (*v1.Ingress, error) {
	createOpts := metav1.CreateOptions{}
	if dryRun {
		createOpts.DryRun = []string{"All"}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3000*time.Millisecond)
	defer cancel()

	ingressSpec := c.buildIngressSpec(host)

	ingresses := c.Client.NetworkingV1().Ingresses(c.config.Namespace)
	return ingresses.Create(ctx, &ingressSpec, createOpts)
}

func (c *Client) buildIngressSpec(host string) v1.Ingress {
	var pc v1.ServiceBackendPort
	if c.config.Service.Port.Number != nil {
		pc = v1.ServiceBackendPort{Number: *c.config.Service.Port.Number}
	} else {
		pc = v1.ServiceBackendPort{Name: *c.config.Service.Port.Name}
	}

	service := v1.IngressServiceBackend{
		Name: c.config.Service.Name,
		Port: pc,
	}

	pathType := v1.PathTypeImplementationSpecific
	annotations := map[string]string{
		"kumojin.com/managed-by":         "k8s-ingress-api",
		"kubernetes.io/ingress.class":    c.config.IngressClass,
		"cert-manager.io/cluster-issuer": c.config.ClusterIssuer,
	}

	for k, v := range c.config.CustomMeta {
		annotations[k] = v
	}

	return v1.Ingress{
		TypeMeta: metav1.TypeMeta{
			APIVersion: ApiVersion,
			Kind:       "Ingress",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        convertHostToName(host) + "-ingress",
			Annotations: annotations,
		},
		Spec: v1.IngressSpec{
			TLS: []v1.IngressTLS{
				{
					Hosts:      []string{host},
					SecretName: convertHostToName(host) + "-tls",
				},
			},
			Rules: []v1.IngressRule{
				{
					Host: host,
					IngressRuleValue: v1.IngressRuleValue{
						HTTP: &v1.HTTPIngressRuleValue{
							Paths: []v1.HTTPIngressPath{
								{
									Path:     "/",
									PathType: &pathType,
									Backend: v1.IngressBackend{
										Service: &service,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func convertHostToName(host string) string {
	return host
}
