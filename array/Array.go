package array

import "fmt"

func Array() {

	// array 1D

	fmt.Println("============= array 1D ============")
	var a [10]int
	for i := 0; i < 10; i++ {
		a[i] = i
	}

	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	fmt.Println(len(a))
	fmt.Println(a[4])

	// matrix - array 2D

	fmt.Println("============= matrix ============")
	var arr2D [10][10]int

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			arr2D[i][j] = i + j
		}
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Print(arr2D[i][j], " ")
		}
		fmt.Println()

	}

	// output
	//============= array 1D ============
	//	0 1 2 3 4 5 6 7 8 9
	//10
	//4
	//============= matrix ============
	//	0 1 2 3 4 5 6 7 8 9
	//1 2 3 4 5 6 7 8 9 10
	//2 3 4 5 6 7 8 9 10 11
	//3 4 5 6 7 8 9 10 11 12
	//4 5 6 7 8 9 10 11 12 13
	//5 6 7 8 9 10 11 12 13 14
	//6 7 8 9 10 11 12 13 14 15
	//7 8 9 10 11 12 13 14 15 16
	//8 9 10 11 12 13 14 15 16 17
	//9 10 11 12 13 14 15 16 17 18
}
