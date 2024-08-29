package slugifyr

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// Slugify converts a string into a slug format, supporting UTF-8, Persian, and English characters.
func Slugify(input string) (string, error) {
	// Normalize input string to remove diacritical marks.
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	normalized, _, err := transform.String(t, input)
	if err != nil {
		return "", err
	}

	// Convert to lowercase
	normalized = strings.ToLower(normalized)

	// Replace spaces and underscores with hyphens
	normalized = strings.ReplaceAll(normalized, " ", "-")
	normalized = strings.ReplaceAll(normalized, "_", "-")

	// Remove non-alphanumeric characters except for hyphens
	reg := regexp.MustCompile(`[^\p{L}\p{N}-]+`)

	normalized = reg.ReplaceAllString(normalized, "")
	normalized = strings.Trim(normalized, "-")

	// Replace multiple hyphens with a single hyphen
	reg = regexp.MustCompile(`-+`)
	normalized = reg.ReplaceAllString(normalized, "-")

	return normalized, nil
}

// isMn checks if a rune is a non-spacing mark.
func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r)
}
