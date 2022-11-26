package config

import (
	"flag"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var configFileNamePointer *string

func init() {
	// flag -c to read filename from command line
	configFileNamePointer = flag.String("c", "config", "file to read config from")
	flag.Parse()
}

func ReadConfig(base, key string) string {
	configFileName := *configFileNamePointer
	// configFileName := "config"
	v := viper.New()
	v.SetConfigName(configFileName) //name of config file w/o extension
	v.SetConfigType("yaml")         //works even if commented
	v.AddConfigPath(".")
	v.AddConfigPath("./configs")
	err := v.ReadInConfig()
	if err != nil {
		log.Println("viper error:", err)
	}
	return v.GetString(fmt.Sprintf("%v.%v", base, key)) //v.GetString(mongo.password)
}
