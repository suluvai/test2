package subhello

import (
	"testing"
)

func TestSubhello(t *testing.T) {
	want := "Hello, world."
	if got := Subhello(); got != want {
		t.Errorf("Subhello() = %q, want %q", got, want)
	}
}

func TestSubproverb(t *testing.T) {
	want := "Concurrency is not parallelism."
	if got := Subpoverb(); got != want {
		t.Errorf("Subproverb() = %q, want %q", got, want)
	}
}
