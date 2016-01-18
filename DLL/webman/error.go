package webman

import (
	"github.com/num5/web"
	"net/http"
)

const (
	BadRequest = iota + 400
	ResourceNotFound
	PermissionDenied

	ValidationError
)

func BadRequest() *web.Error {
	return &web.Error{
		ID: "400",
		Links: &web.ErrLinks{
			About: "http://golune.com",
		},
		Status: http.StatusBadRequest,
		Code: "bad_request",
		Title: "bad request",
		Detail: "Bad Request",
		Source: &web.ErrSource{

		},
	}
}
