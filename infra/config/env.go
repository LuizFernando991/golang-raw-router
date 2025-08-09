package config

import "github.com/spf13/viper"

var env *Envconfig

type Envconfig struct {
	API_PORT    string `mapstructure:"API_PORT"`
	DB_USER     string `mapstructure:"DB_USER"`
	DB_PASSWORD string `mapstructure:"DB_PASSWORD"`
	DB_DATABASE string `mapstructure:"DB_DATABASE"`
	DB_HOST     string `mapstructure:"DB_HOST"`
	DB_PORT     string `mapstructure:"DB_PORT"`
	ENV         string `mapstructure:"ENVIRONMENT"`
}

func GetEnvConfig() *Envconfig {
	return env
}

func LoadEnv(path string) (*Envconfig, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		panic(err)
	}

	return env, nil
}
