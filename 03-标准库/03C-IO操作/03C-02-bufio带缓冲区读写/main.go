package _3C_02_bufio带缓冲区读写

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 写入文件
// 创建文件  ->  使用bufio写入文件  -> 刷新缓冲区    ->  延迟关闭
func WriteFile() {

	// 创建文件
	file, err := os.Create("sweet.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	// 使用bufio写入文件
	BufFile := bufio.NewWriter(file)
	BufFile.WriteString("this is a sweet candy\n")

	for i := 0; i < 5; i++ {
		BufFile.WriteString("this is a sweet story\n")
	}

	// 刷新缓冲区
	BufFile.Flush()
}

// 读取文件
// 创建文件  ->  使用bufio读取文件  -> 一行一行读取    ->  延迟关闭
func ReadFile() {

	// 打开文件
	file, err := os.Open("./sweet.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	// 读取文件
	readFile := bufio.NewReader(file)

	// 一行一行读取
	for {
		line, _, err := readFile.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
		fmt.Println(string(line))
	}
}

func main() {
	WriteFile()
	ReadFile()
}
