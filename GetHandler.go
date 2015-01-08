// GetHandler
package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"net/http"
	"os"
	"time"
)

const (
	// MongoDBHosts Setu constants
	MongoDBHosts = "localhost:27017"
	AuthDatabase = "goinggo"
	AuthUserName = "guest"
	AuthPassword = "welcome"
	TestDatabase = "goinggo"
)

type bubbleRole struct {
	Name    string    `bson:"name" json:"name"`
	Value   int       `bson:"value"`
	Active  bool      `bson:"active"`
	Created time.Time `bson:"created"`
}

func (a *bubbleRole) SetCreated(created time.Time) {
	a.Created = created
}

func GetHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/json")
	fmt.Println("Origin >>" + req.Header.Get("Origin"))

	if origin := req.Header.Get("Origin"); origin != "" {
		res.Header().Set("Access-Control-Allow-Origin", origin)
	}
	res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
	res.Header().Set("Access-Control-Allow-Credentials", "true")

	uri := os.Getenv("MONGOHQ_URL")
	if uri == "" {
		uri = MongoDBHosts
	}
	fmt.Println("uri of the host equals " + uri)
	sess, err := mgo.Dial(uri)
	defer sess.Close()
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("we are connected")
	}
	sess.SetMode(mgo.Monotonic, true)
	collection := sess.DB("mystack-dev").C("bubbleroles")
	var bubbleRoles []bubbleRole
	err = collection.Find(nil).All(&bubbleRoles)

	if err != nil {
		fmt.Printf("RunQuery : ERROR : %s\n", err)
		return
	} else {
		js, err := json.Marshal(bubbleRoles)
		if err != nil {
			fmt.Printf("JSON Error : ERROR : %s\n", err)
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		} else {
			res.Write(js)
		}
	}

}
