package main

import (
	"reflect"
	"testing"
)

func TestMinSize(t *testing.T) {
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
		result := MinSize(data.password, data.value)
		if result != data.want {
			t.Errorf("MinSize(%s, %d): %t, want: %t", data.password, data.value, result, data.want)
		}
	}
}

func TestMinUppercase(t *testing.T) {
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
		result := MinUppercase(data.password, data.value)
		if result != data.want {
			t.Errorf("MinUppercase(%s, %d): %t, want: %t", data.password, data.value, result, data.want)
		}
	}
}

func TestMinLowercase(t *testing.T) {
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
		result := MinLowercase(data.password, data.value)
		if result != data.want {
			t.Errorf("MinLowercase(%s, %d): %t, want: %t", data.password, data.value, result, data.want)
		}
	}
}

func TestMinDigit(t *testing.T) {
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
		result := MinDigit(data.password, data.value)
		if result != data.want {
			t.Errorf("MinDigit(%s, %d): %t, want: %t", data.password, data.value, result, data.want)
		}
	}
}

func TestMinSpecialChars(t *testing.T) {
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
		result := MinSpecialChars(data.password, data.value)
		if result != data.want {
			t.Errorf("MinSpecialChars(%s, %d): %t, want: %t", data.password, data.value, result, data.want)
		}
	}
}

func TestNoRepeted(t *testing.T) {
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
		result := NoRepeted(data.password)
		if result != data.want {
			t.Errorf("NoRepeted(%s, %d): %t, want: %t", data.password, data.value, result, data.want)
		}
	}
}

func TestVarifyRules(t *testing.T) {
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
		result := data.verifier.VarifyRules()
		if !reflect.DeepEqual(result, data.want) {
			t.Errorf("VarifyRules(%v, %v): %T %v, want: %T %v", data.verifier.Password, data.verifier.Rules, result, result, data.want, data.want)
		}
	}
}
