package controller

import "net/http"

func RegisterControllers(ctx *http.ServeMux) {
	UserHandlerRegister(ctx)
	//Add Other Controllers
}
