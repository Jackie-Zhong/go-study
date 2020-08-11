package _3C_01_简单写读拷贝文件

import (
	"fmt"
	"io"
	"os"
)

// 简单掌握写入文件，读取文件，拷贝文件

// 写入文件
// 创建文件  ->  写入文件  -> 延迟关闭文件
func WriteFile() {

	// 创建文件
	file, err := os.Create("zzh.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	// 写入
	file.WriteString("i am a chinese\n")
	// 写入
	file.Write([]byte("i am a boy\n"))
	// 循环写入
	for i := 0; i < 5; i++ {
		file.WriteString("i am a student\n")
	}

}

// 读取文件
// 打开文件  ->   读取文件   ->  延迟关闭文件
func ReadFile() {

	// 打开文件
	file, err := os.Open("./zzh.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// 读取文件
	//var buf [128]byte
	//var content []byte
	buf := make([]byte, 1024)     // 有缓冲读取
	content := make([]byte, 1024) // 有缓冲读取
	for {
		bufFile, err := file.Read(buf[:])
		if err == io.EOF { //读取结束
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		content = append(content, buf[:bufFile]...)
	}
	fmt.Println(string(content))
}

// 拷贝文件
// 打开源文件 ->  新建目标文件 -> 拷贝源文件  -> 写入到目标文件  -> 延迟关闭文件
func CopyFile() {

	// 打开文件
	srcFile, err := os.Open("./zzh.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer srcFile.Close()

	// 创建新文件
	dstFile, err := os.Create("zhong.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer dstFile.Close()

	// 读取文件
	//var buf [128]byte
	buf := make([]byte, 1024) // 有缓冲读取
	for {
		bufSrcFile, err := srcFile.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		dstFile.Write(buf[:bufSrcFile])
	}

}

func main() {
	WriteFile()
	ReadFile()
	CopyFile()
}
