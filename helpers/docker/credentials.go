package docker_helpers

type DockerCredentials struct {
	Host                   *string  `toml:"host" json:"host" long:"host" env:"DOCKER_HOST" description:"Docker daemon address"`
	CertPath               *string  `toml:"tls_cert_path" json:"tls_cert_path" long:"cert-path" env:"DOCKER_CERT_PATH" description:"Certificate path"`
	TLSVerify              *bool    `toml:"tls_verify" json:"tls_verify" long:"tlsverify" env:"DOCKER_TLS_VERIFY" description:"Use TLS and verify the remote"`
}

