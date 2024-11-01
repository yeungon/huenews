package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Env struct {
	Token      string
	Channel_ID string
}

var once sync.Once
var env *Env

func New() {
	once.Do(func() {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		token := os.Getenv("TOKEN")
		channel_id := os.Getenv("CHANNEL_ID")

		env = &Env{
			Token:      token,
			Channel_ID: channel_id,
		}

	})
}

func Get() *Env {
	return env
}
