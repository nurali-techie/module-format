package format

import "strings"

// Name returns given name after triming spaces and making it title case.
func Name(name string) string {
	name = strings.TrimSpace(name)
	name = strings.Title(name)
	return name
}
