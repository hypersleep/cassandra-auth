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

func registrationHandler(rw http.ResponseWriter, r *http.Request) {
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

func signInHandler(rw http.ResponseWriter, r *http.Request) {
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

	session, _ := store.Get(r, "session")

	session.Values["email"] = user.Email
	session.Save(r, rw)

	http.Redirect(rw, r, "/users/check", 302)

	responseBuffer = jsend(true, "Successfully signed in!")
	return
}

func logOutHandler(rw http.ResponseWriter, r *http.Request) {
	var responseBuffer []byte

	rw.Header().Set("Content-Type", "application/json")

	defer func() { fmt.Fprintln(rw, string(responseBuffer)) }()

	session, err := store.Get(r, "session")
	if err != nil {
		responseBuffer = jsend(false, "Session error!")
		return
	}

	session.Values["email"] = nil
	session.Save(r, rw)

	responseBuffer = jsend(true, "Successfully logged out!")
	return
}

func checkHandler(rw http.ResponseWriter, r *http.Request) {
	var err error
	var responseBuffer []byte

	rw.Header().Set("Content-Type", "application/json")

	defer func() { fmt.Fprintln(rw, string(responseBuffer)) }()

	user := &User{}

	session, err := store.Get(r, "session")
	if err != nil {
		responseBuffer = jsend(false, "Session error!")
		return
	}

	if str, ok := session.Values["email"].(string); ok {
		user.Email = str
		responseBuffer = jsend(true, "Successfully checked! Hello, " + user.Email + "!")
	} else {
		responseBuffer = jsend(true, "Error!")
	}

	return
}
