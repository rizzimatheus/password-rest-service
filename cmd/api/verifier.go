package main

import (
	"fmt"
	"regexp"
	"unicode"
)

type Verifier struct {
	Password string  `json:"password"`
	Rules    []Rules `json:"rules"`
}

type Rules struct {
	Rule  string `json:"rule"`
	Value int    `json:"value"`
}

// verifyRules returns which 'rules' have not been satisfied by the 'password'
func (v *Verifier) verifyRules() []string {
	nomatch := []string{}
	for _, rule := range v.Rules {
		switch rule.Rule {
		case "minSize":
			if !minSize(v.Password, rule.Value) {
				nomatch = append(nomatch, "minSize")
			}
		case "minUppercase":
			if !minUppercase(v.Password, rule.Value) {
				nomatch = append(nomatch, "minUppercase")
			}
		case "minLowercase":
			if !minLowercase(v.Password, rule.Value) {
				nomatch = append(nomatch, "minLowercase")
			}
		case "minDigit":
			if !minDigit(v.Password, rule.Value) {
				nomatch = append(nomatch, "minDigit")
			}
		case "minSpecialChars":
			if !minSpecialChars(v.Password, rule.Value) {
				nomatch = append(nomatch, "minSpecialChars")
			}
		case "noRepeted":
			if !noRepeted(v.Password) {
				nomatch = append(nomatch, "noRepeted")
			}
		default:
			nomatch = append(nomatch, fmt.Sprintf("Error: '%s' is an invalid rule", rule.Rule))
		}
	}

	return nomatch
}

// minSize checks if the 'password' has at least 'value' characters
func minSize(password string, value int) bool {
	return len(password) >= value
}

// minUppercase checks if the 'password' has at least 'value' uppercase characters
func minUppercase(password string, value int) bool {
	count := 0
	for _, r := range password {
		if unicode.IsUpper(r) {
			count++
			if count >= value {
				return true
			}
		}
	}
	return value == 0
}

// minLowercase checks if the 'password' has at least 'value' lowercase characters
func minLowercase(password string, value int) bool {
	count := 0
	for _, r := range password {
		if unicode.IsLower(r) {
			count++
			if count >= value {
				return true
			}
		}
	}
	return value == 0
}

// minDigit checks if the 'password' has at least 'value' digits
func minDigit(password string, value int) bool {
	count := 0
	for _, r := range password {
		if unicode.IsDigit(r) {
			count++
			if count >= value {
				return true
			}
		}
	}
	return value == 0
}

// minSpecialChars checks if the 'password' has at least 'value' special characters. Special characters: "!@#$%^&*()-+\/{}[]"
func minSpecialChars(password string, value int) bool {
	re := regexp.MustCompile(`[!@#$%^&*()-+\/{}\[\]]`)
	count := len(re.FindAll([]byte(password), -1))
	return count >= value
}

// noRepeted checks if the 'password' has no repeated characters in sequence. ("aab": false, "aba": true)
func noRepeted(password string) bool {
	if len(password) <= 1 {
		return true
	}

	r := password[0]
	for i := 1; i < len(password); i++ {
		if r == password[i] {
			return false
		}

		r = password[i]
	}

	return true
}
