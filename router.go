package main

import (
	"github.com/num5/web"
	"net/http"
	"deep-space/DLL/webman"
)

var router = []*web.Route{
	{
		"add_item",
		"GET",
		"/item",
		webman.NewItem,
	},
}

func Router() {
	web.SetTrac(true)

	r := web.Register(router)

	http.Handle("/", r)
}
