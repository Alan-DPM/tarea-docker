package main

import "testing"

func TestSumar(t *testing.T) {
	if Sumar(5, 5) != 10 {
		t.Error("Error: Se esperaba 5")
	}
}
