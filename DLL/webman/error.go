package webman
import (
	"net/http"
	"encoding/json"
)

type Error struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`

	status int    `json:"_"`
}

func (err Error) Error() string {
	return err.Detail
}

func NewError(w http.ResponseWriter, err *Error) {
	w.WriteHeader(err.status)
	w.Header().Set("content-type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(err)
}


var (
	ErrBadRequest     = &Error{"bad_request", "Bad request", "Request body is not well-formed. It must be JSON.", 400}
	ErrInternalServer = &Error{"internal_server_error", "Internal Server Error", "Something went wrong.", 500}
	ErrNoData         = &Error{"no_data", "No data", `Key "data" in the top level of the JSON document is missing or contains no data`, 422}
	ErrUnauthorized   = &Error{"unauthorized", "Unauthorized", "Access token is invalid.", 401}
	ErrNotFound       = &Error{"not_found", "Not found", "Route not found.", 404}
)

func ErrMissParam(param string) *Error {
	return &Error{
		"miss_param",
		"Miss Param",
		"miss query param: " + param,
		400,
	}
}

type Success struct {
	Title  string 	   `json:"title,omitempty"`
	Data   interface{} `json:"data,omitempty"`

	status int
}

func Response(w http.ResponseWriter, title string, data interface{}) {
	success := &Success{title, data, http.StatusOK}
	w.WriteHeader(success.status)
	w.Header().Set("content-type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(success)
}
