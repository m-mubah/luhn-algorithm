package web

import (
	"github.com/gorilla/csrf"
	"html/template"
	"log"
	"net/http"
	"os"
)

// functions to generate templates in a more reuseable manner

//GenerateTemplate - builds and returns a given http template with the given data.
func GenerateTemplate(w http.ResponseWriter, req *http.Request, templateName string, data interface{}) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	//compile layouts
	tmpl := template.Must(template.ParseGlob(wd + "/web/templates/layouts/*.gohtml"))
	//compile page
	template.Must(tmpl.ParseFiles(wd + "/web/templates/" + templateName))

	err = tmpl.ExecuteTemplate(w, "main", TemplateData{
		CsrfField: csrf.TemplateField(req),
		Data:      data,
	})

	if err != nil {
		log.Fatal(err)
	}
}

// InsertHTMLIntoTemplate - inserts a string with valid html into a given template
func InsertHTMLIntoTemplate(w http.ResponseWriter, req *http.Request, templatePath, templateName, errMessage string) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	tmpl := template.Must(template.ParseFiles(wd + "/web/templates/" + templatePath))
	err = tmpl.ExecuteTemplate(w, templateName, template.HTML(errMessage))
	if err != nil {
		log.Panic(err)
	}
}

type TemplateData struct {
	CsrfField template.HTML
	Data      interface{}
}
