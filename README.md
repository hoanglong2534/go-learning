# Golang Learning

## 1. Hello World

```go
func HelloGolang() {
    fmt.Println("Hello GoLang")
}
```

## 2. Values, Variables & Constants

```go
fmt.Println("Day la" + " " + "string!") // Values
a := 3                                  // Variables (Type Inference)
const n = 500000000                     // Constant (Bất biến)
```

## 3. For Loop
Go chỉ có duy nhất vòng lặp `for` (không có while, do-while).
```go
for i := 1; i <= 10; i++ {
    fmt.Print(i, " ")
}
```

## 4. Array & Matrix
Mảng có độ dài cố định. Matrix là mảng 2 chiều.
```go
// Array 1D
var a [10]int

// Matrix (Array 2D)
var arr2D [10][10]int
arr2D[0][0] = 1
```

## 5. Slice
Mảng động, linh hoạt hơn Array, dùng `append` để thêm phần tử. 
```go
s := []int{}
s = append(s, 1, 2, 3)
```

## 6. Map & Set
Lưu trữ dạng Key-Value. Go không có Set sẵn, thường dùng `map[T]bool` để giả lập.
```go
// Map
mp := make(map[string]int)
mp["Long"] = 1

// Set (Giả lập)
mySet := make(map[int]bool)
mySet[1] = true
```

## 7. Function
Hàm trong Go có thể trả về nhiều giá trị cùng lúc (thường là kết quả + lỗi).
```go
func Tong(a, b int) (int, string) {
    return a + b, "OK"
}
```

## 8. Struct

```go
type User struct {
    Name string
    Age  int
}
```

## 9. Method
Hàm gắn liền với Struct. Dùng con trỏ `*User` nếu muốn sửa dữ liệu gốc.
```go
func (u *User) TangTuoi() {
    u.Age++
}
```

## 10. Interface
Định nghĩa hành vi, không cần từ khóa `implements` 
```go
type HinhHoc interface {
    DienTich() float64
}
```

## 11. Enum
Go không có Enum, dùng `type` mới + `const` để giả lập.
```go
type Status string
const ON Status = "ON"
```

## 12. Struct Embedding (Kế thừa)
Nhúng struct này vào struct khác để tái sử dụng code (Composition over Inheritance).
```go
type Person struct { Name string }
type Student struct {
    Person // Embedding: Student có luôn field Name của Person
    Code string
}
```

## 13. Error Handling
Go không dùng try-catch. Lỗi được trả về như một giá trị bình thường.
Hàm trả về gồm kết quả và error nếu có lỗi 
```go
if err != nil { fmt.Println(err) }
```

## 14. Concurrency

### Goroutine & Channel
Goroutine là luồng siêu nhẹ. Channel dùng để giao tiếp/trao đổi dữ liệu giữa các luồng.
```go
go func() { 
	ch <- "Done" 
}() 
msg := <-ch                 
```

### WaitGroup
Dùng để chờ một nhóm Goroutine hoàn thành công việc.
```go
wg.Add(1)
go func() { defer wg.Done(); ... }()
wg.Wait() // Chặn Main lại chờ
```

### Mutex (Race Condition)
Dùng khóa (`Lock`) để bảo vệ dữ liệu chung khi nhiều luồng cùng sửa, tránh sai lệch số liệu.
```go
mu.Lock()   // Khóa cửa
money++     // Sửa an toàn
mu.Unlock() // Mở cửa
```
