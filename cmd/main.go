package main

import (
	"fmt"

	"github.com/RaoTheVeloper/math-functions/pkg"
)

func main() {
	fmt.Println("Welcome to math functions!")

	fmt.Printf("GCD(%d, %d) = %d", 12, 18, pkg.GCD(12, 18))
}
