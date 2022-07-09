
package main

import (
	"fmt"
	"os"
)

func main(){

	var s,sep string

	for _,sep=range os.Args[1:]{
		s+=" "+sep

	}
	fmt.Println(s)
}