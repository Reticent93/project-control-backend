package helpers

import (
	"fmt"
	"net/http"
	"project-control-backend/internal/config"
	"runtime/debug"
)

var app *config.AppConfig

//ServerError is a helper function that writes an error message and stack trace to the errorLog, then sends a generic 500 Internal Server Error response to the user.
func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	fmt.Println(trace)
	if err != nil {
		return
	}
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	
}
