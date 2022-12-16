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

func (v *Verifier) VarifyRules() []string {
	nomatch := []string{}
	for _, rule := range v.Rules {
		switch rule.Rule {
		case "minSize":
			if !MinSize(v.Password, rule.Value) {
				nomatch = append(nomatch, "minSize")
			}
		case "minUppercase":
			if !MinUppercase(v.Password, rule.Value) {
				nomatch = append(nomatch, "minUppercase")
			}
		case "minLowercase":
			if !MinLowercase(v.Password, rule.Value) {
				nomatch = append(nomatch, "minLowercase")
			}
		case "minDigit":
			if !MinDigit(v.Password, rule.Value) {
				nomatch = append(nomatch, "minDigit")
			}
		case "minSpecialChars":
			if !MinSpecialChars(v.Password, rule.Value) {
				nomatch = append(nomatch, "minSpecialChars")
			}
		case "noRepeted":
			if !NoRepeted(v.Password) {
				nomatch = append(nomatch, "noRepeted")
			}
		default:
			nomatch = append(nomatch, fmt.Sprintf("Error: '%s' is an invalid rule", rule.Rule))
		}
	}

	return nomatch
}

func MinSize(password string, value int) bool {
	return len(password) >= value
}

func MinUppercase(password string, value int) bool {
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

func MinLowercase(password string, value int) bool {
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

func MinDigit(password string, value int) bool {
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

func MinSpecialChars(password string, value int) bool {
	re := regexp.MustCompile(`[!@#$%^&*()-+\/{}\[\]]`)
	count := len(re.FindAll([]byte(password), -1))
	return count >= value
}

func NoRepeted(password string) bool {
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
