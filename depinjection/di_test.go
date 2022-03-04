package depinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Gert")

	got := buffer.String()
	want := "Hello, Gert"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
