package config

import (
	"github.com/flyteorg/flytestdlib/config"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

//go:generate enumer --type=SecretManagerType --trimprefix=SecretManagerType -json -yaml
//go:generate pflags Config --default-var=DefaultConfig

var (
	DefaultConfig = &Config{
		SecretName:        "flyte-pod-webhook",
		ServiceName:       "flyte-pod-webhook",
		MetricsPrefix:     "flyte:",
		CertDir:           "/etc/webhook/certs",
		ListenPort:        9443,
		SecretManagerType: SecretManagerTypeK8s,
		AWSSecretManagerConfig: AWSSecretManagerConfig{
			SidecarImage: "docker.io/amazon/aws-secrets-manager-secret-sidecar:v0.1.5",
			Resources: corev1.ResourceRequirements{
				Requests: corev1.ResourceList{
					corev1.ResourceMemory: resource.MustParse("500Mi"),
					corev1.ResourceCPU:    resource.MustParse("200m"),
				},
				Limits: corev1.ResourceList{
					corev1.ResourceMemory: resource.MustParse("500Mi"),
					corev1.ResourceCPU:    resource.MustParse("200m"),
				},
			},
		},
	}

	configSection = config.MustRegisterSection("webhook", DefaultConfig)
)

// SecretManagerType defines which secret manager to use.
type SecretManagerType int

const (
	// SecretManagerTypeGlobal defines a global secret manager that can read env vars and mounted secrets to the webhook
	// pod.
	SecretManagerTypeGlobal SecretManagerType = iota

	// SecretManagerTypeK8s defines a secret manager webhook that injects K8s volume mounts to mount K8s secrets.
	SecretManagerTypeK8s

	// SecretManagerTypeAWS defines a secret manager webhook that injects a side car to pull secrets from AWS Secret
	// Manager and mount them to a local file system (in memory) and share that mount with other containers in the pod.
	SecretManagerTypeAWS
)

type Config struct {
	MetricsPrefix          string                 `json:"metrics-prefix" pflag:",An optional prefix for all published metrics."`
	CertDir                string                 `json:"certDir" pflag:",Certificate directory to use to write generated certs. Defaults to /etc/webhook/certs/"`
	ListenPort             int                    `json:"listenPort" pflag:",The port to use to listen to webhook calls. Defaults to 9443"`
	ServiceName            string                 `json:"serviceName" pflag:",The name of the webhook service."`
	SecretName             string                 `json:"secretName" pflag:",Secret name to write generated certs to."`
	SecretManagerType      SecretManagerType      `json:"secretManagerType" pflag:"-,Secret manager type to use if secrets are not found in global secrets."`
	AWSSecretManagerConfig AWSSecretManagerConfig `json:"awsSecretManager" pflag:",AWS Secret Manager config."`
}

type AWSSecretManagerConfig struct {
	SidecarImage string `json:"sidecarImage" pflag:",Specifies the sidecar docker image to use"`
	Resources    corev1.ResourceRequirements
}

func GetConfig() *Config {
	return configSection.GetConfig().(*Config)
}
