package webman

import (
	"web"
	"deep-space/DAL/model"
)

var item *model.Item

func NewItem(ctx *web.Context) {
	name := ctx.GetString("name")
	if name == "" {
		ctx.Json("miss query param: name")
		return
	}

	ctx.WriteString("aaaa")

}
