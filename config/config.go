package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Config is the filing resource handler config
type Config struct {
	BindAddr                string        `envconfig:"BIND_ADDR"`
	Brokers                 []string      `envconfig:"KAFKA_ADDR"`
	ConsumerGroup           string        `envconfig:"CONSUMER_GROUP"`
	ConsumerTopic           string        `envconfig:"HIERARCHY_BUILT_TOPIC"`
	ElasticSearchAPIURL     string        `envconfig:"ELASTIC_SEARCH_URL"`
	GracefulShutdownTimeout time.Duration `envconfig:"GRACEFUL_SHUTDOWN_TIMEOUT"`
	HealthcheckInterval     time.Duration `envconfig:"HEALTHCHECK_INTERVAL"`
	HealthcheckTimeout      time.Duration `envconfig:"HEALTHCHECK_TIMEOUT"`
	HierarchyAPIURL         string        `envconfig:"HIERARCHY_API_URL"`
	KafkaMaxBytes           string        `envconfig:"KAFKA_MAX_BYTES"`
	MaxRetries              int           `envconfig:"REQUEST_MAX_RETRIES"`
	ProducerTopic           string        `envconfig:"PRODUCER_TOPIC"`
}

var cfg *Config

// Get configures the application and returns the configuration
func Get() (*Config, error) {
	if cfg != nil {
		return cfg, nil
	}

	cfg = &Config{
		BindAddr:                ":22900",
		Brokers:                 []string{"localhost:9092"},
		ConsumerGroup:           "dp-search-builder",
		ConsumerTopic:           "hierarchy-built",
		ElasticSearchAPIURL:     "http://localhost:9200",
		GracefulShutdownTimeout: 5 * time.Second,
		HealthcheckInterval:     time.Minute,
		HealthcheckTimeout:      2 * time.Second,
		HierarchyAPIURL:         "http://localhost:22600",
		KafkaMaxBytes:           "2000000",
		MaxRetries:              3,
		ProducerTopic:           "search-built",
	}

	return cfg, envconfig.Process("", cfg)
}
