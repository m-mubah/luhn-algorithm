package regexutils

import "regexp"

// Expression - struct used for matching multiple regexes. Holds a regex and an action
type Expression struct {
	Regexp *regexp.Regexp
	Action func() (error, string)
}

// Match - given a slice of expressions and string, executes the given action if the regex matches. Otherwise returns false.
func Match(expressions []*Expression, str string) (func() (error, string), bool) {
	for _, expression := range expressions {
		if expression.Regexp.MatchString(str) {
			return expression.Action, true
		}
	}
	return nil, false
}
