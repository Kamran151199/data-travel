package config

import (
	"fmt"
	"github.com/Kamran151199/dbmigrate/internal/config/target"
	"github.com/Kamran151199/dbmigrate/pkg/validation"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type Config[O target.Target, D target.Target] struct {
	Source O `json:"source" mapstructure:"source" validate:"required"`
	Sink   D `json:"sink" mapstructure:"sink" validate:"required"`
}

func LoadTarget[T target.Target](path string, name string, ext string) (*T, error) {
	var config T

	// set viper config params
	reader := viper.New()
	reader.SetConfigName(name)
	reader.AddConfigPath(path)
	reader.SetConfigType(ext)

	// read config
	if err := reader.ReadInConfig(); err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}

	// unmarshal config to target struct
	if err := reader.Unmarshal(&config); err != nil {
		log.Fatalf("Could not unmarshal: %s \n", err)
	}

	// enable env replacer
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// enable real time change lookup
	reader.WatchConfig()
	reader.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name, e.Op)
	})

	// validate config by also setting default values
	validator := validation.NewWithDefaultsValidator()
	if err := validator.ValidateWithDefaults(&config); err != nil {
		return nil, fmt.Errorf("Could not validate: %s \n", err.Error())
	}
	return &config, nil
}

// LoadConfig loads configuration from a file.
func LoadConfig[O target.Target, D target.Target](sourcePath, sourceName, sourceExt,
	sinkPath, sinkName, sinkExt string) (*Config[O, D], error) {

	source, err := LoadTarget[O](sourcePath, sourceName, sourceExt)
	if err != nil {
		return nil, fmt.Errorf("Could not load source: %s \n", err)
	}

	sink, err := LoadTarget[D](sinkPath, sinkName, sinkExt)
	if err != nil {
		return nil, fmt.Errorf("Could not load sink: %s \n", err)
	}

	return &Config[O, D]{
		Source: *source,
		Sink:   *sink,
	}, nil
}
