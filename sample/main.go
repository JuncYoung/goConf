package main

import (
	"log"

	"github.com/JuncYoung/goConf"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const ProjectName = "Sample"

func main() {
	goConf.AddSearchPath("E:\\qnzs\\aiqc\\src\\goConf\\sample")
	if err := goConf.InitViperConfig(viper.GetViper(), ProjectName, ""); err != nil {
		panic(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("config file changed: %s", e.Name) // update configuration
	})

	value := viper.GetString("test.myKey")
	log.Printf("value: %s\n", value)
}
