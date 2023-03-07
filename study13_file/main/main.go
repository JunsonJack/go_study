package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// 打开文件，只读
	file,err := os.Open("./main.go")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 读取文件里面的内容
	var strSlice []byte
	var tempSlice = make([]byte,128) //每次读取128个字节
	for {
		n,err := file.Read(tempSlice)
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		// fmt.Printf("读取到了%v个字节",n)
		strSlice = append(strSlice, tempSlice[:n]...)
	}
	
	
	fmt.Println(string(tempSlice))
}