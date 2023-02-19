package pkg_test

import (
	"testing"

	"github.com/RaoTheVeloper/math-functions/pkg"
)

type testCase struct {
	a, b, result int
}

var tests = []testCase{
	{a: 10, b: 5, result: 5},
	{a: 10, b: 0, result: 10},
	{a: 10, b: 1, result: 1},
	{a: 10, b: 2, result: 2},
	{a: 12, b: 18, result: 6},
	{a: -5, b: 5, result: 5},
	{a: -5, b: -5, result: 5},
	{a: 5, b: -5, result: 5},
}

func TestGCD(t *testing.T) {
	for _, test := range tests {
		if result := pkg.GCD(test.a, test.b); result != test.result {
			t.Errorf("GCD(%d, %d) = %d; want %d", test.a, test.b, result, test.result)
		}
	}
}
