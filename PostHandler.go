// GetHandler
package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	//"io/ioutil"
	"net/http"
	"os"
	"time"
)

const (
	// MongoDBHosts Setu constants
	PMongoDBHosts = "localhost:27017"
	PAuthDatabase = "goinggo"
	PAuthUserName = "guest"
	PAuthPassword = "welcome"
	PTestDatabase = "goinggo"
)

func PostHandler(res http.ResponseWriter, req *http.Request) {
	//var role bubbleRole
	fmt.Println("Origin >>" + req.Header.Get("Origin"))
	if origin := req.Header.Get("Origin"); origin != "" {
		res.Header().Set("Access-Control-Allow-Origin", origin)
	}
	//res.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("in post handler ")
	res.Header().Set("Content-Type", "text/json")

	if origin := req.Header.Get("Origin"); origin != "" {
		res.Header().Set("Access-Control-Allow-Origin", origin)
	}

	res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
	res.Header().Set("Access-Control-Allow-Credentials", "true")

	fmt.Println("past headers")

	decoder := json.NewDecoder(req.Body)
	var bubbleR bubbleRole
	err := decoder.Decode(&bubbleR)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Name is " + bubbleR.Name)

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
		fmt.Println("we are connected and ready to post")
	}
	sess.SetMode(mgo.Monotonic, true)
	collection := sess.DB("mystack-dev").C("bubbleroles")
	bubbleR.SetCreated(time.Now())
	err = collection.Insert(bubbleR)

	if err != nil {
		fmt.Printf("Can't insert in to mongo, go error %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("Inserted")
	}
	res.Write([]byte("hello"))
}
