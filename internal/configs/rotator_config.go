package rotatorconfig

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server   Server   `yaml:"server"`
	DSN      Dsn      `yaml:"dsn"`
	RmqCreds RmqCreds `yaml:"rmq"`
}

type Server struct {
	Host     string `yaml:"host"`
	GrpcPort string `yaml:"grpcPort"`
	LogLevel string `yaml:"logLevel"`
}

type RmqCreds struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type Dsn struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	Ssl      string `yaml:"ssl"`
}

func GetConfig(path string) (*Config, error) {
	configYaml, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	err = yaml.Unmarshal(configYaml, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
