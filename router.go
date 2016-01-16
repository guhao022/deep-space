package main

import (
	"web"
	"net/http"
	"deep-space/DLL/webman"
	"deep-space/utils/log"
	"strconv"
)

func Routes(r *web.Router) {
	r.Group("/item",
		r.Get("/aa", webman.NewItem),
		r.Get("/bb", webman.NewItem),
	)
	//r.Get("/aa", webman.NewItem).PathPrefix("/item").Subrouter()

}


func RunHttp(addr int) {
	web.SetTrac(true)

	r := web.New()

	Routes(r)

	log.Tracf("Server start listen on %d", addr)
	err := http.ListenAndServe(":" + strconv.Itoa(addr), r)

	if err != nil {
		panic(err)
	}
}
