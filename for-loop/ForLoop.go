package for_loop

import "fmt"

func ForLoop() {
	i := 1

	for i <= 3 {
		fmt.Println(i, "for loop with Long")
		i++
	}

	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}

	fmt.Println()

	fmt.Println("Số chẵn < 10:")

	for i := range 10 {
		if i%2 == 0 {
			fmt.Print(i, " ")
		}
	}

	//output
	//1 for loop with Long
	//2 for loop with Long
	//3 for loop with Long
	//1 2 3 4 5 6 7 8 9 10
	//Số chẵn < 10:
	//0 2 4 6 8
}
