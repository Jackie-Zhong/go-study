package _2A_条件语句switch

import "fmt"

func main() {

	var mark int
	fmt.Print("请输入您的分数：")
	fmt.Scan(&mark)

	switch {
	case mark > 90:
		fmt.Println("a")
	case mark > 80 && mark < 90:
		fmt.Println("b")
		//fallthrough             // 不跳出循环，继续执行。 输出：b c
		//break
	default:
		fmt.Println("c")
	}

	/*----------------------------------------------------------------------*/

	var grade string
	fmt.Print("请输入您的等级：")
	fmt.Scan(&grade)

	switch {
	case grade == "A":
		fmt.Println("优秀")
	case grade == "B":
		fmt.Println("良好")
	case grade == "C":
		fmt.Println("及格")
	default:
		fmt.Println("差")
	}

	/*----------------------------------------------------------------------*/

	var x interface{} // 不能写成 var x string 或者 var x int 等等

	switch i := x.(type) {
	case nil:
		fmt.Println(i)
	case int:
		fmt.Println(i)
	case float64:
		fmt.Println(i)
	case bool:
		fmt.Println(i)
	case string:
		fmt.Println(i)
	case func(int) float64:
		fmt.Println(i)
	default:
		fmt.Println("未知型")
	}

	/*----------------------------------------------------------------------*/

	var m = 0

	switch m {
	case 0:
		fmt.Println("0") // 如果这里被注释掉，则不输出信息
		fallthrough      // 这里fallthrough执行之后，输出 0，1。后面的不会输出
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("def")
	}

}
