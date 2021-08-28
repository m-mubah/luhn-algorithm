package handlers

import (
	"cryptosystem/utilities/web"
	"net/http"
)

type TemplateData struct {
	CsrfField interface{}
	Data      interface{}
}

func Index(w http.ResponseWriter, req *http.Request) {
	web.GenerateTemplate(w, req, "luhn-algorithm/index.gohtml", "")
}
