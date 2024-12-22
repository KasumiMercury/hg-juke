package config

import (
	"errors"
	"github.com/spf13/viper"
	"log"
	"log/slog"
	"os"
	"path/filepath"
)

const (
	confDirName  = "hg-juke"
	confFileName = "config"
	confFileType = "yaml"
)

func Load() (bool, error) {
	viper.SetConfigName(confFileName)
	viper.SetConfigType(confFileType)

	confDir, err := os.UserConfigDir()
	if err != nil {
		confDir = "~"
	}

	confPath := filepath.Join(confDir, confDirName)
	viper.AddConfigPath(confPath)
	slog.Debug("Loading config file", "path", confPath)

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			return true, err
		}

		if err := createConfigDir(confPath); err != nil {
			return false, err
		}

		fileName := confFileName + "." + confFileType
		filePath := filepath.Join(confPath, fileName)
		slog.Debug("Creating config file", "path", filePath)
		file, err := os.Create(filePath)
		if err != nil {
			return false, err
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Fatalf("failed to close file: %v", err)
			}
		}(file)

		return false, nil
	}

	return true, nil
}

func Set(key string, value interface{}) {
	viper.Set(key, value)
}

func Write() error {
	err := viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

func createConfigDir(path string) error {
	if path == "" {
		panic("path can't be empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}
