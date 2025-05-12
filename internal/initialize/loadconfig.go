package initialize

import (
	"fmt"
	"vulcan_labs_cinema/global"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// configure structure
	err = viper.Unmarshal(&global.Config)

	if err != nil {
		panic(fmt.Errorf("unable to decode into struct: %w", err))
	}

	// Load Config to global.Config
	fmt.Println("load config success to global.Config")
}
