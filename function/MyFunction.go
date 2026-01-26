package function

import "fmt"

func Tong(a int, b int) (int, string) {
	if a == 0 && b == 0 {
		sum := 0
		message := "Example Error"
		return sum, message
	}

	return a + b, "Accepted!"
}

func MyFunction() {
	var a, b int

	fmt.Println("Nhap a va b: ")
	fmt.Scan(&a, &b)

	sum, msg := Tong(a, b)
	fmt.Println(sum, msg)

}
