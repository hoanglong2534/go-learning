package methods

import "fmt"

type SinhVien struct {
	ten   string
	tuoi  int
	email string

	// struct trong golang chỉ có data , ko có function như java , c..
	// ví dụ ko có kiểu này trong golang
	//func study void
	//	fmt.Println("study")
}

// thay vào đó golang có method và receiver
func (sv SinhVien) say() {
	fmt.Println("hello ")
}

func (sv SinhVien) tangTuoi(tang int) {
	sv.tuoi += tang
}

func (sv *SinhVien) tangTuoi2(tang int) {
	sv.tuoi += tang
}

func MyMethods() {
	s := SinhVien{"Long", 20, "long@gmail.com"}
	s.say()

	s.tangTuoi(2)
	fmt.Println(s.tuoi)

	// pointer
	s.tangTuoi2(2)
	fmt.Println(s.tuoi)

	//output:

	// hello
	// 20 --> ko tăng thành 22 vì golang sẽ tạo sv là 1 bản sao của s rồi sửa ở đó
	// 22 --> đúng, sửa trực tiếp trên s

}
