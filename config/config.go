package config

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// C alias for config
var C Config

// InitConfig with yaml file path
func InitConfig(configPath string) {
	if configPath != "" {
		viper.AddConfigPath(configPath)
	} else {
		viper.AddConfigPath("./")
	}
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		logrus.Errorln("viper.ReadInConfig error -> ", err)
	}

	C = Config{}

	C.Server.Mode = viper.GetString("server.mode")
	C.Server.Port = viper.GetInt("server.port")

	C.Redis.Addr = viper.GetString("redis.addr")
	C.Redis.Password = viper.GetString("redis.password")
	C.Redis.DB = viper.GetInt("redis.db")

	C.Mongo.Addrs = viper.GetStringSlice("mongo.addrs")
	C.Mongo.Database = viper.GetString("mongo.database")
	C.Mongo.ReplicaSet = viper.GetString("mongo.replicaSet")
	C.Mongo.Source = viper.GetString("mongo.source")
	C.Mongo.Username = viper.GetString("mongo.username")
	C.Mongo.Password = viper.GetString("mongo.password")
}
