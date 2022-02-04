package main

import (
	"fizzbuzz/pkg/fizzbuzz"
	"fmt"
)

// Entry point of this program
func main() {
	for _, v := range fizzbuzz.UpToLength(20) {
		fmt.Println(v)
	}
}
