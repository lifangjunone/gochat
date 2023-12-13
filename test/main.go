package main

import (
	"fmt"
	"os"
	"sync"
)

// 需要执行一次的函数
func myFunction() {
	fmt.Println("这个函数只会执行一次")
}

func main() {
	var once sync.Once
	fmt.Println(os.Getenv("RUN_MODE"))
	// 需要执行一次的函数通过 once.Do() 运行
	once.Do(myFunction) // 这里传入函数名，注意不要加括号，直接传函数名
	fileName, _ := os.Getwd()
	fmt.Println(fileName)
	// 可以再次调用，但函数内部不会再次执行
	once.Do(myFunction)
}
