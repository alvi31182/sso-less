package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string        `yaml:"env" env-default:"local"`
	StoragePath string        `yaml:"storage_path" env-required:"true"`
	TokenTtl    time.Duration `yaml:"token_ttl" env-default:"1h"`
	GRPC        GRPCconfig    `yaml:"grpc"`
}

type GRPCconfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout" env-default:"10h"`
}

func MustLoad() *Config {
	path := fefecthConfigPath()
	if path == "" {
		panic("Error path empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("Error path empty")
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("Error read config: " + err.Error())
	}

	return &cfg
}

func fefecthConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "config/config.yaml", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	if res == "" {
		res = "config/config.yaml"
	}

	return res
}
