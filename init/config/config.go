package config

import (
	"EffectiveMobile/init/logger"
	"EffectiveMobile/pkg/constants"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var ServerConfig Config

type Config struct {
	APIPort  string `mapstructure:"API_PORT"`
	APIDebug bool   `mapstructure:"API_DEBUG"`

	POSTGRESDsn string `mapstructure:"POSTGRES_DSN"`
}

func InitConfig() error {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./configs")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		logger.FatalF("Error reading config file, %v", logrus.Fields{constants.LoggerCategory: constants.Config}, err.Error())

		return err
	}

	if err := viper.Unmarshal(&ServerConfig); err != nil {
		logger.ErrorF("error unmarshal config, %v", logrus.Fields{constants.LoggerCategory: constants.Config}, err.Error())

		return err
	}

	if ServerConfig.APIPort == "" || ServerConfig.POSTGRESDsn == "" {
		logger.Error("missing requirement variable!", logrus.Fields{constants.LoggerCategory: constants.Config})

		return constants.ErrEmptyVar
	}

	return nil
}
