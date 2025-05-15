package config

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

var appConfig Config

type Config struct {
	Address     string
	DatabaseDNS string
	PrivateKey  string
	PublicKey   string
}

func GetConfig() Config {
	return appConfig
}

func InitConfig() {
	config := Config{}
	getEnv("DATABASE_DSN", "d", "", "file storage path ", &config.DatabaseDNS)
	getEnv("ADDRESS", "a", ":8080", "Server address", &config.Address)
	flag.Parse()
	appConfig = config
	log.Info("Server config : ", config)

}

func getEnv(envName, flagName, defaultValue, usage string, config *string) {
	flag.StringVar(config, flagName, defaultValue, usage)

	if value := os.Getenv(envName); value != "" {
		log.Info("Using environment variable "+envName, "- value "+value)
		*config = value
	}
}

func getEnvInt(envName string, flagName string, defaultValue int64, usage string, config *int64) {
	flag.Int64Var(config, flagName, defaultValue, usage)

	if value := os.Getenv(envName); value != "" {
		if parsed, err := strconv.ParseInt(value, 10, 64); err == nil {
			*config = parsed
		}
	}
}

func getEnvBool(envName string, flagName string, defaultValue bool, usage string, config *bool) {
	flag.BoolVar(config, flagName, defaultValue, usage)

	if value := os.Getenv(envName); value != "" {
		if parsed, err := strconv.ParseBool(value); err == nil {
			*config = parsed
		}
	}
}
