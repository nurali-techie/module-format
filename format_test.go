package format

import "testing"

func TestName(t *testing.T) {
	want := "Nurali"
	if got := Name(" nurali "); got != want {
		t.Errorf("Name() = %q, want =%q", got, want)
	}
}
