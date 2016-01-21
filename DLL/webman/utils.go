package webman
import "gopkg.in/mgo.v2/bson"

func IsObjectId(id string) bool {
	return bson.IsObjectIdHex(id)
}

func ObjectIdHex(id string) bson.ObjectId {
	return bson.ObjectIdHex(id)
}
