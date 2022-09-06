package env

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type server struct {
	Host string
	Port int
}

type database struct {
	DbHost     string
	DbPort     int
	DbName     string
	DbUser     string
	DbPassword string
}

type keys struct {
	SecretKey string
}

type Env struct {
	Server   server
	Database database
	Keys     keys
}

func GetEnv() *Env {
	viper.SetConfigName("appsettings")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	viper.SetDefault("database.dbname", "redstring")

	c := &Env{}
	err := viper.Unmarshal(&c)
	if err != nil {
		log.Fatal("Unable to load configuration file")
	}
	return c
}
