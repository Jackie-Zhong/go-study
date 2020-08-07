package _1C_多态

import "fmt"

// 4. 多态：一种事物的多种形态，一定是要定义接口。比如动物是父类，是个大类，而大象，狮子是子类，是小类。

// 定义一个Animal接口
type Animal interface {
	Walk()
}

// 定义一个Elephont结构体
type Elephant struct {
	Weight string `json:"weight"`
	Height string `json:"height"`
}

// 定义一个Elephont方法
func (e *Elephant) Walk() {
	fmt.Printf("i am a elephont,my weight is %s,my height is %s\n", e.Weight, e.Height)
}

// 定义一个Lion结构体
type Lion struct {
	Color  string `json:"yellow"`
	Length string `json:"length"`
}

// 定义一个Lion方法
func (l *Lion) Walk() {
	fmt.Printf("i am a lion,my color is %s,my length is %s\n", l.Color, l.Length)
}

// 多态
func WhoWalk(a Animal) {
	a.Walk()
}

func main() {
	e1 := &Elephant{Weight: "500kg", Height: "100m"}
	l1 := &Lion{Color: "yellow", Length: "20"}

	WhoWalk(e1)
	WhoWalk(l1)

}
