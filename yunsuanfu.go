package main

import "fmt"

/**
Go 内置运算符有：算数运算符，关系运算符，逻辑运算符，位运算符，赋值运算符，其他运算符
算数运算符：
	+ - *  /  %  ++  --
关系运算符：
	==  ！=  >  < >=  <=
逻辑运算符：
	&& || ！
位运算符：
	& | ^  <<  >>
Go 语言支持的位运算符如下表所示。假定 A 为60，B 为13：
运算符	描述	实例
&	按位与运算符"&"是双目运算符。 其功能是参与运算的两数各对应的二进位相与。 	(A & B) 结果为 12, 二进制为 0000 1100
|	按位或运算符"|"是双目运算符。 其功能是参与运算的两数各对应的二进位相或	(A | B) 结果为 61, 二进制为 0011 1101
^	按位异或运算符"^"是双目运算符。 其功能是参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。	(A ^ B) 结果为 49, 二进制为 0011 0001
<<	左移运算符"<<"是双目运算符。左移n位就是乘以2的n次方。 其功能把"<<"左边的运算数的各二进位全部左移若干位，由"<<"右边的数指定移动的位数，高位丢弃，低位补0。 	A << 2 结果为 240 ，二进制为 1111 0000
>>	右移运算符">>"是双目运算符。右移n位就是除以2的n次方。 其功能是把">>"左边的运算数的各二进位全部右移若干位，">>"右边的数指定移动的位数。 	A >> 2 结果为 15 ，二进制为 0000 1111
位运算符参考链接 http://www.runoob.com/go/go-operators.html
赋值运算符：
	运算符	描述	实例
=	简单的赋值运算符，将一个表达式的值赋给一个左值	C = A + B 将 A + B 表达式结果赋值给 C
+=	相加后再赋值	C += A 等于 C = C + A
-=	相减后再赋值	C -= A 等于 C = C - A
*=	相乘后再赋值	C *= A 等于 C = C * A
/=	相除后再赋值	C /= A 等于 C = C / A
%=	求余后再赋值	C %= A 等于 C = C % A
<<=	左移后赋值 	C <<= 2 等于 C = C << 2
>>=	右移后赋值 	C >>= 2 等于 C = C >> 2
&=	按位与后赋值	C &= 2 等于 C = C & 2
^=	按位异或后赋值	C ^= 2 等于 C = C ^ 2
|=	按位或后赋值	C |= 2 等于 C = C | 2

其他运算符：
	&	返回变量存储地址	&a; 将给出变量的实际地址。
*	指针变量。	*a; 是一个指针变量
指针变量 * 和地址值 & 的区别：指针变量保存的是一个地址值，会分配独立的内存来存储一个整型数字。当变量前面有 * 标识时，才等同于 & 的用法，否则会直接输出一个整型数字。

func main() {
   var a int = 4
   var ptr *int
   ptr = &a
   println("a的值为", a);    // 4
   println("*ptr为", *ptr);  // 4
   println("ptr为", ptr);    // 824633794744
}
运算符优先级

有些运算符拥有较高的优先级，二元运算符的运算方向均是从左至右。下表列出了所有运算符以及它们的优先级，由上至下代表优先级由高到低：
优先级	运算符
7 	^ !
6 	* / % << >> & &^
5 	+ - | ^
4 	== != < <= >= >
3 	<-
2 	&&
1 	||

当然，你可以通过使用括号来临时提升某个表达式的整体运算优先级。
*/
func main() {
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("第一行 -c 的值位%d\n", c)
	c = a - b
	fmt.Printf("第二行 -c 的值位%d\n", c)
	c = a * b
	fmt.Printf("第三行 -c 的值位%d\n", c)
	c = a / b
	fmt.Printf("第四行 -c 的值位%d\n", c)
	c = a % b
	fmt.Printf("第五行 -c 的值位%d\n", c)
	a++
	fmt.Printf("第六行 -c 的值位%d\n", a)
	a = 21
	a--
	fmt.Printf("第七行 -c 的值位%d\n", a)
	//关系运算符
	var a1 int = 21
	var b1 int = 10
	if a1 == b1 {
		fmt.Printf("第一行- a 等于 b \n")
	} else {
		fmt.Printf("第一行 -a 不等于 b\n")
	}
	if a1 < b1 {
		fmt.Printf("第二行 -a 小于b \n")
	} else {
		fmt.Printf("第二行 -a 不小于b \n")
	}
	if a1 > b1 {
		fmt.Printf("第三行 -a 大于b\n")
	} else {
		fmt.Printf("第三行 -a 不大于b\n")
	}
	/*lets change value of a and b*/
	a1 = 5
	b1 = 20
	if a1 <= b1 {
		fmt.Printf("第四行 -a 小于等于b \n")
	}
	if b1 >= a1 {
		fmt.Printf("第五行 -b 大于等于a \n")
	}
	//逻辑运算符
	var a2 bool = true
	var b2 bool = false
	if a2 && b2 {
		fmt.Printf("第一行条件为 true \n")
	}
	if a2 || b2 {
		fmt.Printf("第二行条件为 true \n")
	}
	/*修改a2 和 b2的值*/
	a2 = false
	b2 = true
	if a2 && b2 {
		fmt.Printf("第三行条件为 true \n")
	} else {
		fmt.Printf("第三行条件为false\n")
	}
	if !(a2 && b2) {
		fmt.Printf("第四行条件为 true\n")
	}
	//位运算符
	var a3 uint = 60 /* 60= 0011 1100*/
	var b3 uint = 13 /* 60= 0000 1101*/
	var c3 uint = 0
	c3 = a3 & b3 /* 12 = 0000 1100*/
	fmt.Printf("第一行 -c 的值为 %d\n", c3)
	c3 = a3 | b3 /* 61 = 0011 1101*/
	fmt.Printf("第二行 -c 的值为 %d\n", c3)
	c3 = a3 ^ b3 /* 49 = 0011 0001*/
	fmt.Printf("第三行 -c 的值为 %d\n", c3)
	c3 = a3 << 2 /* 240= 1111 0000*/
	fmt.Printf("第四行 -c 的值为 %d\n", c3)
	c3 = a3 >> 2 /* 15= 0000 1111*/
	fmt.Printf("第五行 -c 的值为 %d\n", c3)
	//赋值运算符
	var a4 int = 21
	var c4 int
	c4 = a4
	fmt.Printf("第1行 - = 运算符实例，c值为 = %d\n", c4)
	c4 += a4
	fmt.Printf("第2行 - = 运算符实例，c值为 = %d\n", c4)
	c4 -= a4
	fmt.Printf("第3行 - = 运算符实例，c值为 = %d\n", c4)
	c4 *= a4
	fmt.Printf("第4行 - = 运算符实例，c值为 = %d\n", c4)
	c4 /= a4
	fmt.Printf("第5行 - = 运算符实例，c值为 = %d\n", c4)
	c4 = 200
	fmt.Printf("第6行 - = 运算符实例，c值为 = %d\n", c4)
	c4 <<= 2
	fmt.Printf("第7行 - = 运算符实例，c值为 = %d\n", c4)
	c4 &= 2
	fmt.Printf("第8行 - = 运算符实例，c值为 = %d\n", c4)
	c4 ^= 2
	fmt.Printf("第9行 - = 运算符实例，c值为 = %d\n", c4)
	c |= 2
	fmt.Printf("第10行 - = 运算符实例，c值为 = %d\n", c4)
	//其他运算符
	var a5 int = 4
	var b5 int32
	var c5 float32
	var ptr *int
	/*运算符实例*/
	fmt.Printf("第1行 -a 变量类型为 = %T \n", a5)
	fmt.Printf("第2行 -b 变量类型为 = %T \n", b5)
	fmt.Printf("第3行 -c 变量类型为 = %T \n", c5)

	/* &和 * 运算符实例*/
	ptr = &a5 /* 'ptr' 包含了 ‘a5’变量的地址*/
	fmt.Printf("a5的值为%d \n", a5)
	fmt.Printf("ptr为%d \n", *ptr)

}
