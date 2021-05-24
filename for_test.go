package main

import (
	"fmt"
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

func TestForWeight(t *testing.T) {
	Equal(t, OrderWeight("103 123 4444 99 2000"), "2000 103 123 4444 99")
	Equal(t, OrderWeight("2000 10003 1234000 44444444 9999 11 11 22 123"), "11 11 2000 10003 22 123 1234000 44444444 9999")
	Equal(t, OrderWeight(" "), "")
	Equal(t, OrderWeight("333333 "), "333333")
}

func TestConsonantValue(t *testing.T) {
	Equal(t, ConsonantValue("a"), (0))
	Equal(t, ConsonantValue("bcd"), (9))
	Equal(t, ConsonantValue("zea"), (26))
	Equal(t, ConsonantValue("az"), (26))
	Equal(t, ConsonantValue("baz"), (26))
	Equal(t, ConsonantValue("aeiou"), (0))
	Equal(t, ConsonantValue("uaoczei"), (29))
	Equal(t, ConsonantValue("abababababfapeifapefijaefaepfjavnefjnfbhwyfnjsifjapnes"), (143))
	Equal(t, ConsonantValue("codewars"), (37))
	Equal(t, ConsonantValue("bup"), (16))
}

func TestFistNonRepeating(t *testing.T) {
	Equal(t, FirstNonRepeating("a"), "a")
	Equal(t, FirstNonRepeating("stress"), "t")
	Equal(t, FirstNonRepeating("moonmen"), "e")

	Equal(t, FirstNonRepeating(""), "")

	Equal(t, FirstNonRepeating("abba"), "")
	Equal(t, FirstNonRepeating("aa"), "")

	Equal(t, FirstNonRepeating("~><#~><"), "#")
	Equal(t, FirstNonRepeating("hello world, eh?"), "w")

	Equal(t, FirstNonRepeating("sTreSS"), "T")
	Equal(t, FirstNonRepeating("Go hang a salami, I'm a lasagna hog!"), ",")
}

func TestMultiple3And5(t *testing.T) {
	Equal(t, Multiple3And5(10), 23)

	// 测试30
	// 3,6,9,12,15,18,21,24,27 = 135
	// 5,10,(15),20,25 = 60
	Equal(t, Multiple3And5(30), 195)
}

func TestValidParentheses(t *testing.T) {
	Equal(t, ValidParentheses("()"), true)
	Equal(t, ValidParentheses(")(()))"), false)
	Equal(t, ValidParentheses("("), false)
	Equal(t, ValidParentheses("(())((()())())"), true)
}

func TestCreateSpiral(t *testing.T) {
	for n := 0; n < 7; n++ {
		t.Run(fmt.Sprintf("n=%d", n), func(t *testing.T) {
			ret := CreateSpiral(n)
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					fmt.Print(ret[i][j])
					fmt.Print(",")
				}
				fmt.Println()
			}
		})
	}
}

func TestDuplicate_count(t *testing.T) {
	Equal(t, duplicate_count("abcde"), 0)
	Equal(t, duplicate_count("abcdea"), 1)
	Equal(t, duplicate_count("abcdeaB11"), 3)
	Equal(t, duplicate_count("indivisibility"), 1)
}

func TestFirstVariationonCaesarCipher(t *testing.T) {
	{
		var sol = []string{"T p", "oc ", "iwl", "yo", ""}
		var u = "S mkx bocod"
		Equal(t, MovingShift(u, 1), sol)

		Equal(t, DemovingShift(sol, 1), u)

	}

	{
		var sol = []string{"J vltasl rlhr ", "zdfog odxr ypw", " atasl rlhr p ", "gwkzzyq zntyhv", " lvz wp!!!"}
		var u = "I should have known that you would have a perfect answer for me!!!"

		Equal(t, MovingShift(u, 1), sol)

		Equal(t, DemovingShift(sol, 1), u)
	}

	{
		var u = "I should have known that you would have a perfect answer for me!!!"
		var sol = []string{"J vltasl rlhr ", "zdfog odxr ypw", " atasl rlhr p ", "gwkzzyq zntyhv", " lvz wp!!!"}

		Equal(t, MovingShift(u, 1), sol)
		Equal(t, DemovingShift(sol, 1), u)
	}
}

func TestPartitions(t *testing.T) {
	Equal(t, Partitions(1), 1)
	Equal(t, Partitions(5), 7)
	Equal(t, Partitions(10), 42)
	Equal(t, Partitions(25), 1958)
}
