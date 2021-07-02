package main

import (
	"fmt"
	"os"
)

func main() {
	fp, err := os.Create("./a.txt")
	if err != nil {
		fmt.Println("文件创建失败")
		return
	}
	defer fp.Close()
	fmt.Println("文件创建成功")
}
