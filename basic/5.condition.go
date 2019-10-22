package main

import (
	"fmt"

)

func main(){
	var a int =20
	// if 条件语句 condition 不需要写括号
	// else 也可以不带
	if a>10 {
		fmt.Println("Hello!")
	}else{
		fmt.Println("World!")
	}

	// switch 条件语句
	// switch 语句,从上到下进行匹配,不需要添加 break
	var marks int=90
	var grade string
	switch {
	case marks<100 && marks>=90:

		grade="A"
		fmt.Println("A")

	case marks<90:
		grade="B"
		fmt.Println("B")

	}

	// switch 的另一种形式
	switch grade{
	case "A":
		fmt.Println("Excellent")

		// fallthrough 可以强制执行后面的 case 语句，类似C 语言的 switch
		fallthrough
	case "B":
		fmt.Println("Good")
	default:
		fmt.Println("Bad")
	}

	// interface 是 go 的派生类型之一 零值为 nil
	// @Star  interface 还有 nil 需要更加详细的说明
	var x interface{}

	print()
	switch x.(type){
	case nil:
		fmt.Printf("x 的类型是: %T\n",x)
		// fallthrough  cannot fallthrough in type switch
	case int:
		fmt.Printf("int\n")
	case float64:
		fmt.Printf("float64\n")
	default:
		fmt.Printf("Unknow\n")
	}

	var c1,c2,c3 chan int
	var i1,i2 int

	// select 语句
	// select 的每个 case 必须是一个通信操作，要么是发送要么是接收
	// select 随机执行一个可运行的case，如果没有case可执行，select 将阻塞，直到有case可运行
	// @Star select 语句与 chan 类型需要了解


	select{
	case i1=<-c1:
		fmt.Printf("Received ", i1," from c1\n")
	case c2<-i2:
		fmt.Print("Sent",i2,"to c2 \n")
	case i3,ok:=<-c3:
		if ok{
			fmt.Printf("Received",i3,"from c3\n")
		}else{
			fmt.Printf("c3 is closed\n")
		}

	// 如果有 default 那么 select 总是可以执行的
	default:
		fmt.Printf("No communication\n")

	}

	// for 循环
	// Go 的循环支持 break continue 语句
	// numbers 是一个有6个元素的数组
	// 前四个为 1,2,3,5 后两个为0值
	numbers:=[6]int{1,2,3,5}

	// 类似 java 的 for(,,) 的语法
	// for 的条件变量属于 循环的局部变量  不会和func的变量发生冲突
	for a:=0;a<10;a++{
		fmt.Printf("a 的值为 %d\n",a)

	}

	var a0 int
	var b0 int = 3

	// 类似于 java 的 while(){} 的语法
	for a0<b0{
		a0++
		fmt.Printf("a0 的值为 %d\n",a0)
	}
	// 死循环
	//for true {
	//
	//}

	// 类似于 Python 的 dict.items() 返回key,value 的用法
	// for .. range .. 的形式类似于 Python 的 for .. in .. 可以对 slice map 数组 字符串 进行循环迭代
	for i,x:=range numbers{
		fmt.Println(i,x)

	}

	// goto
	// 无条件的跳转到指定行
	var tag int = 10
customLabel:
	for tag < 20 {
		tag++
		if tag > 15 {
			goto customLabel
		}
		fmt.Printf("tag 的值为: %d\n", tag)
	}

}