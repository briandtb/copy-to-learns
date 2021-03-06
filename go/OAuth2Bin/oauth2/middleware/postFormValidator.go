package middleware

import (
	"net/http"
	"oauth2bin/oauth2/utils"
)

// PostFormValidator verifies if a request has method POST and content-type
// "application/x-www-form-urlencoded". This is turned into a middleware since
// it is a common task in OA2B.
//
// VisualError: boolean which determines whether to present a visual (HTML)
// or textual (JSON) error in case the request doesn't satisfy the above conditions
type PostFormValidator struct {
	VisualError bool
}

// NewPostFormValidator returns a new instance of PostFormValidator
func NewPostFormValidator(visualError bool) PostFormValidator {
	return PostFormValidator{VisualError: visualError}
}

// Handle implements the Middleware interface
// and performs the above mentioned job
func (pfv PostFormValidator) Handle(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			title := "Method Not Allowed"
			desc := r.Method + " not allowed"
			pfv.presentError(w, r, http.StatusMethodNotAllowed, title, desc)
			return
		}

		contentType := r.Header.Get("Content-Type")
		if contentType != "application/x-www-form-urlencoded" {
			title := "Bad Request"

			var desc string
			if contentType == "" {
				desc = "Expecting content type"
			} else {
				desc = "Content type not allowed: " + contentType
			}

			pfv.presentError(w, r, http.StatusBadRequest, title, desc)
			return
		}

		handler.ServeHTTP(w, r)
	}
}

func (pfv PostFormValidator) presentError(w http.ResponseWriter, r *http.Request, status int, title, desc string) {
	if pfv.VisualError {
		utils.ShowError(w, r, status, title, desc)
	} else {
		utils.ShowJSONError(w, r, status, utils.RequestError{
			Error: title,
			Desc:  desc,
		})
	}
}
