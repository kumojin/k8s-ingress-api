package k8s

import (
	v1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type IngressCreateTrimOptions struct {
	Name              string `json:"name"`
	Host              string `json:"host"`
	TargetServiceName string `json:"targetServiceName"`
	TargetServicePort int    `json:"targetServicePort"`
}

func BuildIngressCreateConfig(options *IngressCreateTrimOptions) *v1.Ingress {
	service := v1.IngressServiceBackend{
		Name: options.TargetServiceName,
		Port: v1.ServiceBackendPort{
			Number: int32(options.TargetServicePort),
		},
	}

	pathType := v1.PathTypeImplementationSpecific

	return &v1.Ingress{
		TypeMeta: metav1.TypeMeta{
			APIVersion: ApiVersion,
			Kind:       "Ingress",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: options.Name + "-" + options.Host,
			Annotations: map[string]string{
				"kubernetes.io/ingress.class":                    "nginx",
				"cert-manager.io/cluster-issuer":                 "letsencrypt",
				"nginx.ingress.kubernetes.io/force-ssl-redirect": "true",
			},
		},
		Spec: v1.IngressSpec{
			TLS: []v1.IngressTLS{
				{
					Hosts:      []string{options.Host},
					SecretName: options.Name + "-" + options.Host,
				},
			},
			Rules: []v1.IngressRule{
				{
					Host: options.Host,
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
