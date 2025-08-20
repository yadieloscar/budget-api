package configs

import (
    "github.com/spf13/viper"
    "log"
)

type Config struct {
    Port string `mapstructure:"PORT"`
}

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