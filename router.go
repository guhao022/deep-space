package main

import (
	"github.com/num5/web"
	"deep-space/DLL/item"
	"net/http"
)

func Router() {
	r := web.New()
	r.SetTrac(true)
	r.Get("/", item.NewItem)

	http.Handle("/", r)
}
