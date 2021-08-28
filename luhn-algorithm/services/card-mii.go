package services

import "strconv"

// GetMajorIndustryIdentifier - returns the industry to which the card belongs to.
func GetMajorIndustryIdentifier(cardNumber string) string {
	number, _ := strconv.Atoi(cardNumber)
	result := "unknown"

	switch number {
	case 0:
		result = "ISO/TC 68"
		break
	case 1, 2:
		result = "Airlines"
		break
	case 3:
		result = "Travel and entertainment and banking/financial"
		break
	case 4, 5:
		result = "Banking and financial"
		break
	case 6:
		result = "Merchandising and banking/financial"
		break
	case 7:
		result = "Petroleum"
		break
	case 8:
		result = "Healthcare and telecommunications"
		break
	case 9:
		result = "National assignment"
		break
	default:
		break
	}

	return result
}
