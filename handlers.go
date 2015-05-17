package main

import(
	"fmt"
	"net/http"
	"io/ioutil"

	"github.com/pquerna/ffjson/ffjson"
)

type Data struct {
	Message string `json:"message"`
}

type Status struct {
	Status  bool `json:"status"`
	Data    Data `json:"data"`
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegistrationHandler(rw http.ResponseWriter, r *http.Request) {

	var responseBuffer []byte

	rw.Header().Set("Content-Type", "application/json")

	defer func() { fmt.Fprintln(rw, string(responseBuffer)) }()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseBuffer, _ = ffjson.Marshal(&Status{ false, Data{ "Failed to read request body" } })
		return
	}

	user := &User{}
	err = ffjson.Unmarshal(body, user)
	if err != nil {
		responseBuffer, _ = ffjson.Marshal(&Status{ false, Data{ "Failed to unmarshal JSON"  } })
		return
	}

	err = addCassandraUser(user)
	if err != nil {
		responseBuffer, _ = ffjson.Marshal(&Status{ false, Data{ "Registration failed!" + err.Error() } })
		return
	}

	responseBuffer, _ = ffjson.Marshal(&Status{ true, Data{ "Successfully registred!" }})
	return
}

func SignInHandler(rw http.ResponseWriter, r *http.Request) {

	var responseBuffer []byte

	rw.Header().Set("Content-Type", "application/json")

	defer func() { fmt.Fprintln(rw, string(responseBuffer)) }()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseBuffer, _ = ffjson.Marshal(&Status{ false, Data{ "Failed to read request body" } })
		return
	}

	user := &User{}
	err = ffjson.Unmarshal(body, user)
	if err != nil {
		responseBuffer, _ = ffjson.Marshal(&Status{ false, Data{ "Failed to unmarshal JSON"  } })
		return
	}

	err = authCassandraUser(user)
	if err != nil {
		responseBuffer, _ = ffjson.Marshal(&Status{ false, Data{ "Sign in failed!" } })
		return
	}

	responseBuffer, _ = ffjson.Marshal(&Status{ true, Data{ "Successfully signed in!" }})
	return
}

func LogOutHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	responseBuffer, _ := ffjson.Marshal(&Status{ true, Data{} })
	fmt.Fprintln(rw, string(responseBuffer))
}

func CheckHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	responseBuffer, _ := ffjson.Marshal(&Status{ true, Data{} })
	fmt.Fprintln(rw, string(responseBuffer))
}