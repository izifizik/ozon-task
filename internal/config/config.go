package config

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

type Config struct {
	DB   *pgx.Conn
	Port string
	Host string
}

var once sync.Once
var instance Config

//NewConfig - create config from env
func NewConfig() Config {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Println(err.Error())
		}
		ctx := context.Background()
		instance.Port = os.Getenv("PORT_APP")
		instance.Host = os.Getenv("HOST_APP")

		url := os.Getenv("URL_DB")
		instance.DB = connectToDB(ctx, url)
	})
	return instance
}

func connectToDB(ctx context.Context, url string) *pgx.Conn {
	log.Println("try to connect to psql")
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		log.Fatalln("error with connect to psql:" + err.Error())
	}
	log.Println("connect to psql success")
	log.Println("try to ping psql")
	err = conn.Ping(ctx)
	if err != nil {
		log.Fatalln("error to ping psql:" + err.Error())
	}
	log.Println("connect to psql success")

	return conn
}
