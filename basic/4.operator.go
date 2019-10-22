package main

// Go 支持 ++ --
// && 逻辑 AND
// || 逻辑 OR
// ! 逻辑 NOT
// & 按位与运算
// | 按位或运算
// ^ 按位异或
// << 左移 高位丢弃，低位补0
// >> 右移 低位丢弃，高位补0


// & 返回变量的指针  &a 将给出a地址。
// * 指针变量 	*a 将给出a所保存的地址所指的内容

import "fmt"

func main(){
	var a int=4
	var b int32
	var c float32
	var ptr *int

	fmt.Println(a,b,c,ptr)

	ptr=&a
	fmt.Println(ptr,*ptr,&ptr)

	var ptrptr =&ptr

	fmt.Println(ptrptr,*ptrptr,&ptrptr)
	fmt.Println(a)

}