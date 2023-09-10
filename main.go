package main

import (
	"fmt"
	"log"
	"net/http"

	controller "github.com/alijkdkar/ArvanChallenge/Controller"
	repository "github.com/alijkdkar/ArvanChallenge/Repository"
	"github.com/alijkdkar/ArvanChallenge/pkg"
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
	mux := http.DefaultServeMux

	controller.RegisterControllers(mux)

	log.Fatal(http.ListenAndServe(":4879", nil))
}
