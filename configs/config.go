package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	Db *DbConfig
}

type DbConfig struct {
	Host                   string `mapstructure:"DB_HOST"`
	Port                   string `mapstructure:"DB_PORT"`
	UserName               string `mapstructure:"DB_USER_NAME"`
	Password               string `mapstructure:"DB_PWD"`
	DbName                 string `mapstructure:"DB_NAME"`
	MaxIdleConnection      int    `mapstructure:"DB_MAX_IDLE_CONN"`
	MaxOpenConnection      int    `mapstructure:"DB_MAX_OPEN_CONN"`
	ConnMaxLifetimeMinutes int    `mapstructure:"DB_CONN_MAX_LIFE_MIN"`
	MigrationDirName       string `mapstructure:"MIGRATION_DIR_NAME"`
}

func LoadConfig(path string, configName string, configType string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	var dbConf *DbConfig
	err = viper.Unmarshal(&dbConf)
	config.Db = dbConf
	return config, nil
}
