package common

import (
	"fmt"
	"log"
	"net/url"
	"sync"

	"github.com/tructn/redstring/env"
	"github.com/tructn/redstring/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var instance *gorm.DB
var lock = &sync.Mutex{}

func GetDatabase(env env.Env) *gorm.DB {
	if instance != nil {
		return instance
	}

	lock.Lock()
	defer lock.Unlock()

	dsn := url.URL{
		User:   url.UserPassword(env.Database.DbUser, env.Database.DbPassword),
		Scheme: "postgres",
		Host:   fmt.Sprintf("%s:%d", env.Database.DbHost, env.Database.DbPort),
		Path:   env.Database.DbName,
	}

	instance, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn.String(),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	//TODO:
	instance.AutoMigrate(&model.User{}, &model.Player{})

	fmt.Printf("Load %s database successfully", env.Database.DbName)

	return instance
}
