
package main

import "fmt"

// 多变量声明，一般用于声明全局变量
// 全局变量允许声明但是不使用
var (
	v1,v2 int
	v3 string
)

var v4,v5 int


func main(){

	// 声明的局部变量必须使用，不然会编译错误。
	// 声明一个变量并且初始化。

	var a string="Hello"
	fmt.Println(a)

	var b,c int=1,2
	fmt.Println(b,c)

	// 没有初始化的变量会自动赋值为系统默认的零值。
	// 数值类型赋值为 0  布尔类型赋值为 false 字符串类型赋值为""  派生类型为nil
	var d int
	fmt.Println(d)

	var e bool
	fmt.Println(e)

	// 根据值自动判断变量类型, 变量必须声明类型或者赋值（Python 不需要)
	var f=true
	fmt.Println(f)

	// 省略 var 需要添加 :=
	// := 是一个声明语句, 左侧必须是新的变量
	g:=10
	fmt.Println(g)

	// 和 var h string = "Ahri" 相同
	// 但是这种方式只能用在局部变量当中，而不能用于全局变量的声明和赋值
	h:="Ahri"
	fmt.Println(h)

	fmt.Println(v1,v2,v3,v4,v5)

	// 交换 b c 的值, 但是类型必须是相同的
	b,c=c,b
	// 和 Python 一样，可以使用 _ 作为占位符，占位符中的值将被抛弃
	a1,a2,_,a4:=1,2,3,4

	fmt.Println(a1,a2,a4)

	/*
	我们可以看到Go 的内存模型和 Python 的明显不同
	Go 是不区分可变类型和不可变类型的
	对于基本类型的赋值 都是重新申请一块内存 不会有不同变量指向同一基本类型内存地址

	Go 代码与 Python代码对比 可以看到明显的不同
	-- Python --
	i=10
	k=i
	print(id(i),i)
	print(id(k),k)
	i=20
	print(id(i),i)
	print(id(k),k)

	 */
	var i int=10
	var k =i
	fmt.Println(&i,i)
	fmt.Println(&k,k)
	i=20
	fmt.Println(&i,i)
	fmt.Println(&k,k)
}