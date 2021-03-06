package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

func main() {
	var configFile *string = flag.String("c", "myConfig", "Setting the configuration file")
	flag.Parse()

	_, err := os.Stat(*configFile)

	if err == nil {
		fmt.Println("Using User Specified Configuration file!")
		viper.SetConfigFile(*configFile)
	} else {
		viper.SetConfigName(*configFile)
		viper.AddConfigPath("/tmp")
		viper.AddConfigPath("$HOME")
		viper.AddConfigPath(".")
	}

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())

	if viper.IsSet("item1.key1") {
		fmt.Println("item1.val1:", viper.Get("item1.key1"))
	} else {
		fmt.Println("item1.key1 not set!")
	}

	if viper.IsSet("item1.key2") {
		fmt.Println("item1.val2:", viper.Get("item1.key2"))
	} else {
		fmt.Println("item1.key2 not set!")
	}

	if !viper.IsSet("item3.key1") {
		fmt.Println("item3.key1 not set!")
	}
}
