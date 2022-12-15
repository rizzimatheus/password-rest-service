package main

import (
	"regexp"
	"unicode"
)

type Rules struct {
	Rule  string `json:"rule"`
	Value int    `json:"value"`
}

func minSize(password string, value int) bool {
	return len(password) >= value
}

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
	return false
}

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
	return false
}

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
	return false
}

func minSpecialChars(password string, value int) bool {
	re := regexp.MustCompile(`[!@#$%^&*()-+\/{}\[\]]`)
	count := len(re.FindAll([]byte(password), -1))
	return count >= value
}

func noRepeted(password string) bool {
	if len(password) <= 1 {
		return true
	}

	r := password[0]
	for i := 1 ; i < len(password); i++ {
		if r == password[i] {
			return false
		}

		r = password[i]
	}

	return true
}