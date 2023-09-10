package controller

import (
	"encoding/json"
	"net/http"

	repository "github.com/alijkdkar/ArvanChallenge/Repository"
	"github.com/alijkdkar/ArvanChallenge/domain"
)

func UserHandlerRegister(ctx *http.ServeMux) {
	baseUrl := "/User"
	ctx.HandleFunc(baseUrl+"", User)
}

func User(w http.ResponseWriter, r *http.Request) {
	rep := repository.NewUserRepository()
	if r.Method == http.MethodPost {

		user := domain.CreateNewUser("behzad", "bluekian", "09132120832")
		err := rep.Create(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - some erorr occerd "))
		}
		w.Write([]byte("User Saved"))
	} else if r.Method == http.MethodGet {
		allUser := rep.GetUsers()
		json, err := json.Marshal(allUser)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - some erorr occerd "))
		}
		w.Write(json)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400-Bad Requst"))
	}

}
