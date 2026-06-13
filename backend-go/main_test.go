package main

import "testing"

func TestSumar(t *testing.T) {
	if Sumar(2, 3) != 5 {
		t.Error("Error: Se esperaba 5")
	}
}
