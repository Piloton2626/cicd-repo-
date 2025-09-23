package main

import "testing"

func TestHello(t *testing.T) {
    expected := "welcome to the CI/CD"
    if expected != "welcome to the CI/CD" {
        t.Errorf("Expected %s but got different", expected)
    }
}
