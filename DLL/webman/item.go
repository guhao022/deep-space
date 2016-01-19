package webman

import (
	"github.com/num5/web"
	"deep-space/DAL/model"
	"net/http"
	"strconv"
)

func NewItem(w http.ResponseWriter, r *http.Request) {
	var err error
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

	level := r.FormValue("level")
	level, err = strconv.Atoi(level)
	if err != nil || level == 0 {
		NewError(w, ErrMissParam("level"))
		return
	}

	price := r.FormValue("price")
	price, err = strconv.Atoi(level)
	if err != nil || price == 0 {
		NewError(w, ErrMissParam("price"))
		return
	}

	abstract := r.FormValue("abstract")
	if abstract == "" {
		NewError(w, ErrMissParam("abstarct"))
		return
	}

	var item *model.Item

	item.Name = name
	item.Cid = cid
	item.Level = level
	item.Price = price
	item.FallNum =  0
	item.Abstract = abstract

	err = item.AddItem()
	if err != nil {
		NewError(w, ErrInternalServer)
		return
	}

	Success(w, "New Item", item.Id)
}
