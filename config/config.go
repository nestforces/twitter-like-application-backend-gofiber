package config

import (
    "github.com/spf13/viper"
)

type Config struct {
    DBUrl    string `mapstructure:"DB_URL"`
    JWTSecret string `mapstructure:"JWT_SECRET"`
}

var AppConfig *Config

func LoadConfig() error {
    viper.AddConfigPath(".")
    viper.SetConfigFile(".env")
    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err != nil {
        return err
    }
    if err := viper.Unmarshal(&AppConfig); err != nil {
        return err
    }
    return nil
}
