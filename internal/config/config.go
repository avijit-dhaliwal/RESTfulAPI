package config

import (
    "github.com/spf13/viper"
)

type Config struct {
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
    ServerPort string
    JWTSecret  string
}

func LoadConfig() (Config, error) {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    viper.AutomaticEnv()

    err := viper.ReadInConfig()
    if err != nil {
        return Config{}, err
    }

    var config Config
    err = viper.Unmarshal(&config)
    return config, err
}