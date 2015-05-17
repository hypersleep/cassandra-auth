package main

import(
	"fmt"
	"net/http"

	"github.com/pquerna/ffjson/ffjson"
)

type Data struct {
	Message string `json:"message"`
}

type Status struct {
	Status  bool `json:"status"`
	Data    Data `json:"data"`
}

func jsend(status bool, message string) []byte {
	responseBuffer, _ := ffjson.Marshal(&Status{ status, Data{ message } })
	return responseBuffer
}

func RegistrationHandler(rw http.ResponseWriter, r *http.Request) {
	var err error
	var responseBuffer []byte

	rw.Header().Set("Content-Type", "application/json")

	defer func() { fmt.Fprintln(rw, string(responseBuffer)) }()

	user := &User{}

	err = user.read(r)
	if err != nil {
		responseBuffer = jsend(false, "Wrong request!")
		return
	}

	err = user.register()
	if err != nil {
		responseBuffer = jsend(false, "Registration failed!")
		return
	}

	responseBuffer = jsend(true, "Successfully registred!")
	return
}

func SignInHandler(rw http.ResponseWriter, r *http.Request) {
	var err error
	var responseBuffer []byte

	rw.Header().Set("Content-Type", "application/json")

	defer func() { fmt.Fprintln(rw, string(responseBuffer)) }()

	user := &User{}

	err = user.read(r)
	if err != nil {
		responseBuffer = jsend(false, "Wrong request!")
		return
	}

	err = user.auth()
	if err != nil {
		responseBuffer = jsend(false, "Sign in failed!")
		return
	}

	responseBuffer = jsend(true, "Successfully signed in!")
	return
}

func LogOutHandler(rw http.ResponseWriter, r *http.Request) {
	var err error
	var responseBuffer []byte

	rw.Header().Set("Content-Type", "application/json")

	defer func() { fmt.Fprintln(rw, string(responseBuffer)) }()

	user := &User{}

	err = user.read(r)
	if err != nil {
		responseBuffer = jsend(false, "Wrong request!")
		return
	}

	responseBuffer = jsend(true, "Successfully logged out!")
	return
}

func CheckHandler(rw http.ResponseWriter, r *http.Request) {
	var err error
	var responseBuffer []byte

	rw.Header().Set("Content-Type", "application/json")

	defer func() { fmt.Fprintln(rw, string(responseBuffer)) }()

	user := &User{}

	err = user.read(r)
	if err != nil {
		responseBuffer = jsend(false, "Wrong request!")
		return
	}

	responseBuffer = jsend(true, "Successfully checked!")
	return
}
