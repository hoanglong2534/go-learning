package concurrency

import (
	"fmt"
)

func sayHi(channel chan string) {
	fmt.Println("Hi")

	channel <- "sayHi da chay xong!"
}

func MyRoutine() {

	myChannel := make(chan string, 2)

	myChannel <- "ptit"

	go sayHi(myChannel)

	fmt.Println("Long")

	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)

	//output:
	// Long
	//ptit
	//Hi
	//sayHi da chay xong!

	// myChannel có size = 2, ban đầu đẩy "ptit" vào queue , sau đó chạy sayHi ở 1 go routine khác, tiếp tục
	// chạy in ra "Long" ,
	// in ra "ptit"
	// sau đó hàm sayHi chạy xong in ra Hi, và đẩy tiếp "sayHi da chay xong!" vào queue
	// in từ queuer	ra "sayHi da chay xong!"

	// tùy CPU lập lịch mà "ptit" và "Hi" có thể in ra trước hoặc sau nhau
}
