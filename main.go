package main

import (
	"fmt"
	"log"

	"github.com/alijkdkar/ArvanChallenge/controller"
	"github.com/alijkdkar/ArvanChallenge/pkg"
	"github.com/alijkdkar/ArvanChallenge/repository"
	"github.com/gin-gonic/gin"
)

func main() {

	//todo : movie to cmd

	conf := pkg.Config{}.LOAD()
	fmt.Println(conf)
	_, err := repository.NewPostgres()

	fmt.Println("im", err)
	if err != nil {
		panic(err)
	}
	// mux := http.DefaultServeMux
	mux := gin.Default()
	controller.RegisterControllers(mux)

	log.Fatal(mux.Run(":4879"))
}
