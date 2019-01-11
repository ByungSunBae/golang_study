//상수2
package main

import "fmt"

func main() {
	const a, b int = 10, 99
	const c, d, e = true, 0.84, "string1"
	const (
		x, y int16 = 60, 500
		i, k       = "Data", 7796
	)

	fmt.Println(a, b, c, d, e)
	fmt.Println(x, y, i, k)
}
