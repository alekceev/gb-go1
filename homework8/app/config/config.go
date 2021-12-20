package config

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/kelseyhightower/envconfig"
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
	Port        string `envconfig:"PORT" default:"8080" required:"true"`
	DbURL       Url    `envconfig:"DB_URL" default:"postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable"`
	JaegerURL   Url    `envconfig:"JAEGER_URL" default:"http://jaeger:16686"`
	SentryURL   Url    `envconfig:"SENTRY_URL" default:"http://sentry:9000"`
	KafkaBroker Url    `envconfig:"KAFKA_BROKER" default:"kafka:9092"`
	SomeAppId   string `envconfig:"SOME_APP_ID" default:"testid"`
	SomeAppKey  string `envconfig:"SOME_APP_KEY" default:"testkey"`
}

// init config
func Get() (Config, error) {
	config := Config{}

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
