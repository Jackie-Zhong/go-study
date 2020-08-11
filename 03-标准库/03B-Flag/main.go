package _3B_Flag

import (
	"flag"
	"fmt"
	"time"
)

// 知道咋回事就行了

func main() {

	//flag.Type()

	fmt.Println(flag.Arg(0))                                 // 无输出
	fmt.Println(flag.String("1", "1", "1"))                  // 0xc000088220
	fmt.Println(flag.Int("2", 1, "1"))                       // 0xc0000a00b0
	fmt.Println(flag.Bool("3", true, "1"))                   // 0xc0000a00b8
	fmt.Println(flag.Float64("4", 0.1, "1"))                 // 0xc0000a00c0
	fmt.Println(flag.Duration("5", 2*time.Millisecond, "1")) // 2ms

	// flag.TypeVar()

	var name string
	flag.StringVar(&name, "name", "zhangsan", "姓名")
	fmt.Println(name)

}
