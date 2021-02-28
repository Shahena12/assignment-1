package main

import (
	"fmt"
	"unicode"
)

func main() {

	fmt.Println(isValid("pass"))     
	fmt.Println(isValid("password")) 
	fmt.Println(isValid("Password")) 
	fmt.Println(isValid("P@ssword")) 
	fmt.Println(isValid("P@ssw0rd")) 
}

func isValid(s string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(s) >= 7 {
		hasMinLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}