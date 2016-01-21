package webman
import (
	"net/http"
	"deep-space/DAL/model"
	"strings"
)

// @name 新建
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

	cate := &model.ItemCate{
		Name: name,
		Pid: ObjectIdHex(pid),
		Abstract: strings.TrimSpace(abstract),
	}

	err := cate.AddItemCate()
	if err != nil {
		NewError(w, ErrInternalServer(err.Error()))
		return
	}
	Response(w, "new_item_cate", cate.Id)
}

// @name 物品分类列表
func ItemCateList(w http.ResponseWriter, r *http.Request) {
	/*name := r.FormValue("name")
	if name == "" {
		NewError(w, ErrMissParam("name"))
		return
	}*/

	var cate model.ItemCate

	v, err := cate.FindAll(0,0)
	if err != nil {
		NewError(w, ErrInternalServer(err.Error()))
		return
	}

	Response(w, "findall_item_cate", v)

}

// @name 删除分类
func DelItemCate(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		NewError(w, ErrMissParam("id"))
		return
	}
	if !IsObjectId(id) {
		NewError(w, ErrForbidden("id must be ObjectId format"))
	}

	var cate model.ItemCate

	err := cate.DelById(id)
	if err != nil {
		NewError(w, ErrInternalServer(err.Error()))
		return
	}

	Response(w, "del_item_cate", "")

}


