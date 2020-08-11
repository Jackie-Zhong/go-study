package _2B_条件语句select

import "fmt"

// 少用到select，所以，要求学会看懂代码。 select条件语句的case都是io操作。

func main() {

	var a, b int
	var ah, bh, ch chan int

	select {
	case a = <-ah: // 读取;检测有无数据可以读
		fmt.Println(a)
	case bh <- b: // 写入：检测有无数据可以写
		fmt.Println(b)
	case c, ok := (<-ch): // 读取：检测有无数据可以读
		if ok {
			fmt.Println(c)
		} else {
			fmt.Println("c3 is closed\n")
		}
	default:
		fmt.Println("no output\n")

	}

}
