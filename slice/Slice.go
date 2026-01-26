package slice

import "fmt"

func MySlice() {
	var a = []int{} // java: list

	a = append(a, 1) // java : list.add
	a = append(a, 2)
	a = append(a, 3, 4, 5)

	// duyệt
	for _, value := range a {
		fmt.Println(value)
	}

	fmt.Println(a)
}
