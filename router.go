package main

import (
	"net/http"
	"strconv"

	"deep-space/DLL/webman"
	"deep-space/log"
	"github.com/num5/web"
)

func Routes(r *web.Router) {
	// item
	r.Post("/item/new", webman.NewItem)
	r.Post("/item/findbyname", webman.FindItemByName)

	//item_cate
	r.Post("/item/cate/new", webman.NewItemCate)
	r.Get("/item/cate/list", webman.ItemCateList)

}

func RunHttp(addr int) {
	web.SetTrac(true)

	r := web.New()

	Routes(r)

	log.CLog("[TRAC] Server start listen on # %d #\n", addr)

	err := http.ListenAndServe(":"+strconv.Itoa(addr), r)

	if err != nil {
		panic(err)
	}
}