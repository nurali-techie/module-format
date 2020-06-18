package format

import "strings"

func Name(name string) string {
	name = strings.TrimSpace(name)
	name = strings.Title(name)
	return name
}
