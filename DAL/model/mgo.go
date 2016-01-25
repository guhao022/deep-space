package model

import (
	"deep-space/DAL/mgo"
	"os"
)

type Mgo struct {
	*mgo.Exec
}

func NewMgo(collection string) *Mgo {
	username := os.Getenv("MGO_USERNAME")
	password := os.Getenv("MGO_PASSWORD")
	return &Mgo{
		&mgo.Exec{
			Database: "deep-space",
			Username: username,
			Password: password,

			Collection: collection,

			Query:  make(map[string]interface{}),
			Change: make(map[string]interface{}),
		},
	}
}
