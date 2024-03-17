package configuration

import (
	"fmt"
	cfg "github.com/PiotrFerenc/mash2/internal/consts"
	"github.com/spf13/viper"
)

type configuration struct {
}

func CreateYmlConfiguration() Configuration {
	return &configuration{}
}

func (config *configuration) LoadConfiguration() *Config {
	viper.SetConfigName(cfg.ConfigurationFileName)
	viper.AddConfigPath(cfg.ConfigurationFolderName)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return &Config{
		Queue: QueueConfig{
			QueueRunPipe:      viper.GetString(cfg.QueueRunPipe),
			QueueStageSucceed: viper.GetString(cfg.QueueStageSucceed),
			QueueStageFailed:  viper.GetString(cfg.QueueStageFailed),
			QueueHost:         viper.GetString(cfg.QueueHost),
			QueueVhost:        viper.GetString(cfg.QueueVhost),
			QueueUser:         viper.GetString(cfg.QueueUser),
			QueuePassword:     viper.GetString(cfg.QueuePassword),
			QueuePort:         viper.GetString(cfg.QueuePort),
		},
	}
}
