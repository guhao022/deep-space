package webman
import (
	"net/http"
	"encoding/json"
)

type Error struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`

	Status int    `json:"status"`
}

func (err Error) Error() string {
	return err.Detail
}

func NewError(w http.ResponseWriter, err *Error) {
	w.WriteHeader(err.Status)
	w.Header().Set("content-type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(err)
}


var (
	ErrBadRequest     = &Error{"bad_request", "Bad request", "Request body is not well-formed. It must be JSON.", 400}
	ErrNoData         = &Error{"no_data", "No data", `Key "data" in the top level of the JSON document is missing or contains no data`, 422}
	ErrNotFound       = &Error{"not_found", "Not found", "Route not found.", 404}
)

func ErrMissParam(param string) *Error {
	return &Error{
		"miss_param",
		"Miss Param",
		"miss query param: " + param,
		http.StatusBadRequest,
	}
}

func ErrForbidden(detail string) *Error {
	return &Error{
		"forbidden",
		"Forbidden",
		"Forbidden: " + detail,
		http.StatusForbidden,
	}
}

func ErrInternalServer(detail string) *Error {
	return &Error{
		"internal_server_error",
		"Internal Server Error",
		detail,
		http.StatusInternalServerError,
	}
}

func ErrUnauthorized(detail string) *Error {
	return &Error{
		"unauthorized",
		"Unauthorized",
		detail,
		http.StatusUnauthorized,
	}
}

type Success struct {
	ID  string 	   `json:"title,omitempty"`
	Data   interface{} `json:"data,omitempty"`

	Status int	`json:"status"`
}

func Response(w http.ResponseWriter, id string, data interface{}) {
	success := &Success{id, data, http.StatusOK}
	w.WriteHeader(success.Status)
	w.Header().Set("content-type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(success)
}

