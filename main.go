package main

import(
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	port := "8080"

	r := mux.NewRouter()

	users := r.PathPrefix("/users").Subrouter()

	users.Methods("POST").
		  Path("/create").
		  HandlerFunc(RegistrationHandler)
	users.Methods("POST").
		  Path("/signin").
		  HandlerFunc(SignInHandler)
	users.Methods("DELETE").
		  Path("/logout").
		  HandlerFunc(LogOutHandler)
	users.Methods("POST").
		  Path("/check").
		  HandlerFunc(CheckHandler)

	fmt.Println("Server running on port:", port)
	http.ListenAndServe(":" + port, r)
}