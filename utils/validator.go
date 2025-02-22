package utils

import "regexp"

func ValidatePhoneNumber(phone string) bool {
	pattern := `^(?:\+?88)?01[3-9]\d{8}$`
	match, _ := regexp.MatchString(pattern, phone)
	return match
}
