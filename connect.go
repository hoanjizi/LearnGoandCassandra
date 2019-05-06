package main
import (
	"github.com/gocql/gocql"
  	"log"
    "time"
)

func PerformOperations() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "hoan"

	cluster.Timeout = 5*time.Second
	cluster.ProtoVersion = 4
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Could not connect")
	}

	keySpaceMeta, _ := session.KeyspaceMetadata("hoan")

	if _, exists := keySpaceMeta.Tables["person"]; exists != true {
		session.Query("CREATE TABLE person (" +
		"id text, name text, phone text, " +
		"PRIMARY KEY (id))").Exec()
	}
	session.Query("INSERT INTO person (id, name, phone) VALUES ('shalabh', 'Shalabh Aggarwal', '1234567890')").Exec()
	var name string
	var phone string
	iter := session.Query("SELECT name, phone FROM person").Iter()
    for iter.Scan(&name, &phone) {
        log.Printf("Iter Name: %v", name)
        log.Printf("Iter Phone: %v", phone)
    }
}

	func main() {
		PerformOperations()
	}
