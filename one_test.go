package main

import (
	"testing"

	. "github.com/go-playground/assert/v2"
)

func TestCreatePhoneNumber(t *testing.T) {
	result := CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	Equal(t, result, "(123) 456-7890")
}
