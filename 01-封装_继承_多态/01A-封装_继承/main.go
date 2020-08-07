package _1A_封装_继承

import "fmt"

// 1. 封装：指的是字段名，方法名的大小写，决定其是公有还是私有

// 2. 继承
type Person struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

type Student struct {
	Person // 继承Person

	Class  string `json:"class"`
	Number string `json:"number"`
}

func (s *Student) Introduce() {
	fmt.Printf("Hi,my name is %s,i am %s years old,i am class %s,my number is %s", s.Name, s.Age, s.Class, s.Number)
}

func main() {

	lucy := Student{Person{Name: "Lucy", Age: "12"}, "five", "15"}
	lucy.Introduce()

}


