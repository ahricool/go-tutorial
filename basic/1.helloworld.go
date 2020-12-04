
package main
// 类似Java，必须在源文件中非注释的第一行指名这个文件属于哪个包
// package main 表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包

import "fmt"

// 需要用到的包， fmt 包实现了格式化IO 的函数

//	func main() 是程序开始执行的函数，main是可执行程序所必须包含的，是启动后第一个执行的函数
//  如果有 init() 函数会先执行该函数

func main(){
	/*
	完全类似 Java 的注释形式
	 */

	//fmt.Println("Hello,World!")
	//// Println 会自动换行
	//res:= make(map[string]int)
	//res["sad"]=1
	//for k,v:=range res{
	//	fmt.Println(k)
	//	fmt.Println(v)
	//}
	//var key string
	//if (key ==""){
	//	fmt.Println("What???")
	//}
	//
	//
	//var aa float64=111.23
	//fmt.Println(int(aa))
	//
	//notFillinDimension := []string{"123","asd","zcx"}
	//if _,ok:=notFillinDimension.;!ok{
	//	print("hello")
	//}else{
	//	print("word")
	//}

	//wg := sync.WaitGroup{}
	//wg.Add(20)
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		fmt.Println("A: ", i)
	//		wg.Done()
	//	}()
	//}
	//for i := 0; i < 10; i++ {
	//	go func(i int) {
	//		fmt.Println("B: ", i)
	//		wg.Done()
	//	}(i)
	//}
	//wg.Wait()
	fmt.Print("Hello")
	var a=fmt.Sprintf("%+v","asdasdsad"[0:100])
	fmt.Print(a)

}

// go run 1.HelloWorld.go 可以直接运行go 语言代码
// go build 2.HelloWorld.go 可以生成二进制文件  ./HelloWorld 可以直接执行

// 如果一个标识符以大写字母开头，那么被标识的对象可以被包外的代码所使用(只需要import 这个包)，称作导出。类似于 Java 中的 public
// 如果一个标识符以小写字符开头，那么被标识的对象对包外是不可见的。类似于 Java 中的 protected
// Go 的标识符是 Java 风格的

/*

Go 语言的数据类型
布尔型
	true false
数字类型

	无符号整型 uint8 uint16 uint32 uint64
	有符号整型 int8 int16 int32 int64
	浮点型 float32 float64 complex64 complex128
	其他数字类型 byte(unit8) rune(int32) uint int uintptr(pointer)

字符串类型
	UTF-8
派生类型 零值为nil
	// @Star  派生类型需要更加详细的说明
	指针、数组、struct、Channel、func、切片、interface、map

 */