package _map

import "fmt"

func MyMap() {

	//create
	mp := make(map[string]int) // java: Map<String, Integer> mp = new HashMap<>();

	// put
	mp["Long"] = 1
	mp["Pham"] = 2

	//get
	var val1 = mp["Long"]
	fmt.Println(val1)
	var val2 = mp["Pham"]
	fmt.Println(val2)

	//duyệt
	for k, v := range mp {
		fmt.Println(k, v)
	}

	// kiểm tra có thuộc map ko
	_, prs := mp["Long"]
	if prs {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

	//delete
	delete(mp, "Long")
	fmt.Println(mp)

	//Set  = map[T]bool
	mySet := make(map[int]bool)
	mySet[1] = true
	mySet[2] = true
	mySet[2] = true
	mySet[1] = true

	for k, v := range mySet {
		fmt.Println(k, v, " ")
	}

}
