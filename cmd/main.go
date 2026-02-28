package main

import (
	"27_02_2026/perror"
	"fmt"
)

func NotNegative(x int) (int, error) {
	if x < 0 {
		return 0, perror.New("Negative Number Does not proper Divisor", 3233)
	}
	return x * x, nil
}
func main() {
	val, err := NotNegative(1)

	if err != nil {
		fmt.Println("error occured")
		return
	}
	fmt.Println(val)
	fmt.Println("Hello pulock just i am test")
}
