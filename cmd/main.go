package main

import (
	"fmt"

	"github.com/RaoTheVeloper/math-functions/pkg"
)

func main() {
	fmt.Println("Hello, World!")

	fmt.Printf("GCD(%d, %d) = %d", 10, 5, pkg.GCD(10, 5))
}
