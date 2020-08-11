package _2C_条件语句for

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var a = "abcdefg"
	for i, n := 0, len(a); i < n; i++ {
		fmt.Println(i, a[i])
	}




}
