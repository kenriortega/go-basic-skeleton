package main

import (
	"reflect"
	"testing"
)

func TestGreeting(t *testing.T) {
	want := "Hello world"
	got := Greeting()

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected %v got %v", want, got)
	}

}
