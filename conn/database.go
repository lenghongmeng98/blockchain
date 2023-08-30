package conn

import (
	"context"
	"fmt"

	"github.com/go-kivik/couchdb/v3"
	"github.com/go-kivik/kivik/v3"
)

var username = "admin"
var password = "123"
var url = "http://localhost:5984/"
var dbName = "student"
var DBConn *kivik.DB

func ConnectToDB() {
	client, err := kivik.New("couch", url)

	if err != nil {
		panic(err)
	} else {
		fmt.Print("Connected!")
	}

	client.Authenticate(context.TODO(), couchdb.BasicAuth(username, password))
	DBConn = client.DB(context.TODO(), dbName)
}
