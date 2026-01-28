package errors

import (
	"errors"
	"fmt"
)

// golang ko có try-catch-finally  như java
// --> trả error như 1 giá trị bình thường

func chia(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("ko dc chia cho 0!")
	}
	return a / b, nil
}

func MyError() {

	val, err := chia(10, 0)
	if err != nil {
		fmt.Println("Lỗi: ", err)
	} else {
		fmt.Println(val)
	}

	// output:
	//Lỗi:  ko dc chia cho 0!

}
