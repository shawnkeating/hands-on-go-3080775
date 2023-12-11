// types/conversion/begin/main.go
package main

import "fmt"

func main() {
	// declare float and convert it to an unsigned int
	var a = 1.5676
	var b = uint(a)
	fmt.Printf("%v=%T, %v=%T", a, a, b, b)
}
