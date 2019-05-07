package main
import (
	"github.com/gocql/gocql"
  	"log"
	"time"
	"fmt"
)

func PerformOperations() {
	cluster := gocql.NewCluster("127.0.0.1")
	// cluster.Authenticator = gocql.PasswordAuthenticator{
    //     Username: "cassandra",
    //     Password: "cassandra",
    // }
	cluster.Keyspace = "demo"

	cluster.Timeout = 5*time.Second
	cluster.ProtoVersion = 4
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	
	if err != nil {
		log.Fatalf("Could not connect")
	}

	

	keySpaceMeta, _ := session.KeyspaceMetadata("demo")

	if _, exists := keySpaceMeta.Tables["person"]; exists != true {
		session.Query("CREATE TABLE person (" +
		"id text, name text, phone text, " +
		"PRIMARY KEY (id))").Exec()
	}
	//session.Query("INSERT INTO person (id, name, phone) VALUES ('shalh', 'Shalabh Aggarwal', '1234567890');").Exec()
	if err := session.Query(`INSERT INTO person (id, name, phone) VALUES (?, ?, ?)`,
		"125", "hoan", "hello world").Exec(); err != nil {
		log.Fatal(err)
	}
	var name string
	var phone string
	iter := session.Query("SELECT name, phone FROM person").Iter()
    for iter.Scan(&name, &phone) {
		fmt.Println("Iter Name: "+name)
		fmt.Println("Iter Phone: "+phone)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
}

	func main() {
		PerformOperations()
	}
