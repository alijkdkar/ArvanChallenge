package pkg

import (
	"log"
	"os"
	"strconv"
	"strings"
)

var SecretKey []byte

type MyPostgres struct {
	User       string
	Password   string
	Host       string
	Port       int
	Db         string
	DataSource string
}

type Config struct {
	MyPostgres
	SecretKey []byte
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
	SecretKey = []byte(os.Getenv("HashSecertKey"))
	return &Config{pos, SecretKey}
}
