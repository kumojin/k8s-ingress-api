package k8s

import (
	"k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type IngressCreateTrimOptions struct {
	Name        string
	Host        string
	TargetHost  string
	ServicePort int
}

func BuildIngressCreateConfig(options *IngressCreateTrimOptions) *v1beta1.Ingress {
	return &v1beta1.Ingress{
		TypeMeta: metav1.TypeMeta{
			APIVersion: ApiVersion,
			Kind:       "Ingress",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: options.Name,
		},
		Spec: v1beta1.IngressSpec{
			TLS: []v1beta1.IngressTLS{
				{
					Hosts:      []string{options.TargetHost},
					SecretName: options.Name + "--" + options.TargetHost,
				},
			},
			Rules: []v1beta1.IngressRule{
				{
					Host: options.Host,
					IngressRuleValue: v1beta1.IngressRuleValue{
						HTTP: &v1beta1.HTTPIngressRuleValue{
							Paths: []v1beta1.HTTPIngressPath{
								{
									Backend: v1beta1.IngressBackend{
										ServiceName: options.Name,
										ServicePort: intstr.FromInt(options.ServicePort),
									},
								},
								{
									Backend: v1beta1.IngressBackend{
										ServiceName: options.Name,
										ServicePort: intstr.FromInt(options.ServicePort),
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
