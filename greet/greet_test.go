package greet

import "testing"

func TestGreet(t *testing.T) {
	expected := "hello Vesa"
	received := Greet("Vesa")
	if expected != received {
		t.Errorf("expected %s, received %s", expected, received)
	}
}
