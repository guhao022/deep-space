package webman
import (
	"net/http"
	"deep-space/DAL/model"
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

	var item_cate model.ItemCate

	item_cate.Name = name
	item_cate.Pid = ObjectIdHex(pid)
	item_cate.Abstract = abstract

	/*err := item_cate.AddItemCate()
	if err != nil {
		NewError(w, ErrInternalServer(err.Error()))
		return
	}
*/
	Response(w, "New Item", item_cate.Id)

}
