package model

import (
	"deep-space/DAL/mgo"
	"os"
)


func NewMgo(collection string) *mgo.Exec {
	username := os.Getenv("MGO_USERNAME")
	password := os.Getenv("MGO_PASSWORD")
	println("dbname:"+username)
	return &mgo.Exec{
		Database: "deep-space",
		Username: username,
		Password: password,

		Collection: collection,

		Query: make(map[string]interface{}),
		Change: make(map[string]interface{}),
	}
}



