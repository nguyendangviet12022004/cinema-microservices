package app

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("cinema")
	viper.SetConfigType("yaml")
	// viper.AddConfigPath("/etc/cinema/")
	// viper.AddConfigPath("$HOME/.cinema/")
	viper.AddConfigPath("./config/")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

}
