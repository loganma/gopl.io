// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

// Echo1 prints its command-line arguments.
// package 需要加注释， main 注释包含一句或者几句话
package main

// 列表形式用的多
import (
	"fmt"
	"os"
)

func main() {
	// 声明变量会初始化，这里隐式赋值空字符 ""
	var s, sep string
	// := 短变量声明 short variable declaration
	// 根据初始值为这些变量赋予适当类型的声明语句
	// go 只有 for 循环一种循环语句
	// for initialization; condition; post {
	// }
	for i := 1; i < len(os.Args); i++ {
		// += 是赋值运算符 assignment operator
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

//!-
