package main

import (
	"fmt"
	"runtime"
)

func main() {
	numCPU := runtime.NumCPU()
	fmt.Println("当前系统 CPU 核心数量:", numCPU)
	runtime.GOMAXPROCS(numCPU)
	fmt.Println("设置 GOMAXPROCS 完成")
}
