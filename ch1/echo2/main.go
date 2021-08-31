// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	// _ blank identifier 空标识符
	for _, arg := range os.Args[1:] {
		// 如果数据量大，代价很高，使用 join 函数替代
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

//!-
