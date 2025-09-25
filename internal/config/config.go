package config

import (
	"github.com/spf13/viper"
)

type GormCfg struct {
	Path string `mapstructure:"path"`
}

type JSONCfg struct {
	Path string `mapstructure:"path"`
}

type StorageCfg struct {
	Type   string   `mapstructure:"type"`
	Gorm   GormCfg  `mapstructure:"gorm"`
	JSON   JSONCfg  `mapstructure:"json"`
	Memory struct{} `mapstructure:"memory"`
}

type AppConfig struct {
	Storage StorageCfg `mapstructure:"storage"`
}

func Load() (AppConfig, error) {
	var cfg AppConfig
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		return cfg, err
	}
	if err := v.Unmarshal(&cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}
