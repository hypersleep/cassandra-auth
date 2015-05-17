package main

import(
	"log"
	"time"
	"errors"

	"github.com/gocql/gocql"
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
			err = errors.New("Can't connect to Cassadra! Quit by timeout!")
	}

	return
}

func addKeyspace(cassandraKeyspace string, cassandraSession *gocql.Session) (err error) {
	return cassandraSession.Query("CREATE KEYSPACE " +
								cassandraKeyspace +
								" WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 }").
								Exec()
}

func addUsers(cassandraUsers map[string]string, cassandraSession *gocql.Session) (err error) {
	for _, user := range cassandraUsers {
		log.Println(user)
		// cassandraSession.Query("CREATE KEYSPACE " +
		// 						cassandraKeyspace +
		// 						" WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 }").
		// 						Exec()
	}

	return
}