package simple_factory

import (
	"testing"
	"fmt"
)

func TestBird_Say(t *testing.T) {
	animal := NewAnimal(1)
	s := animal.Say("Tom1")
	if s != "bird" {
		t.Fatal("Type1 test fail")
	}
}

func TestDog_Say(t *testing.T) {
	animal := NewAnimal(2)
	s := animal.Say("Tom2")
	if s != "dog" {
		t.Fatal("Type1 test fail")
	}
}
