package config

import (
	"strings"

	"github.com/spf13/viper"
)

// Config holds application configuration values.
type Config struct {
	Port        string
	KafkaBroker string
	KafkaTopic  string
	LogTag      string
}

// Load reads configuration from environment variables (and defaults).
// It uses Viper so you can also add file-based config later if needed.
func Load() *Config {
	v := viper.New()

	// sensible defaults
	v.SetDefault("PORT", "8081")
	v.SetDefault("KAFKA_BROKER", "localhost:9092")
	v.SetDefault("KAFKA_TOPIC", "raw-logs")
	v.SetDefault("LOG_TAG", "wazuh-dc")

	// allow environment variables to override
	v.AutomaticEnv()
	// map environment variable names like LOG_TAG or kafka.broker to same keys
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	return &Config{
		Port:        v.GetString("PORT"),
		KafkaBroker: v.GetString("KAFKA_BROKER"),
		KafkaTopic:  v.GetString("KAFKA_TOPIC"),
		LogTag:      v.GetString("LOG_TAG"),
	}
}
