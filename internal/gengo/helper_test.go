package gengo

import "testing"

func Test_toCamel(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{name: "SnakeCase", in: "snake_case", want: "SnakeCase"},
		{name: "CamelCase", in: "camelCase", want: "CamelCase"},
		{name: "Word", in: "word", want: "Word"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toCamel(tt.in); got != tt.want {
				t.Errorf("toCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toSingular(t *testing.T) {

	tests := []struct {
		name string
		in   string
		want string
	}{
		{name: "Plural", in: "names", want: "name"},
		{name: "Singular", in: "name", want: "name"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toSingular(tt.in); got != tt.want {
				t.Errorf("toSingular() = %v, want %v", got, tt.want)
			}
		})
	}
}
