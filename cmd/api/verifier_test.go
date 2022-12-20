package main

import (
	"reflect"
	"testing"
)

func Test_rules_minSize(t *testing.T) {
	table := []struct {
		password string
		value    int
		want     bool
	}{
		{"TesteSenhaForte!123&", 20, true},
		{"TesteSenhaForte!123&", 8, true},
		{"Test!1&", 8, false},
		{"TesteSenhaForte!123&", 0, true},
		{"", 5, false},
		{"", 0, true},
	}

	for _, data := range table {
		result := minSize(data.password, data.value)
		if result != data.want {
			t.Errorf("minSize(%s, %d): %t, want: %t", data.password, data.value, result, data.want)
		}
	}
}

func Test_rules_minUppercase(t *testing.T) {
	table := []struct {
		password string
		value    int
		want     bool
	}{
		{"TesteSenhaForte!123&", 3, true},
		{"TesteSenhaForte!123&", 2, true},
		{"Test!1&", 3, false},
		{"TesteSenhaForte!123&", 0, true},
		{"", 5, false},
		{"", 0, true},
	}

	for _, data := range table {
		result := minUppercase(data.password, data.value)
		if result != data.want {
			t.Errorf("minUppercase(%s, %d): %t, want: %t", data.password, data.value, result, data.want)
		}
	}
}

func Test_rules_minLowercase(t *testing.T) {
	table := []struct {
		password string
		value    int
		want     bool
	}{
		{"TesteSenhaForte!123&", 12, true},
		{"TesteSenhaForte!123&", 10, true},
		{"Test!1&", 5, false},
		{"TesteSenhaForte!123&", 0, true},
		{"", 5, false},
		{"", 0, true},
	}

	for _, data := range table {
		result := minLowercase(data.password, data.value)
		if result != data.want {
			t.Errorf("minLowercase(%s, %d): %t, want: %t", data.password, data.value, result, data.want)
		}
	}
}

func Test_rules_minDigit(t *testing.T) {
	table := []struct {
		password string
		value    int
		want     bool
	}{
		{"TesteSenhaForte!123&", 3, true},
		{"TesteSenhaForte!123&", 1, true},
		{"Test!1&", 5, false},
		{"TesteSenhaForte!123&", 0, true},
		{"", 5, false},
		{"", 0, true},
	}

	for _, data := range table {
		result := minDigit(data.password, data.value)
		if result != data.want {
			t.Errorf("minDigit(%s, %d): %t, want: %t", data.password, data.value, result, data.want)
		}
	}
}

func Test_rules_minSpecialChars(t *testing.T) {
	table := []struct {
		password string
		value    int
		want     bool
	}{
		{"TesteSenhaForte!123&", 2, true},
		{"TesteSenhaForte!123&", 1, true},
		{"Test!1&", 5, false},
		{"TesteSenhaForte!123&", 0, true},
		{"", 5, false},
		{"", 0, true},
	}

	for _, data := range table {
		result := minSpecialChars(data.password, data.value)
		if result != data.want {
			t.Errorf("minSpecialChars(%s, %d): %t, want: %t", data.password, data.value, result, data.want)
		}
	}
}

func Test_rules_noRepeted(t *testing.T) {
	table := []struct {
		password string
		value    int
		want     bool
	}{
		{"TesteSenhaForte!123&", 0, true},
		{"TtesteSenhaForte!123&", 0, true},
		{"TTesteSenhaForte!123&", 0, false},
		{"TesteSenhaForte!123&&", 0, false},
		{"TesteSenhaForte!123123&", 0, true},
		{"TesteSenhaForte!1123&", 0, false},
		{"TesteSenhaForte!!123&", 0, false},
		{"", 0, true},
	}

	for _, data := range table {
		result := noRepeted(data.password)
		if result != data.want {
			t.Errorf("noRepeted(%s, %d): %t, want: %t", data.password, data.value, result, data.want)
		}
	}
}

func Test_rules_verifyRules(t *testing.T) {
	table := []struct {
		verifier Verifier
		want     []string
	}{
		{Verifier{
			Password: "TesteSenhaForte!123&",
			Rules: []Rules{
				{Rule: "minSize", Value: 8},
				{Rule: "minSpecialChars", Value: 2},
				{Rule: "noRepeted", Value: 0},
				{Rule: "minDigit", Value: 4},
			},
		}, []string{"minDigit"}},
		{Verifier{
			Password: "TTesteSenhaForte!123&",
			Rules: []Rules{
				{Rule: "minSize", Value: 25},
				{Rule: "minSpecialChars", Value: 5},
				{Rule: "noRepeted", Value: 0},
				{Rule: "minDigit", Value: 4},
				{Rule: "minUppercase", Value: 6},
				{Rule: "minLowercase", Value: 13},
			},
		}, []string{"minSize", "minSpecialChars", "noRepeted", "minDigit", "minUppercase", "minLowercase"}},
		{Verifier{
			Password: "TesteSenhaForte!123&",
			Rules: []Rules{
				{Rule: "minSize", Value: 8},
				{Rule: "minSpecialChars", Value: 2},
				{Rule: "noRepeted", Value: 0},
				{Rule: "minDigit", Value: 3},
				{Rule: "minUppercase", Value: 3},
				{Rule: "minLowercase", Value: 3},
			},
		}, []string{}},
	}

	for _, data := range table {
		result := data.verifier.verifyRules()
		if !reflect.DeepEqual(result, data.want) {
			t.Errorf("verifyRules(%v, %v): %T %v, want: %T %v", data.verifier.Password, data.verifier.Rules, result, result, data.want, data.want)
		}
	}
}
