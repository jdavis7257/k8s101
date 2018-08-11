package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var mongoSession *mgo.Session
var mongoCollection *mgo.Collection

const DB_NAME = "demo"
const COLLECTION_NAME = "comments"

type Comment struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	Message string        `bson:"message,omitempty"`
}

func apiRequestHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("api request received")
	log.Println(r.RequestURI)
	if strings.HasSuffix(r.RequestURI, "new") {
		log.Println("Saving comment...")
		r.ParseMultipartForm(2)
		log.Println(r.Form)
		text := r.Form.Get("comment")
		log.Println(text)
		err := mongoCollection.Insert(&Comment{Message: text})
		if err != nil {
			log.Println("Did not save!!")
		}
	}

	if strings.HasSuffix(r.RequestURI, "/comment/getAll") {
		log.Println("Getting docs...")
		var results []Comment
		err := mongoCollection.Find(nil).All(&results)
		if err != nil {
			log.Println("Could not get documents due to an error.")
		} else {
			log.Println(results)
			bytes, errM := json.Marshal(results)
			if errM == nil {
				w.Write(bytes)
			}
		}
	}
}

func main() {

	session, err := mgo.Dial("mongo")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(DB_NAME).C(COLLECTION_NAME)

	mongoSession = session
	mongoCollection = c

	fs := http.FileServer(http.Dir("static"))
	http.HandleFunc("/api/v1/", apiRequestHandler)
	http.Handle("/", fs)
	http.ListenAndServe(":8080", nil)
}
