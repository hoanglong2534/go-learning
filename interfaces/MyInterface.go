package interfaces

import "fmt"

func display(h HinhHoc) {
	fmt.Println("Chu vi: ", h.chuVi())
	fmt.Println("Dien tich: ", h.dienTich())
}

func MyInterface() {

	hcn := HCN{10, 5}
	hv := HV{3}

	display(&hcn)
	display(&hv)
}
