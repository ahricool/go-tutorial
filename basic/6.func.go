package main

import (
	"fmt"
	"math"
)

// 函数是基本的代码块 每个可运行的Go程序 至少有一个main() 函数


// func 关键字 表示一个函数
// 和 Java 类似 函数名和参数列表一起构成了函数的签名
// 参数列表 包括 参数的类型 参数的个数 参数的顺序
// 最后的一个 int 表示函数的返回类型
// 有些函数不需要返回值 返回类型 不是必须的
// 函数体 函数定义的代码的集合


func max(num1,num2 int) int {
	var res =0

	if (num1>num2){
		res=num1
	}else{
		res=num2
	}
	return res
}

// 函数可以返回多个值
// 返回类型 需要定义成 类似 Python 元组的形式
func swap(x,y string) (string,string){
	return y,x
}

// Go 采用和C 类似的内存模型，所以传入的参数值区分为 值传递 和 引用传递

// 声明 x 和 y 都是 int 的指针类型
// 函数名定义采用驼峰的形式
func swapPoint(x *int,y *int){
	var tmp int

	// 交换地址 a b 所指向的值
	// 这个过程指针是没有任何变化的，不过是指向的值发生了改变
	// 指针传入的时候还是按值传递 你在函数里面修改指针是没有用的 你要修改的是 指针指向的东西
	tmp=*x
	*x=*y
	*y=tmp
}


func main(){
	var a int =100
	var b int =200

	res:=max(a,b)
	fmt.Println("max value is ",res)

	str_a,str_b:=swap("World","Hello")
	fmt.Println(str_a,str_b)

	fmt.Println(a,b)
	fmt.Println(&a,&b)

	// 取 a b 的指针传入
	swapPoint(&a,&b)
	fmt.Println(a,b)
	fmt.Println(&a,&b)

// Go 也有类似于 Python 的lambda 表达式的形式
	lam:=func(x float64) float64{
		return math.Sqrt(x)
	}
	fmt.Println(lam(9))

}