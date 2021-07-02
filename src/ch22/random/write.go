package main

import (
	"fmt"
	"os"
)

// func main() {
// 	fp, err := os.Create("./a.txt")
// 	if err != nil {
// 		fmt.Println("文件创建成功")
// 		return
// 	}
// 	defer fp.Close()
// 	n, _ := fp.WriteString("美女\n")
// 	fmt.Println(n)
// 	fp.WriteString("来玩啊")
// }

func main() {
	//创建文件
	fp, err := os.Create("./a.txt")
	if err != nil {
		fmt.Println("文件创建失败")
		return
	}
	//关闭文件
	defer fp.Close()
	//1、将字符切片写入文件中
	a := []byte{'h', 'e', 'l', 'l', 'o'}
	//
	////文件对象.write(字符切片)
	n, _ := fp.Write(a)
	fmt.Println(n)

	//2、将字符串转成字符切片写入文件中
	//str:="helloworld"
	str := "锄禾日当午"
	//字符串和字符切片允许转换
	b := []byte(str)

	fp.Write(b)
}
