package main

import (
	"fmt"
	"regexp"
)

var EmailExpr *regexp.Regexp

func init() {
	compiled, ok := regexp.Compile(`.+@.+\..+`)
	if ok != nil {
		panic("failed to compile regular expression")
	}
	EmailExpr = compiled
	fmt.Println("regular expression compiled successfully")
}

func isValidEmail(email string) bool {
	return EmailExpr.MatchString(email)
}

func main() {
	fmt.Println(isValidEmail("invalid"))
	fmt.Println(isValidEmail("valid@valid.com"))
	fmt.Println(isValidEmail("invalid@example"))
}
