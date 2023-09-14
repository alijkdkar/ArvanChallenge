package pkg

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type MyPostgres struct {
	User       string
	Password   string
	Host       string
	Port       int
	Db         string
	DataSource string
}
type MyRedis struct {
	RedisAddress string
}

type Config struct {
	MyPostgres
	MyRedis
}

func (post Config) LOAD() *Config {

	pos := MyPostgres{}

	pos.User = os.Getenv("POSTGRES_USER")
	pos.Password = os.Getenv("POSTGRES_PASSWORD")
	pos.Host = os.Getenv("POSTGRES_HOST")
	pos.DataSource = os.Getenv("POSTGRES_SOURCE")
	_port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		log.Fatalln(err)
	}
	pos.Port = _port

	pos.Db = strings.ToLower(os.Getenv("POSTGRES_DB"))

	redis := MyRedis{
		RedisAddress: os.Getenv("REDIS_ADDRESS"),
	}

	return &Config{MyPostgres: pos, MyRedis: redis}
}
