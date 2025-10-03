package config

import (
    "log"

    "github.com/spf13/viper"
)

type Config struct {
    DBHost     string `mapstructure:"DB_HOST"`
    DBPort     string `mapstructure:"DB_PORT"`
    DBUser     string `mapstructure:"DB_USER"`
    DBPassword string `mapstructure:"DB_PASSWORD"`
    DBName     string `mapstructure:"DB_NAME"`
    ServerPort string `mapstructure:"SERVER_PORT"`
}

func LoadConfig() (Config, error) {
    var config Config

    viper.SetConfigFile(".env")
    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err != nil {
        log.Println("No .env file found, using only environment variables")
    }

    if err := viper.Unmarshal(&config); err != nil {
        return config, err
    }

    return config, nil
}
