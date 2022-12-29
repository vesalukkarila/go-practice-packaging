package greet

import "testing"

func TestValidGreet(t *testing.T) {
	expected := "hello Vesa"
	greeting, err := Greet("Vesa")
	if err != nil {
		t.Errorf("expected err nil, received %v", err)
	}
	if expected != greeting {
		t.Errorf("expected %s, received %s", expected, greeting)
	}
}

func TestInvalidGreed(t *testing.T) {
	expected := ""
	greeting, err := Greet("")
	if err == nil {
		t.Errorf("expected err not nil, received %v", err)
	}
	if expected != greeting {
		t.Errorf("expected %s, received %s", expected, greeting)
	}
}
