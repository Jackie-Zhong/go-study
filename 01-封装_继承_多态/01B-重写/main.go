package _1B_重写

import "fmt"

// 3. 重写：重新写个一模一样的方法名

type Person struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

type Student struct {
	Person // 继承Person

	Class  string `json:"class"`
	Number string `json:"number"`
}

type Teacher struct {
	Person // 继承Person

	Sex    string
	weight string
}

func (s *Student) Introduce() {
	fmt.Printf("Hi,my name is %s,i am %s years old,i am class %s,my number is %s\n", s.Name, s.Age, s.Class, s.Number)
}

// 重写
func (t *Teacher) Introduce() {
	fmt.Printf("Hi,my name is %s,i am %s years old,i am a %s,my weight is %s\n", t.Name, t.Age, t.Sex, t.weight)
}

func main() {

	lucy := Student{Person{Name: "Lucy", Age: "12"}, "five", "15"}
	lucy.Introduce()

	marry := Teacher{Person{Name: "Marry", Age: "29"}, "girl", "45kg"}
	marry.Introduce()

}
