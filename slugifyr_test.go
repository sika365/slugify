package main

import "testing"

func TestSlugify(t *testing.T) {
	// Define test cases.
	testCases := []struct {
		input    string
		expected string
	}{
		{"Hello World", "hello-world"},
		{"This is a Test!", "this-is-a-test"},
		{"GoLang is awesome", "golang-is-awesome"},
		{"Multiple     Spaces", "multiple-spaces"},
		{"Special # Characters !", "special-characters"},
		{"این یک تست است", "این-یک-تست-است"},
		{"تست، بررسی", "تست-بررسی"},
		{"فارسی &*|false| *!123", "فارسی-false-123"},
		{"سلام_دنیا", "سلام-دنیا"},
		{"Hello{}-سلام", "hello-سلام"},
		{"GoLang و گرامر", "golang-و-گرامر"},
	}

	for _, tc := range testCases {
		output, _ := Slugify(tc.input)
		if output != tc.expected {
			t.Errorf("Slugify(%q) = %q; want %q", tc.input, output, tc.expected)
		}
	}
}
