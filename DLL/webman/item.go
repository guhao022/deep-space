package webman

import (
	"deep-space/DAL/model"
	"net/http"
	"strconv"
)

// 新建
func NewItem(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	if name == "" {
		NewError(w, ErrMissParam("name"))
		return
	}

	cid := r.FormValue("cid")
	if cid == "" {
		NewError(w, ErrMissParam("cid"))
		return
	}
	if !IsObjectId(cid) {
		NewError(w, ErrForbidden("cid must be ObjectId format"))
	}

	level, err := strconv.Atoi(r.FormValue("level"))
	if err != nil || level == 0 {
		NewError(w, ErrMissParam("level"))
		return
	}

	price, err := strconv.Atoi(r.FormValue("price"))
	if err != nil || price == 0 {
		NewError(w, ErrMissParam("price"))
		return
	}

	abstract := r.FormValue("abstract")
	/*if abstract == "" {
		NewError(w, ErrMissParam("abstarct"))
		return
	}*/

	var item = &model.Item{}

	item.Name = name
	item.Cid = ObjectIdHex(cid)
	item.Level = level
	item.Price = price
	item.FallNum =  0
	item.Abstract = abstract

	err = item.AddItem()
	if err != nil {
		NewError(w, ErrInternalServer(err.Error()))
		return
	}

	Response(w, "new_item", item.Id)

}

// 根据名称查找
func FindItemByName(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if name == "" {
		NewError(w, ErrMissParam("name"))
		return
	}

	var item model.Item
	item.Name = name

	item.FindByName()

	Response(w, "find_item_by_name", item)

}
