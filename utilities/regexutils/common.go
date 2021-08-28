package regexutils

import "regexp"

//contains common regular expressions

var DigitsOnly = regexp.MustCompile("^\\d+$")
