package cassandra

import (
	"github.com/gocql/gocql"
)

func NewCassandraSession() *gocql.Session {
	cluster := gocql.NewCluster("scylla")
	cluster.Keyspace = "wallet"
	session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	return session
}
