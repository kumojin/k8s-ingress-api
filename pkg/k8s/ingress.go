package k8s

type IngressCreateOptions struct {
	Name        string
	Host        string
	TargetHost  string
	ServiceName string
	ServicePort string
}

type IngressCreateRuleOptions struct {
	Host        string
	ServiceName string
	ServicePort string
}

func IngressBuildCreateConfig(options *IngressCreateOptions) map[string]interface{} {

	hostRule := IngressBuildCreateRuleConfig(&IngressCreateRuleOptions{
		Host:        options.Host,
		ServiceName: options.ServiceName,
		ServicePort: options.ServicePort,
	})

	targetHostRule := IngressBuildCreateRuleConfig(&IngressCreateRuleOptions{
		Host:        options.TargetHost,
		ServiceName: options.ServiceName,
		ServicePort: options.ServicePort,
	})

	return map[string]interface{}{
		"apiVersion": "networking.k8s.io/v1beta1",
		"Kind":       "Ingress",
		"metadata": map[string]interface{}{
			"name":                           options.Name,
			"kubernetes.io/ingress.class":    "nginx",
			"cert-manager.io/cluster-issuer": "letsencrypt",
			"nginx.ingress.kubernetes.io/force-ssl-redirect": "true",
		},
		"spec": map[string]interface{}{
			"tls": []map[string]interface{}{
				{
					"hosts":      []string{options.TargetHost},
					"secretName": options.Name,
				},
			},
			"rules": []map[string]interface{}{
				hostRule,
				targetHostRule,
			},
		},
	}
}

func IngressBuildCreateRuleConfig(options *IngressCreateRuleOptions) map[string]interface{} {
	return map[string]interface{}{
		"host": options.Host,
		"http": map[string]interface{}{
			"paths": []map[string]interface{}{
				{
					"backend": map[string]interface{}{
						"serviceName": options.ServiceName,
						"servicePort": options.ServicePort,
					},
				},
			},
		},
	}
}
