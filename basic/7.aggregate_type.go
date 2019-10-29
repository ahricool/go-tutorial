package main
import "fmt"


// 未定义长度的数组只能传给不限制数组长度的函数
// 定义了长度的数组只能传给限制了相同数组长度的函数
func avg(arr [10]int )float32 {
	var sum int =0
	var count int =0
	fmt.Println(arr)
	for _,value:= range arr{
		sum+=value
		count++
	}
	fmt.Println(sum,count)
	return float32(sum)/float32(count)
}


func main(){

	// 数组的声明
	var n [10] int
	// 数组的初始化，未初始化的部分会自动补零值
	var m = [10] float64{1000,20,30,5}
	// 根据初始化元素个数 自动判断数组大小
	var k=[...]float32{20,10,5,2.5}

	fmt.Println(n)
	fmt.Println(m)
	fmt.Println(k)

	var i,j int
	for i =0;i<10;i++{
		n[i]=i+100
	}
	for j=0;j<10;j++{
		fmt.Printf("n[%d]=%d\n",j,n[j])
	}

	// 多维数组
	var multidim [10][5][4] int
	// 初始化多维数组
	var multidim2 =[3][4]int{
		{0,1,3,4},
		{1,2,3,4},
		{9,2,3,4},
	}

	fmt.Println(multidim)
	fmt.Println(multidim2)

	var arr=[10]int{1,2,3,4,5,6,7,8,9,10}
	var res=avg(arr)
	fmt.Println(res)





}