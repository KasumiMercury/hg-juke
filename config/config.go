package config

import (
	"errors"
	"github.com/spf13/viper"
	"log/slog"
	"os"
	"path/filepath"
)

func Load() (bool, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	confDir, err := os.UserConfigDir()
	if err != nil {
		confDir = "~"
	}

	confPath := filepath.Join(confDir, "hg-juke")
	viper.AddConfigPath(confPath)
	slog.Debug("Loading config file: %s", confPath)

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			return true, err
		}

		return false, nil
	}

	return true, nil
}
