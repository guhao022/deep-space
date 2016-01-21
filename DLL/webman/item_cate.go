package webman
import (
	"net/http"
	"deep-space/DAL/model"
	"strings"
)

func NewItemCate(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	if name == "" {
		NewError(w, ErrMissParam("name"))
		return
	}

	pid := r.FormValue("pid")
	if pid == "" {
		NewError(w, ErrMissParam("pid"))
		return
	}
	if !IsObjectId(pid) {
		NewError(w, ErrForbidden("cid must be ObjectId format"))
	}

	abstract := r.FormValue("abstract")

	item_cate := &model.ItemCate{
		Name: name,
		Pid: ObjectIdHex(pid),
		Abstract: strings.TrimSpace(abstract),
	}

	err := item_cate.AddItemCate()
	if err != nil {
		NewError(w, ErrInternalServer(err.Error()))
		return
	}
	Response(w, "new_item_cate", item_cate.Id)
}

func ItemCateList(w http.ResponseWriter, r *http.Request) {
	var item model.ItemCate

	v, err := item.FindAll(0,0)
	if err != nil {
		NewError(w, ErrInternalServer(err.Error()))
		return
	}

	Response(w, "findall_item_cate", v)

}
