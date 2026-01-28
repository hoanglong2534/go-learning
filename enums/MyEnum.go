package enums

import "fmt"

func updateStatus(status StatusEnum) StatusEnum {
	if status == PENDING {
		status = OFF
	} else if status == OFF {
		status = ON
	}
	return status
}

func MyEnums() {
	res := updateStatus(PENDING)
	fmt.Println(res)
	res = updateStatus(OFF)
	fmt.Println(res)

	// Output:
	// OFF
	// ON
}
