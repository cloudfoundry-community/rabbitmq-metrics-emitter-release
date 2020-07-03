package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	RMQManagement     RMQManagement     `yaml:"rmq_management"`
	LoggregatorConfig LoggregatorConfig `yaml:"loggregator"`
	CfConfig          CfConfig          `yaml:"cloud_foundry"`
}

type RMQManagement struct {
	Endpoint string `yaml:"endpoint"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
type CfConfig struct {
	ApiUrl            string `yaml:"api_url"`
	SkipSslValidation bool   `yaml:"skip_ssl_validation"`
	Username          string `yaml:"username"`
	Password          string `yaml:"password"`
}

type LoggregatorConfig struct {
	MetronAddress string   `yaml:"metron_address"`
	TLS           TLSCerts `yaml:"tls"`
}

type TLSCerts struct {
	KeyFile    string `yaml:"key_file" json:"keyFile"`
	CertFile   string `yaml:"cert_file" json:"certFile"`
	CACertFile string `yaml:"ca_file" json:"caCertFile"`
}

func LoadConfig(path string) (Config, error) {
	var config Config

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return config, err
	}

	return config, yaml.Unmarshal(yamlFile, &config)

}
