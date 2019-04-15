package main

import "fmt"

/**
Go语言提供了数组类型和数据结构
数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型，例如整形、字符串或者自定义类型

声明数组：
go语言数组声明需要指定元素类型和元素个数，语法格式：
var variable_name [SIZE] variable_type  一维数组定义方式
eq: var balance [10] float32
初始化数组：
var balance =[5]float32{1000.0,2.0,3.4,7.0,50.0}
初始化数组中的｛｝中的个数不能大于[]中的数字
如果忽略[]中的数字不设置数组大小，Go语言回根据元素的个数来这只数组的大小
var balance  =[...]float32{1000,0,2.0,3.4,7.0,50.0}
多维数组：
声明方式： var variable_name [SIZE1][SIZE@]...[SIZEN] variable_type
var threedim [5][10][4] int
向函数传递数组：
*/
func main() {
	var n [10]int
	var i, j int
	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d]= %d\n", j, n[j])
	}
	//二维数组
	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	var k, m int
	for k = 0; k < 5; k++ {
		for m = 0; m < 2; m++ {
			fmt.Printf("a[%d][%d] = %d\n", k, m, a[k][m])
		}
	}

	var balance = []int{1000, 2, 3, 17, 50}
	var avg float32
	avg = getAverage(balance, 5)
	fmt.Printf("平均值： %f", avg)
}

//向函数传递数组
func getAverage(arr []int, size int) float32 {
	var i, sum int
	var avg float32
	for i = 0; i < size; i++ {
		sum += arr[i]
	}
	avg = float32(sum) / float32(size)
	return avg
}
