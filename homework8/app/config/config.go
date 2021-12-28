package config

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type Url string

func (u *Url) Decode(value string) error {
	_, err := url.ParseRequestURI(value)
	if err != nil {
		return fmt.Errorf("Url error %v", err)
	}
	*u = Url(value)
	return nil
}

type Config struct {
	Port        string `yaml:"port" envconfig:"PORT" default:"8080" required:"true"`
	DbURL       Url    `yaml:"db_url" envconfig:"DB_URL" default:"postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable"`
	JaegerURL   Url    `yaml:"jaeger_url" envconfig:"JAEGER_URL" default:"http://jaeger:16686"`
	SentryURL   Url    `yaml:"sentry_url" envconfig:"SENTRY_URL" default:"http://sentry:9000"`
	KafkaBroker Url    `yaml:"kafka_broker" envconfig:"KAFKA_BROKER" default:"kafka:9092"`
	SomeAppId   string `yaml:"some_app_id" envconfig:"SOME_APP_ID" default:"testid"`
	SomeAppKey  string `yaml:"some_app_key" envconfig:"SOME_APP_KEY" default:"testkey"`
}

// init config
func Get(path string) (Config, error) {
	config := Config{}

	if path != "" {
		return getConfigFromYaml(path)
	}

	err := envconfig.Process("", &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func (c Config) String() string {
	conf, _ := json.MarshalIndent(c, "", "  ")
	return string(conf)
}

func getConfigFromYaml(path string) (Config, error) {
	config := Config{}

	file, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()
	err = yaml.NewDecoder(file).Decode(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
