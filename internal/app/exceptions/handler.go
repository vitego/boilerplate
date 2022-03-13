package exceptions

import (
	"github.com/vitego/router/httperror"
	"github.com/vitego/router/response"
	"net/http"
	"reflect"
)

type Handler struct{}

// Render a panic exception into an HTTP Response.
func (Handler) Render(w http.ResponseWriter, r *http.Request, err interface{}) {
	status := http.StatusInternalServerError
	text := err
	code := 0

	if reflect.ValueOf(err).Type().String() == "httperror.Error" {
		status = err.(httperror.Error).Status
		text = err.(httperror.Error).Error()
		code = err.(httperror.Error).Code
	}

	response.Error(w, status, code, text)
}
