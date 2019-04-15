package main

import "fmt"

/**
条件语句
语句	描述
if 语句	if 语句 由一个布尔表达式后紧跟一个或多个语句组成。

if...else 语句	if 语句 后可以使用可选的 else 语句, else 语句中的表达式在布尔表达式为 false 时执行。
if 嵌套语句	你可以在 if 或 else if 语句中嵌入一个或多个 if 或 else if 语句。
switch 语句	switch 语句用于基于不同条件执行不同动作。
 switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上直下逐一测试，直到匹配为止。

switch 语句执行的过程从上至下，直到找到匹配项，匹配项后面也不需要再加 break。

switch 默认情况下 case 最后自带 break 语句，匹配成功后就不会执行其他 case，如果我们需要执行后面的 case，可以使用 fallthrough 。
Go 编程语言中 switch 语句的语法如下：

switch var1 {
    case val1:
        ...
    case val2:
        ...
    default:
        ...
}

变量 var1 可以是任何类型，而 val1 和 val2 则可以是同类型的任意值。类型不被局限于常量或整数，但必须是相同的类型；或者最终结果为相同类型的表达式。

您可以同时测试多个可能符合条件的值，使用逗号分割它们，例如：case val1, val2, val3
Type Switch

switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。

Type Switch 语法格式如下：

switch x.(type){
    case type:
       statement(s);
    case type:
       statement(s);
   你可以定义任意个数的case
default:  可选
statement(s)
}
select 语句	select 语句类似于 switch 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。
 select是Go中的一个控制结构，类似于用于通信的switch语句。每个case必须是一个通信操作，要么是发送要么是接收。

select随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。一个默认的子句应该总是可运行的。
select {
    case communication clause  :
       statement(s);
    case communication clause  :
       statement(s);
     你可以定义任意数量的 case
default :  可选
statement(s);
}
以下描述了 select 语句的语法：

    每个case都必须是一个通信
    所有channel表达式都会被求值
    所有被发送的表达式都会被求值
    如果任意某个通信可以进行，它就执行；其他被忽略。
    如果有多个case都可以运行，Select会随机公平地选出一个执行。其他不会执行。
    否则：
        如果有default子句，则执行该语句。
        如果没有default字句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值。


*/

/**
循环语句
Go 语言提供了以下几种类型循环处理语句：
循环类型 	描述
for 循环 	重复执行语句块
//链接https://www.runoob.com/go-test/go-test-for-loop.html
for循环是一个循环控制结构，可以执行指定次数的循环。
语法

Go语言的For循环有3中形式，只有其中的一种使用分号。

和 C 语言的 for 一样：

for init; condition; post { }

和 C 的 while 一样：

for condition { }

和 C 的 for(;;) 一样：

for { }

    init： 一般为赋值表达式，给控制变量赋初值；
    condition： 关系表达式或逻辑表达式，循环控制条件；
    post： 一般为赋值表达式，给控制变量增量或减量。

for语句执行过程如下：

    ①先对表达式1赋初值；
    ②判别赋值表达式 init 是否满足给定条件，若其值为真，满足循环条件，则执行循环体内语句，然后执行 post，进入第二次循环，再判别 condition；否则判断 condition 的值为假，不满足条件，就终止for循环，执行循环体外语句。

for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。格式如下：

for key, value := range oldMap {
    newMap[key] = value
}
循环嵌套 	在 for 循环中嵌套一个或多个 for 循环
以下为 Go 语言嵌套循环的格式：

for [condition |  ( init; condition; increment ) | Range]
{
   for [condition |  ( init; condition; increment ) | Range]
   {
      statement(s);
   }
   statement(s);
}
循环控制语句

循环控制语句可以控制循环体内语句的执行过程。

GO 语言支持以下几种循环控制语句：
控制语句 	描述
break 语句 	经常用于中断当前 for 循环或跳出 switch 语句
continue 语句 	跳过当前循环的剩余语句，然后继续进行下一轮循环。
goto 语句 	将控制转移到被标记的语句。
无限循环

如过循环中条件语句永远不为 false 则会进行无限循环，我们可以通过 for 循环语句中只设置一个条件表达式来执行无限循环：

package main

import "fmt"

func main() {
   for true  {
	   fmt.Printf("这是无限循环。\n");
   }
}
*/

func main() {
	//if 条件语句
	/* 定义局部变量 */
	var a int = 10
	if a < 20 {
		fmt.Printf("a 小于 20 \n")
	}
	fmt.Printf("a的值为：%d \n", a)
	//switch 语句
	/*定义局部变量*/
	var grade string = "B"
	var marks int = 90
	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"

	}
	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好!\n")
	case grade == "D":
		fmt.Printf("及格!\n")
	case grade == "F":
		fmt.Printf("不及格!\n")
	default:
		fmt.Printf("差!\n")
	}
	fmt.Printf("你的等级是%s\n", grade)
	//Type Switch
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf("x 的类型：%T\n", i)
	case int:
		fmt.Printf("x 的int类型")
	case float64:
		fmt.Printf("x 的float64类型：")
	case func(int) float64:
		fmt.Printf("x 的func(int)类型")
	case bool, string:
		fmt.Printf("x 的类型bool或string")
	default:
		fmt.Printf("未知型")
	}
	//fallthrough 使用 fallthrough 会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true
	switch {
	case false:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case false:
		fmt.Println("3、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("4、case 条件语句为 true")
	case false:
		fmt.Println("5、case 条件语句为 false")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}

	//select case
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received", i1, "from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, "to c2 \n")
	case i3, ok := (<-c3): //<- 是什么意思没搞懂？？
		if ok {
			fmt.Printf("received", i3, "from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
	//for 循环
	var b1 int = 15
	var a1 int
	numbers := [6]int{1, 2, 3, 5}
	/* for 循环*/
	for a1 := 0; a1 < 10; a1++ {
		fmt.Printf("a1的值为：%d\n", a1)
	}
	for a1 < b1 {
		a1++
		fmt.Printf("a1的值为：%d\n", a1)
	}

	for i, x := range numbers {
		fmt.Printf("第%d位 x的值 =%d\n", i, x)
	}
	//循环嵌套输出2到100之间的素数
	var ik, j2 int
	for ik = 2; ik < 100; ik++ {
		for j2 = 2; j2 <= (ik / j2); j2++ {
			if ik%j2 == 0 {
				break //如果发现因子，则不是素数
			}
		}
		if j2 > (ik / j2) {
			fmt.Printf("%d 是素数\n", ik)
		}
	}
}
