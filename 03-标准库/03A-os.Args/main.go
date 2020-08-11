package _3A_os_Args

import (
	"fmt"
	"os"
	"strconv"
)

// 这个命令知道干啥的就行了

// os.Args 简单的获取命令行参数
// os.Args []string类型

func main() {

	// go run main.go 1 3 -X
	for i, args := range os.Args {
		fmt.Println("参数"+strconv.Itoa(i)+":", args, len(os.Args))
	}

	// go run main.go 1 2 3
	fmt.Println(os.Args) // 输出： 1 2 3
}
