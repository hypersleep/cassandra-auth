package main

import(
	"log"
	"time"
	"errors"

	"github.com/gocql/gocql"
	"golang.org/x/crypto/bcrypt"
)

func connectCassandra(cassandraHost string, cassandraSessionChannel chan *gocql.Session, cassandraKeyspace string) {
	for {
		cluster := gocql.NewCluster(cassandraHost)
		if cassandraKeyspace != "" {
			cluster.Keyspace = cassandraKeyspace
		}
		session, err := cluster.CreateSession()

		if err == nil {
			cassandraSessionChannel <- session
			return
		}

		log.Println("Can't connect cassandra! Try again after 1 second!")
		time.Sleep(1 * time.Second)
	}
}

func newCassandraSession(cassandraHost string, cassandraKeyspace string) (cassandraSession *gocql.Session, err error) {
	var cassandraSessionChannel chan *gocql.Session = make(chan *gocql.Session)

	go connectCassandra(cassandraHost, cassandraSessionChannel, cassandraKeyspace)

	select {
		case cassandraSession = <- cassandraSessionChannel:
			log.Println("Cassandra connection established!")
		case <- time.After(10 * time.Second):
			err = errors.New("Can't connect to Cassadra! Quit by 10 seconds timeout!")
	}

	return
}

func addKeyspace(cassandraKeyspace string, cassandraSession *gocql.Session) (err error) {
	query := "CREATE KEYSPACE " + cassandraKeyspace + " WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 }"
	return cassandraSession.Query(query).Exec()
}

func addUsers(cassandraUsers map[string]string, cassandraSession *gocql.Session) (err error) {
	var query string
	for email, password := range cassandraUsers {

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
		if err != nil { return err }

		query = "INSERT INTO users (email, encrypted_password) VALUES (?, ?)"
		err = cassandraSession.Query(query, email, hashedPassword).Exec()
		if err != nil { return err }
	}

	return
}

func migrateCassandra(cassandraHost string, cassandraKeyspace string, cassandraUsers map[string]string) error  {

	// Cassandra keyspace creation

	cassandraSession, err := newCassandraSession(cassandraHost, "")
	if err != nil {
		return err
	}

	defer cassandraSession.Close()

	err = addKeyspace(cassandraKeyspace, cassandraSession)
	if err != nil {
		return errors.New("Can't create keyspace: " + err.Error())
	}

	cassandraSession.Close()

	// Cassandra migrations section

	cassandraSession, err = newCassandraSession(cassandraHost, cassandraKeyspace)
	if err != nil {
		return err
	}

	query := "CREATE TABLE users (email text, encrypted_password text, PRIMARY KEY (email))"
	err = cassandraSession.Query(query).Exec()
	if err != nil { return err }

	err = addUsers(cassandraUsers, cassandraSession)
	if err != nil {
		return errors.New("Can't add users: " + err.Error())
	}

	return nil
}
