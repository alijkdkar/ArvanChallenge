package main

import (
	"fmt"
	"log"

	"github.com/alijkdkar/ArvanChallenge/controller"
	discountrepository "github.com/alijkdkar/ArvanChallenge/discount-repository"
	discountcontroller "github.com/alijkdkar/ArvanChallenge/discountController"
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

	if err := discountrepository.NewRedisCoontext(); err != nil {
		panic(err)
	}

	mux := gin.Default()

	//set acount manager controller handler service 1
	controller.RegisterControllers(mux)

	//set discount manager controller handler service 2
	discountcontroller.RegisterDiscountServices(mux)

	//two below services just has relation together
	log.Fatal(mux.Run(":4879"))
}
