package main

import(
	"log"
	"net/http"
	"io/ioutil"

	"github.com/gocql/gocql"
	"golang.org/x/crypto/bcrypt"
	"github.com/pquerna/ffjson/ffjson"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *User) read(r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read request body")
		return err
	}

	err = ffjson.Unmarshal(body, user)
	if err != nil {
		log.Println("Failed to unmarshal JSON")
		return err
	}

	return nil
}

func (user *User) register() (err error) {
	cassandraSession, err := newCassandraSession(cassandraHost, cassandraKeyspace)
	if err != nil { return }

	log.Println("Trying to register:", user.Email)

	defer cassandraSession.Close()

	cassandraUsers := map[string]string{
		user.Email: user.Password,
	}

	err = addUsers(cassandraUsers, cassandraSession)
	if err != nil { return }

	log.Println("Succefully registred:", user.Email)

	return
}

func (user *User) auth() (err error) {
	cassandraSession, err := newCassandraSession(cassandraHost, cassandraKeyspace)
	if err != nil { return }

	log.Println("Trying to autheticate:", user.Email)

	defer cassandraSession.Close()

	var email string
	var hashedPassword string
	err = cassandraSession.Query(`SELECT email, encrypted_password FROM users WHERE email=? LIMIT 1`, user.Email).
							  Consistency(gocql.One).
							  Scan(&email, &hashedPassword)
	if err != nil { return }

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))
	if err != nil { return }

	log.Println("Succefully autheticated:", user.Email)

	return

}
