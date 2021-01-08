package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
	"github.com/vgmdj/ddd-case/infrastructure/repository/persistence"
	"github.com/vgmdj/ddd-case/interfaces/httpfacade"
)

// Config the global config
type Config struct {
	Env        string                `mapstructure:"env"`
	HTTPConfig httpfacade.HTTPConfig `mapstructure:"http"`
	DBConfig   persistence.Config    `mapstructure:"database"`
}

const (
	// EnvDev dev
	EnvDev = "dev"
	// EnvTest test
	EnvTest = "test"
	// EnvPre pre
	EnvPre = "pre"
	// EnvProd prod
	EnvProd = "prod"
)

var (
	configPath = map[string]string{
		EnvDev:  strings.TrimRight(os.Getenv("AUTOMATIONPATH"), "/") + "/config",
		EnvTest: "config/",
		EnvPre:  "config/",
		EnvProd: "config/",
	}

	env          = os.Getenv("env")
	globalConfig *Config
)

func init() {
	c := &Config{}
	viper.AddConfigPath(configPath[env])
	viper.SetConfigName(fmt.Sprintf("config.%s", env))
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}

	err = viper.Unmarshal(c)
	if err != nil {
		panic(err.Error())
	}

	c.HTTPConfig.Env = c.Env

	globalConfig = c
}

// GlobalEntity return the global config
func GlobalEntity() *Config {
	return globalConfig
}
