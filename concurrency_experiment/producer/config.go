package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.SetEnvPrefix("bg_task")
	err := viper.BindEnv("env")
	err = viper.BindEnv("consul_url")
	err = viper.BindEnv("consul_path")
	if err != nil {
		logrus.Panicln("error binding flag", err)
	}

	consulURL := viper.GetString("consul_url")
	consulPath := viper.GetString("consul_path")
	if consulURL == "" {
		logrus.Fatal("CONSUL_URL missing")
	}
	if consulPath == "" {
		logrus.Fatal("CONSUL_PATH missing")
	}

	err = viper.AddRemoteProvider("consul", consulURL, consulPath)
	if err != nil {
		logrus.Panicln("error binding flag", err)
	}

	viper.SetConfigType("yml")
	err = viper.ReadRemoteConfig()
	if err != nil {
		logrus.Fatal(fmt.Sprintf("%s named \"%s\"", err.Error(), consulPath))
	}

	// loading config
	LoadDB()
}
