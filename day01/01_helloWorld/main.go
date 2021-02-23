package main

import "fmt"

func main()  {
	fmt.Println("hello world! 我是小海腾！！！！")
	new()
	fmt.Println(add(3, 8))
}

func new()  {
	fmt.Println("这是new方法")
}

func add(a, b int) int {
	return a + b
}
