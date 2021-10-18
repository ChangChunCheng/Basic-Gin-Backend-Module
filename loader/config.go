// Package loader
package loader

import (
	"os"
	"path"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		logrus.Fatalf("Warning: %s environment variable not set.\n", k)
	}
	return v
}

// loadConfig - Loading config file by viper
func loadConfig() {
	viper.AllowEmptyEnv(false)

	viper.SetConfigFile(path.Join("config", "conf.json"))
	viper.SetConfigType("json")
	viper.WatchConfig()

	err := viper.ReadInConfig()
	if err != nil {
		panic("Fatal error config file")
	}

	viper.AutomaticEnv()
}
