package gengo

import "strings"

// toCamel make CamelCase. A word, snake_case and camelCase is applicable.
func toCamel(name string) string {
	converted := ""

	ar := strings.Split(name, "_")
	if len(ar) == 0 { // snake_case
		converted = strings.ToUpper(name[:1]) + name[1:]
	} else {
		for _, w := range ar {
			converted += strings.ToUpper(w[:1]) + w[1:]
		}
	}
	return converted
}

// toSingular make plural to singular by simply taking last `s`.
func toSingular(name string) string {
	lastCharIndex := len(name) - 1
	if name[lastCharIndex] == 's' {
		return name[:lastCharIndex]
	}
	return name
}
