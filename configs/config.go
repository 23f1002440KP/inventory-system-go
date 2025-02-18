package configs

import (
	"23f1002440KP/inventory-system/logger"
	"flag"

	// "fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/lmittmann/tint"
)

type Configs struct {
	Database struct {
		Name    string `yaml:"name" env-required:"true"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	}`yaml:"database"`
	Server struct {
		Host string `yaml:"host" env-required:"true"`
		Port int    `yaml:"port" env-required:"true"`
	}`yaml:"server"`
}

func LoadConfigs() *Configs {
	var cfgpath string 
	cfgpath = os.Getenv("CONFIG_PATH")
	if cfgpath == "" {
		flags := flag.String("config", "configs/config.yaml", "The path to the config file.")
		flag.Parse()
		cfgpath = *flags
	}

	var cfg Configs
	err := cleanenv.ReadConfig(cfgpath, &cfg)
	if err != nil {
		logger.Logger().Error("Not able to read Config file:",tint.Err(err))
	}
	return &cfg
}