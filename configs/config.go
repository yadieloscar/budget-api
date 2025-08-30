// Package configs contains configuration loading utilities for the service.
package configs

import (
    "github.com/spf13/viper"
    "log"
)

// Config represents application runtime configuration.
type Config struct {
    Port string `mapstructure:"PORT"`
}

// LoadConfig reads configuration from a file named `config.env` (via Viper)
// in the current working directory, and unmarshals it into Config.
func LoadConfig() (config Config, err error) {
    viper.SetConfigName("config")
    viper.SetConfigType("env")
    viper.AddConfigPath(".")

    if err = viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file, %s", err)
        return
    }

    if err = viper.Unmarshal(&config); err != nil {
        log.Fatalf("Unable to decode into struct, %v", err)
    }

    return
}
