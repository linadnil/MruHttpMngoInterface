package MruHttpMngoInterface

import (
	"net/http"
	"log"
	"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(Handler))

	log.Fatal(http.ListenAndServe(":8000", mux))
}

func Handler(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	//Connect to localdb
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	Jses := session.DB("test").C("JSONcollection")
	//Xses := session.DB("test").C("XMLcollection")

	err = Jses.Insert(&result)
	if err != nil {
		log.Fatal(err)
	}

	res := []struct_for_JSON{}
	err = Jses.Find(bson.M{}).All(&res)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", res)

	rw.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, "Hi there, I love Golang!")
}


