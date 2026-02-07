package concurrency

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func sayHello(id int) {
	defer wg.Done()
	fmt.Println("Hello", id)
}

func MyWaitGroup() {
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go sayHello(i)
	}
	wg.Wait()
	fmt.Println("Done")

	// output:
	//Hello 3
	//Hello 5
	//Hello 10
	//Hello 6
	//Hello 7
	//Hello 8
	//Hello 9
	//Hello 1
	//Hello 2
	//Hello 4
	//Done

	//các go routine chạy đồng thời nên kết quả mỗi lần khác nhau về thứ tự in ra

}
