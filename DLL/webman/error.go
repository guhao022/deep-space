package webman
import (
	"net/http"
	"encoding/json"
)

type Error struct {
	Id     string `json:"id"`
	Status int    `json:"_"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func (err Error) Error() string {
	return err.Detail
}

func NewError(w http.ResponseWriter, err Error) {
	w.WriteHeader(err.Status)
	w.Header().Set("content-type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(err)
}


var (
	ErrBadRequest     = &Error{"bad_request", 400, "Bad request", "Request body is not well-formed. It must be JSON."}
	ErrInternalServer = &Error{"internal_server_error", 500, "Internal Server Error", "Something went wrong."}
	ErrNoData         = &Error{"no_data", 422, "No data", `Key "data" in the top level of the JSON document is missing or contains no data`}
	ErrUnauthorized   = &Error{"unauthorized", 401, "Unauthorized", "Access token is invalid."}
	ErrNotFound       = &Error{"not_found", 404, "Not found", "Route not found."}
)

func ErrMissParam(param string) *Error {
	return &Error{
		"miss_param",
		400,
		"Miss Param",
		"miss query param: " + param,
	}
}

type Success struct {
	Status string      `json:"_"`
	Title  string 	   `json:"title,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

func Success(w http.ResponseWriter, title string, data interface{}) {
	success := &Success{200, title, data}
	w.WriteHeader(success.Status)
	w.Header().Set("content-type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(success)
}
