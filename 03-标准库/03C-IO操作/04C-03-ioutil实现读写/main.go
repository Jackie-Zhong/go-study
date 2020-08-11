package _4C_03_ioutil实现读写

import (
	"fmt"
	"io/ioutil"
)

// 使用ioutil工具读取文件，简单粗暴好用

// 写入文件
func WriteFile() {

	err := ioutil.WriteFile("melody.txt", []byte("aaa"), 0066) // 写入文件 aaa
	if err != nil {
		fmt.Println(err)
		return
	}

}

// 读取文件
func ReadFile() {
	content, err := ioutil.ReadFile("./melody.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}

func main() {
	WriteFile()
	ReadFile()

}
