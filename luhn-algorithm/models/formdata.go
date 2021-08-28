package models

import (
	"cryptosystem/utilities/regexutils"
	"net/http"
)

type FormData struct {
	CardNumber string
	Errors     map[string]string
}

// Populate - populates the form with form data from the request
func (formData *FormData) Populate(req *http.Request) {
	formData.CardNumber = req.FormValue("cardNumber")
}

// Validate - performs form validation
func (formData *FormData) Validate() bool {
	digitsOnly := regexutils.DigitsOnly.Match([]byte(formData.CardNumber))
	formData.Errors = make(map[string]string)

	if digitsOnly == false {
		formData.Errors["CardNumber"] = "Please enter digits only!"
	}

	return len(formData.Errors) == 0
}
