package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"path/filepath"
	"runtime"
	"sync"
)

const configName = "config.yml"

type Config struct {
	IsDebug *bool `yaml:"is_debug" env-required:"true"`
	GRPC    struct {
		ServerPort string `yaml:"port" env-default:"8080"`
	} `yaml:"grpc_server_config"`
	AppConfig AppConfig `yaml:"app_config"`
}

type AppConfig struct {
	LogLevel string `yaml:"log_level"`
}

var _, b, _, _ = runtime.Caller(0)
var rootPath = filepath.Join(filepath.Dir(b), "../..")

var help, _ = cleanenv.GetDescription(instance, nil)

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}

		path, err := filepath.Abs(rootPath + "/" + configName)
		if err != nil {
			log.Print(help)
			log.Fatal(err)
		}

		if err := cleanenv.ReadConfig(path, instance); err != nil {
			log.Print(help)
			log.Fatal(err)
		}
	})

	return instance
}
