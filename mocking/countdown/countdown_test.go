package countdown

import (
	"bytes"
	"testing"
)

func TestStart(t *testing.T) {
	buffer := &bytes.Buffer{}

	Start(buffer)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
