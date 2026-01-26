package constant

import (
	"fmt"
)

const s string = "constant"

func Constant() {

	fmt.Println("s: " + s)

	const n = 500000000
	fmt.Println("n: ", n)

	const d = 3e20 / n
	fmt.Println("d: ", d)

	fmt.Println("int64(d): ", int64(d))

	//output
	//s: constant
	//n:  500000000
	//d:  6e+11
	//int64(d):  600000000000
}
