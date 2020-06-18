package format

import "testing"

func TestName(t *testing.T) {
	want := "project1"
	if got := Name(" Project1 "); got != want {
		t.Errorf("Name() = %q, want =%q", got, want)
	}
}
