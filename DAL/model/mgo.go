package model

import (
	"deep-space/DAL/mgo"
	"os"
)

var Mgo = &mgo.Exec{
	Database: "black",
	Username: os.Getenv("MGO_USERNAME"),
	Password: os.Getenv("MGO_PASSWORD"),

	Query: make(map[string]interface{}),
	Change: make(map[string]interface{}),
}

var C mgo.Mgo


