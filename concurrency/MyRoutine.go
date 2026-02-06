package concurrency

import (
	"fmt"
)

func sayHi() {
	fmt.Println("Hi")
}

func MyRoutine() {
	go sayHi()
	fmt.Println("Long")

	//output:
	// Long
	//Không in ra Hi vì goroutine sayHi có thể chưa được go lập lịch để chạy
	//trong khi MyRoutine đã kết thúc làm process kết thúc theo.
}
