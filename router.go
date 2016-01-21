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
	r.Get("/item/new", webman.NewItem)
	r.Get("/item/findbyname", webman.FindItemByName)

}

func RunHttp(addr int) {
	web.SetTrac(true)

	r := web.New()

	Routes(r)

	log.CLog("[TRAC] Server start listen on # %d #\n", addr)
	err := http.ListenAndServeTLS(":"+strconv.Itoa(addr), r)

	if err != nil {
		panic(err)
	}
}