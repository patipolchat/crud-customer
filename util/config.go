package util

import (
	"crud-customer/util/validator"
	"github.com/spf13/viper"
	"strings"
)

func GetConfig[T any]() (*T, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	validate := validator.GetValidator()
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	var cfg T
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	if err := validate.Struct(cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
