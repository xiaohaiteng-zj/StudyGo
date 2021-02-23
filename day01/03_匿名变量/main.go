package main

import "fmt"

func foo() (int, string) {
	return 100, "xiaohaiteng"
}

func main()  {
	//在使用多重赋值的时候，想要忽略某个值，就可以使用匿名变量，匿名变量用一个_表示
	x, _ := foo()
	_, y := foo()
	fmt.Println(x, y)
}
