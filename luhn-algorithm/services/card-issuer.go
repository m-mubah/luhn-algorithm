package services

import (
	"cryptosystem/utilities/regexutils"
	"log"
	"regexp"
)

// GetCardIssuer - gets the issuer of the card and personal account number by evaluating regex.
//Also returns the image path for the card issuer.
func GetCardIssuer(cardNumber string) (string, string, string) {
	var expressions = []*regexutils.Expression{
		{
			Regexp: regexp.MustCompile("^5[1-5][0-9]{14}$"),
			Action: func() (error, string) {
				return nil, "Master Card"
			},
		},
		{
			Regexp: regexp.MustCompile("^4[0-9]{12}(?:[0-9]{3})?$"),
			Action: func() (error, string) {
				return nil, "Visa"
			},
		},
		{
			Regexp: regexp.MustCompile("^3[47][0-9]{13}$"),
			Action: func() (error, string) {
				return nil, "American Express"
			},
		},
		{
			Regexp: regexp.MustCompile("^(?:2131|1800|35\\d{3})\\d{11}$"),
			Action: func() (error, string) {
				return nil, "JCB"
			},
		},
		{
			Regexp: regexp.MustCompile("^3(?:0[0-5]|[68][0-9])[0-9]{11}$"),
			Action: func() (error, string) {
				return nil, "Diner's Club"
			},
		},
		{
			Regexp: regexp.MustCompile("^6(?:011|5[0-9]{2})[0-9]{12}$"),
			Action: func() (error, string) {
				return nil, "Discover"
			},
		},
	}

	matched, personalAccountNumber, imgPath := "unknown", "unknown", ""

	if action, ok := regexutils.Match(expressions, cardNumber); ok {
		err, cardType := action()
		if err != nil {
			log.Panic(err)
		}

		matched = cardType
	}

	switch matched {
	case "unknown":
		break
	case "Master Card":
		personalAccountNumber = cardNumber[len(cardNumber)-10 : len(cardNumber)-1]
		imgPath = "/static/svg/card-vendors/master-card.svg"
		break
	case "Visa":
		personalAccountNumber = cardNumber[len(cardNumber)-10 : len(cardNumber)-1]
		imgPath = "/static/svg/card-vendors/visa.svg"
		break
	case "JCB":
		personalAccountNumber = cardNumber[len(cardNumber)-10 : len(cardNumber)-1]
		imgPath = "/static/svg/card-vendors/jcb.svg"
		break
	case "Discover":
		personalAccountNumber = cardNumber[len(cardNumber)-10 : len(cardNumber)-1]
		imgPath = "/static/svg/card-vendors/discover.svg"
		break
	case "American Express":
		personalAccountNumber = cardNumber[len(cardNumber)-9 : len(cardNumber)-1]
		imgPath = "/static/svg/card-vendors/american-express.svg"
		break
	case "Diner's Club":
		personalAccountNumber = cardNumber[len(cardNumber)-8 : len(cardNumber)-1]
		imgPath = "/static/svg/card-vendors/diners-club.svg"
		break
	}

	return matched, personalAccountNumber, imgPath
}
