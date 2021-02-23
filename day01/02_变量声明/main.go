package main

import "fmt"

//函数外部声明并初始化变量，一次初始化多个变量，并可以将变量的类型省略
var mobile, score = "17521242662", 90

func main()  {

	//变量标准声明 --->>  var 变量名 变量类型

	//var name string
	//var age int
	//var isStudent bool

	//变量批量声明

	//var  (
	//	name string
	//	age int
	//	isStudent bool
	//)

	//变量的初始化 --->>  var 变量名 类型 = 表达式


	//函数内部声明并初始化变量，后面都用这种简略方式
	name := "xiaohaiteng"
	age := 25

	fmt.Println(name, age, mobile, score)
	test()
}


func test()  {
	sclool := "华东师范大学"
	name := "张佳"
	score := 100

	fmt.Println(sclool, name, score)
}