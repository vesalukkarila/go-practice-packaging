package main

import "testing"

func TestValidGetName(t *testing.T) {
	expected := "vesa"
	name, err := getName([]string{"./main", "vesa"}) //list of strings, notice syntax
	if err != nil {
		t.Errorf("expected err nil, received %v", err)
	}

	if expected != name {
		t.Errorf("expected %s, received %s", expected, name)
	}
}

func TestInvalidGetName(t *testing.T) {
	expected := ""
	name, err := getName([]string{"./main"}) //list of strings, notice syntax
	if err == nil {
		t.Errorf("expected err, received %v", err)
	}

	if expected != name {
		t.Errorf("expected %s, received %s", expected, name)
	}
}
