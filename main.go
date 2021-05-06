package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func main() {
	fmt.Println("starting ghproject")
	viper.AddConfigPath(".")
	viper.SetConfigFile("config")
	viper.SetConfigType("yml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	for key, val := range viper.AllSettings() {
		fmt.Println("key: ", key, ", val: ", val)
	}
	os.Exit(0)
}
