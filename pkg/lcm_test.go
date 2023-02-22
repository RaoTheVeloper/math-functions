package pkg_test

import (
	"testing"

	"github.com/RaoTheVeloper/math-functions/pkg"
)

type lcmTestCase struct {
	a, b, result int
}

var lcmTests = []lcmTestCase{
	{a: 10, b: 5, result: 10},
	{a: 10, b: 0, result: 0},
	{a: 10, b: 1, result: 10},
	{a: 10, b: 2, result: 10},
	{a: 12, b: 18, result: 36},
	{a: -5, b: 5, result: 5},
	{a: -5, b: -5, result: 5},
	{a: 5, b: -5, result: 5},
	{a: 0, b: 0, result: 0},
}

func TestLCM(t *testing.T) {
	for _, test := range lcmTests {
		if result := pkg.LCM(test.a, test.b); result != test.result {
			t.Errorf("LCM(%d, %d) = %d; want %d", test.a, test.b, result, test.result)
		}
	}
}

func TestLCMPrimeFactorization(t *testing.T) {
	for _, test := range lcmTests {
		if result := pkg.LCMPrimeFactorization(test.a, test.b); result != test.result {
			t.Errorf("LCMPrimeFactorization(%d, %d) = %d; want %d", test.a, test.b, result, test.result)
		}
	}
}
