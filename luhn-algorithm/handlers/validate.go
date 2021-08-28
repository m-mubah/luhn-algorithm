package handlers

import (
	"cryptosystem/luhn-algorithm/models"
	"cryptosystem/luhn-algorithm/services"
	"cryptosystem/utilities/web"
	"html/template"
	"log"
	"net/http"
	"os"
)

// Validate - POST. Validates a card number by using luhn algorithm
func Validate(w http.ResponseWriter, req *http.Request) {
	wd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}

	err = req.ParseForm()
	if err != nil {
		log.Panic(err)
	}

	//get the data from the request.form and populate the struct
	formData := new(models.FormData)
	formData.Populate(req)

	var errMessage = ""

	//validate if the form has the correct data
	if !formData.Validate() {
		//return validation errors
		for _, error := range formData.Errors {
			errMessage = errMessage + "<p>" + error + "</p>"
		}

		web.InsertHTMLIntoTemplate(w, req, "luhn-algorithm/_invalid.gohtml", "_invalid", errMessage)

	} else {
		//check if the card is valid and find the checksum
		isValid, checksum := services.ValidateCard(formData.CardNumber)

		if isValid {
			//set the data
			industryIdentifier := services.GetMajorIndustryIdentifier(formData.CardNumber[:1])
			issuer, personalAccountNumber, imgPath := services.GetCardIssuer(formData.CardNumber)

			//initialise view data
			data := models.Card{
				ImgPath:           imgPath,
				CardNumber:        formData.CardNumber,
				MII:               industryIdentifier,
				Issuer:            issuer,
				PersonalAccNumber: personalAccountNumber,
				Checksum:          checksum,
			}

			tmpl := template.Must(template.ParseFiles(wd + "/web/templates/luhn-algorithm/_valid.gohtml"))
			err = tmpl.ExecuteTemplate(w, "_valid", data)
			if err != nil {
				log.Panic(err)
			}
		} else {
			errMessage = "<strong>" + formData.CardNumber + "</strong> is not a valid card number."
			web.InsertHTMLIntoTemplate(w, req, "luhn-algorithm/_invalid.gohtml", "_invalid", errMessage)
		}
	}
}
