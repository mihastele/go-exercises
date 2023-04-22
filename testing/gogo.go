package main

import "regexp"

func IsValidEmail(addr string) bool {
	return regexp.MustCompile(`.+@.+\..+`).MatchString(addr)
}
