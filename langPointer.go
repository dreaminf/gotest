package main

import "fmt"

/**
语言指针
Go语言中使用指针可以更简单的执行一些任务
变量是一种使用方便的占位符，用于引用计算机内存的地址
Go语言的取地址符&，放到变量使用就会返回相应变量的内存地址

一个指针变量指向了一个值的内存地址
类似于变量和常量，在使用指针前需要声明指针，指针声明格式如下：
var var_name *var_type
var ip *int
var fp *float32
指针使用流程：
定义指针变量
为指针变量赋值
访问指针变量中指向地址的值
空指针：
当一个指针被定义后没有分配到任何变量时，它的的值为nil
nil指针也称为空指针
nil在概念上和其他语言的null，None，nil，NULL一样，都指向代零值或空值
一个指针变量通常缩写为ptr

 //go指针数组

 //go指向指针的指针
指向指针的指针变量声明格式如下：
var ptr **int
 //go向函数传递指针参数
*/
const MAX int = 3

func main() {
	var a int = 10
	fmt.Printf("变量地址为：%x\n", &a)
	var ip *int
	ip = &a
	fmt.Printf("Ip变量存储的指针地址：%x\n", ip)
	fmt.Printf("*ip变量的值;%d\n", *ip)
	//空指针
	var ptr *int
	fmt.Printf("ptr的值为：%x\n", ptr)
	/**
	空指针判断
	if(ptr != nil) ptr不是空指针
	if(ptr == nil) ptr 是空指针
	*/

	//go指针数组
	b := []int{10, 100, 200}
	var i int
	var ptr1 [MAX]*int
	for i = 0; i < MAX; i++ {
		ptr1[i] = &b[i] /*整数地址赋值给指针数组*/
	}
	for i = 0; i < MAX; i++ {
		fmt.Printf("b[%d]=%d\n", i, *ptr1[i])
	}
	//go指向指针的指针
	var c int
	var ptr2 *int
	var pptr2 **int
	c = 3000
	ptr2 = &c
	pptr2 = &ptr2
	fmt.Printf("变量c=%d\n", c)
	fmt.Printf("指针变量 *ptr2 = %d\n", *ptr2)
	fmt.Printf("指向指针的指针变量 **pptr2 = %d\n", **pptr2)
	//go向函数传递指针参数
	var a1 int = 100
	var b1 int = 200
	swap(&a1, &b1)
	fmt.Printf("交换后a1的值：%d. b1的值：%d\n", a1, b1)
}
func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
