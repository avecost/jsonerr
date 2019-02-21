package jsonerr

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Error hold the information about an error,
// including metadata abouts its JSON structure
type Error struct {
	HTTPCode int    `json:"-"`
	Code     int    `json:"code,omitempty"`
	Message  string `json:"message"`
}

// Errors hold the information about errors,
// including metadata abouts its JSON structure
type Errors struct {
	HTTPCode int               `json:"-"`
	Code     int               `json:"code,omitempty"`
	Messages map[string]string `json:"messages"`
}

// JSONError function is similar to http.Error, but
// the response body in JSON
func JSONError(w http.ResponseWriter, e Error) {
	data := struct {
		Err Error `json:"error"`
	}{e}

	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.HTTPCode)
	fmt.Fprint(w, string(b))
}

// JSONErrors function is similar to http.Error, but
// the response body in JSON
func JSONErrors(w http.ResponseWriter, e Errors) {
	data := struct {
		Err Errors `json:"error"`
	}{e}

	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.HTTPCode)
	fmt.Fprint(w, string(b))
}

// Example usage
// func displayError(w http.ResponseWriter, r *http.Request) {
// e := Error{
// 	HTTPCode: http.StatusForbidden,
// 	Code: 123,
// 	Message: "An Error Occured",
// }

// JSONError(w, e)
// }
