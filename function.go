package main

import (
	"fmt"
	"math"
)

/**
Go语言至少有个main（）函数
函数声明告诉了编译器函数的名称，返回类型和参数
Go语言标准库提供了多种可动用的内置的函数，例如：len()
函数的定义：
func functiong_name ([parameter list]) [return_types] {
	函数体
}
parameter list :参数列表，参数就像一个占位符，当函数被调用时，可以将值传递给参数，这个值被称为实际参数，参数列表指定的是参数类型、顺序及参数个数，参数是可选的，也就是说函数也可以不包含参数
return_types ：返回类型，函数返回一列值，return types 是该列值的数据类型。有些功能不需要返回值，这种情况下return_tyoes不是必须的
函数参数：
函数如果使用参数，该变量可称为函数的形参
形参就像定义在函数体内的局部变量
调用函数，可以通过两种方式来传递参数
值传递：默认情况下，Go语言使用的是值传递，即在调用过程中不影响到实际参数。--》值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数
引用传递：引用传递是指在调用函数式将实际参数的地址传递到函数中，那么在函数中对参数进行的修改，将影响到实际参数
函数用法：
函数作为值：函数定义后可作为值来使用

闭包：闭包是匿名函数，可在动态编程中使用
Go语言支持匿名函数，可作为闭包，匿名函数是一个“内联”语句表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必声明
方法：方法就是包含了接受者的函数
Go语言中同时有函数和方法。一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。所有给定类型的方法属于该类型的方法集。语法格式如下：
func (variable_name variable_data_type) function_name() [return_type] {
	函数体
}
*/
/* 定义结构体*/
type Circle struct {
	radius float64
}

func main() {
	var a int = 100
	var b int = 200
	var ret int
	ret = max(a, b) //调用函数返回最大值
	fmt.Printf("最大值是：%d\n", ret)
	//多个返回值调用函数
	a1, b1 := swap("Mahesh", "Kumar")
	fmt.Println(a1, b1)
	//函数作为值使用
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	//使用函数
	fmt.Println(getSquareRoot(9))
	//闭包的使用
	nextNumber := getSequence()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	//创建新的函数
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
	//方法
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积= ", c1.getArea())
}

/**
函数返回两个数的最大 值
*/
func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

/**
多个返回值
*/
func swap(x, y string) (string, string) {
	return y, x
}

/**
匿名函数 ，闭包的使用
*/
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

//方法  改method 属于类型对象中的方法
func (c Circle) getArea() float64 {
	//c,radius 即为Circle类型对象中的属性
	return 3.14 * c.radius * c.radius
}
