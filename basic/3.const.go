
package main

import "fmt"
import "unsafe"

// 常量是一个简单值的标识符，在程序运行时，不会被修改的量
// 常量的数据类型只能是布尔型、数字型和字符串，不能是派生类型
// 因此，常量往往省略类型，编译器会自动推断其所属类型

// 常量还可以用作枚举
const (
	unknown = 0
	female  = 1
	male    = 2
)

// 常量还可以使用函数来计算，必须为内置函数。常量在编译期间值必须确定下来。
const (
	a1="abc"
	b1=len(a1)
	c1=unsafe.Sizeof(a1)
)

// iota 是一个特殊的常量，会被编译器自动优化为数值
// iota 表示的在 const 语句块中的行索引
// iota 可以被用作枚举值

// const 中未赋值的变量 h2 ，将自动赋值为上一个变量g2
const(
	//a0
	a2=iota //iota=0
	b2=iota //iota=1
	c2      // 常量没有被初始化 则自动初始化为上一个常量的值 c2=2 当然第一常量不允许这样做
	d2,e2,f2=iota,iota,iota // iota=3
	g2=100
	h2  //没有错误 idea 显示有问题

)


// <<iota  <<表示左移的意思 即二进制左移iota位 相当于 *2*iota
const(
	i=1<<iota
	j=3<<iota
	k
	l
)

func constTest(){

	fmt.Println(a1,b1,c1)
	fmt.Println(a2,b2,c2,d2,e2,f2,g2,h2)
	fmt.Println(i,j,k,l)

}

func main() {
	// 常量
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str"
	area = LENGTH * WIDTH

	// Printf 格式化输出
	fmt.Printf("The area is: %d \n", area)


	println(a, b, c)
	println("-------------")

	//println()
	//fmt.Println()
	constTest()
}
