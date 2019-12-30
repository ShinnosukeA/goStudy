package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readFile() {
	//打开文件
	file, e := os.Open("./open.go")
	if e != nil {
		fmt.Printf("there is a error %v\n", e)
		return
	}
	//关闭文件 使用defer来释放资源
	defer file.Close()
	//读文件
	tmp := [128]byte{}
	for {
		n, err := file.Read(tmp[:])
		if err != nil {
			fmt.Printf("读取文件错误 %v\n", err)
			return
		}
		fmt.Println(n)

		fmt.Println(string(tmp[:n]))
		if n < 128 {
			return
		}
	}
}

func readFileByBufio() {
	file, err := os.Open("./open.go")
	if err != nil {
		fmt.Printf("打开文件出错了 %v\n", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		s, err := reader.ReadString('1')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("读取文件出错了 %v \n", err)
			return
		}
		fmt.Println(s)
	}

}
func main() {
	//readFile()
	readFileByBufio()
}
