package struct_embedding

import "fmt"

// golang ko sử dụng extend như java
//--> dùng struct embedding để tái sử dụng code + mô phỏng kế thừa!

type Person struct {
	name string
	age  int
}

func (person Person) SayHi() {
	fmt.Println("Hi")
}

type Student struct {
	Person
	studentCode string
}

type Teacher struct {
	Person
	teacherCode string
}

func (teacher Teacher) SayHi() {
	fmt.Println("Hi Teacher")
}

func MyStructEmbedding() {

	student := Student{
		Person: Person{
			"Long",
			12,
		},
		studentCode: "B22PTIT123",
	}

	teacher := Teacher{
		Person: Person{
			"Hoang",
			30,
		},
		teacherCode: "GVPTIT123",
	}

	fmt.Println(student, teacher)

	student.SayHi()
	teacher.SayHi()

	//output:
	//{{Long 12} B22PTIT123} {{Hoang 30} GVPTIT123}
	//Hi
	//Hi Teacher

}
