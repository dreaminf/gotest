package main

import (
	"fmt"
	"unsafe"
)

/**
Go语言变量名由字母、数字、下划线组成，其中首个字符不能为数字
全局变量允许声明但不使用
函数体内变量一旦声明，必须在代码块中使用，否则编译报错
字符串不能用单引号，需要使用双引号
*/
var a = "abcddd" //根据值自行判定变量类型
var b string = "ruboob.com"
var c bool
var d int = 35
var ( //这种因式分解关键字的写法一般用于声明全局变量
	x int
	y int
)

/*
多变量声明
*/
var vname1, vname2, vname3 string
var d2, d3, d4 = "23", "34", 3 //多变量声明赋值

/**
常量
常量中的数据类型只可以是布尔型，数字型（整数，浮点型和复数）和字符串型
定义格式为：const identifier [type] = value
显式类型定义 const b string = "abc"
隐式类型定义 const b = "abc"
多个相同类型的声明可简写：
const c_name1, c_name2 = value1, value2
常量可以用len(),cap(),unsafe.Sizeof()函数计算表达式的值，常量表达式中，函数必须是内置函数，否则编译错误
*/
//常量 用作枚举
const (
	Unknown = 0
	Female  = 1
	Male    = 2
)
const (
	aa = "abc"
	bb = len(aa)
	cc = unsafe.Sizeof(aa) //字符串类型在go里是个结构，包含指向底层数组的指针和长度，这两部分每部分都是8个字节，所以字符串类型大小为16个字节
)

/**
iota 特殊常量，可以认为是一个被编译器修改的常量
iota在const关键字出现时将被重置为0（const内部的第一行之前），
const中没新增一行常量声明将使iota计数一次（iota可以理解为const语句块中的行索引）
*/
const (
	aaa = iota
	bbb = iota
	ccc = iota
)
const (
	aaa1 = iota
	bbb1 // 在定义常量组时，如果不提供初始值，则表示使用上行的表达式
	ccc1
)
const (
	aaa2 = iota
	bbb2
	ccc2
	ddd2 = "ha"
	eee2
	fff2 = 100
	ggg2
	hhh2 = iota
	iii2
)
const (
	i = 1 << iota // << 表示左移  左移0位，不变仍为1；
	j = 3 << iota // 左移1位，变为二进制110，即6
	k             //左移2位 变为二进制1100，即12
	l             //左移3位，变为二进制11000，即24
)

func main() {
	var age int = 2 //定义变量 定义在函数中变量的必须使用
	/* 我的第一个简单程序 */
	ef := "string"                               //此种变量声明方式,不能放在函数外
	vname1, vname2, vname3 = "abc", "bcd", "def" //不能在函数外赋值
	fmt.Println("hello,world", a, b, c, age)
	fmt.Println(vname1, vname2, vname3, ef)
	const LENGTH int = 10
	const WIDTH int = 5
	const a1, b1, c1 = 1, false, "str"
	var area int
	area = LENGTH * WIDTH
	fmt.Printf("area is : %d", area)
	println()
	println(a1, b1, c1)
	println(aa, bb, cc)
	println(aaa, bbb, ccc)
	println(aaa1, bbb1, ccc1)
	println(aaa2, bbb2, ccc2, ddd2, eee2, fff2, ggg2, hhh2, iii2)
	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)

}
