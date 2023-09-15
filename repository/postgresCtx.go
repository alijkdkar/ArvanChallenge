package repository

import (
	"fmt"
	"html/template"
	"strconv"
	"strings"

	"github.com/alijkdkar/ArvanChallenge/domain"
	"github.com/alijkdkar/ArvanChallenge/pkg"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const connectionString = "postgres://{{.User}}:{{.Password}}@{{.Host}}:{{.Port}}/{{.Db}}?sslmode=disable"

var DB *gorm.DB

func connectionStringMaker(conf pkg.MyPostgres) string {
	sb := strings.Builder{}
	temp := template.Must(template.New("connString").Parse(connectionString))
	fmt.Println(temp)
	if err := temp.Execute(&sb, conf); err != nil {
		panic(err)
	}
	return sb.String()
}

func NewPostgres() (*gorm.DB, error) {
	conf := pkg.Config{}.LOAD()
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		conf.User,
		conf.Password,
		conf.Host,
		strconv.Itoa(conf.Port),
		"postgres")

	// connect to the postgres db just to be able to run the create db statement
	db1, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// check if db exists
	stmt := fmt.Sprintf("SELECT * FROM pg_database WHERE datname = '%s';", conf.Db)
	rs := db1.Raw(stmt)
	if rs.Error != nil {
		fmt.Println("Data base not exists")
	}
	//create database
	var rec = make(map[string]interface{})
	if rs.Find(rec); len(rec) == 0 {
		stmt := fmt.Sprintf("CREATE DATABASE %s;", conf.Db)
		if rs := db1.Exec(stmt); rs.Error != nil {
			fmt.Println("Error on Create DataBase")
		}
	}

	conn := connectionStringMaker(conf.MyPostgres)

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{SkipDefaultTransaction: true})

	if err != nil {
		return db, err
	}

	db.AutoMigrate(&domain.User{}, &domain.CreditCard{}, &domain.Transaction{})
	DB = db

	return db, nil
}
