package main

import (
	"testing"

	. "github.com/go-playground/assert/v2"
)

func TestCreatePhoneNumber(t *testing.T) {
	result := CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	Equal(t, result, "(123) 456-7890")
}

func TestDivisor(t *testing.T) {
	Equal(t, Divisors(5), 2)
	Equal(t, Divisors(4), 3)
	Equal(t, Divisors(12), 6)
	Equal(t, Divisors(30), 8)
}

func TestDigPow(t *testing.T) {
	Equal(t, DigPow(89, 1), 1)
	Equal(t, DigPow(92, 1), -1)
	Equal(t, DigPow(695, 2), 2)
	Equal(t, DigPow(46288, 3), 51)
}

func TestSplitString(t *testing.T) {
	Equal(t, SplitStrings("abc"), []string{"ab", "c_"})
	Equal(t, SplitStrings("a"), []string{"a_"})
	Equal(t, SplitStrings("dawsd"), []string{"da", "ws", "d_"})
	Equal(t, SplitStrings("awsaws"), []string{"aw", "sa", "ws"})
	Equal(t, SplitStrings("普通话"), []string{"普通", "话_"})
}

func TestMakeTheDeadfishSwim(t *testing.T) {
	Equal(t, MakeTheDeadfishSwim("ooo"), []int{0, 0, 0})
	Equal(t, MakeTheDeadfishSwim("ioioio"), []int{1, 2, 3})
	Equal(t, MakeTheDeadfishSwim("idoiido"), []int{0, 1})
	Equal(t, MakeTheDeadfishSwim("isoisoiso"), []int{1, 4, 25})
	Equal(t, MakeTheDeadfishSwim("codewars"), []int{0})
}
