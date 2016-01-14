package model

func InitDB(){
	itemc := Mgo.C("item")
	//icc := Mgo.C("item_cate")

	itemc.Insert()

}
