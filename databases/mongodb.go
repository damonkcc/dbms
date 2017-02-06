package main

import (
	//  "fmt"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

//const MongoDb details
const (
	hosts      = "mongodb-example:27017"
	database   = "admin"
	username   = "admin"
	password   = "admin"
	collection = "system.users"
)

type Role string

type SystemUsers struct {
	//ID       bson.ObjectId `bson:"_id,omitempty"`
	User     string `bson:"user"`
	Database string `bson:"db"`
	Roles    []Role `bson:"roles"`
}

func main() {

	info := &mgo.DialInfo{
		Addrs:    []string{hosts},
		Timeout:  60 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}

	session, err1 := mgo.DialWithInfo(info)
	if err1 != nil {
		log.Printf("ERROR: %+v", err1)
	}
	var result []SystemUsers
	col := session.DB(database).C(collection)
	err := col.Find(nil).All(&result)
	if err != nil {

		log.Printf("ERROR: %+v", err)
	}
	for k, v := range result {
		log.Printf("k: %+v", k)
		log.Printf("v: %+v", v.User)
	}
}
