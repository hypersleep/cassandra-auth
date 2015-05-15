package main

import(
	"fmt"
	"net/http"
)

func RegistrationHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Hello")
}

func SignInHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Hello")
}

func LogOutHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Hello")
}

func CheckHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Hello")
}