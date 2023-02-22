package pkg_test

import (
	"reflect"
	"testing"

	"github.com/RaoTheVeloper/math-functions/pkg"
)

type primeFactorizationTestCase struct {
	n      int
	result []int
}

var primeFactorizationTests = []primeFactorizationTestCase{
	{n: 10, result: []int{2, 5}},
	{n: 0, result: []int{}},
	{n: 1, result: []int{}},
	{n: 2, result: []int{2}},
	{n: 12, result: []int{2, 2, 3}},
	{n: -5, result: []int{5}},
	{n: -10, result: []int{2, 5}},
}

func TestPrimeFactorization(t *testing.T) {
	for _, test := range primeFactorizationTests {
		if result := pkg.PrimeFactorization(test.n); !reflect.DeepEqual(result, test.result) {
			t.Errorf("PrimeFactorization(%d) = %v; want %v", test.n, result, test.result)
		}
	}
}
