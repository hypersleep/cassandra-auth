package main

import(
	"fmt"
	"net/http"

	"github.com/pquerna/ffjson/ffjson"
)

type Data struct {
	Message []string `json:"message"`
}

type Status struct {
	Status  bool `json:"status"`
	Data    Data `json:"data"`
}

func RegistrationHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	buf, _ := ffjson.Marshal(&Status{ true, Data{} })
	fmt.Fprintln(rw, string(buf))
}

func SignInHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	buf, _ := ffjson.Marshal(&Status{ true, Data{} })
	fmt.Fprintln(rw, string(buf))
}

func LogOutHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	buf, _ := ffjson.Marshal(&Status{ true, Data{} })
	fmt.Fprintln(rw, string(buf))
}

func CheckHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	buf, _ := ffjson.Marshal(&Status{ true, Data{} })
	fmt.Fprintln(rw, string(buf))
}