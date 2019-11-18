//請寫出一個可以自行輸入值並且輸出兩個數相加,相減的函數來調用

package main

import （
	"fmt"
	"goproject/ch6/2/utils"
)

func main() {
	var a int
	var b int
	
	fmt.Println("請輸入第一個整數:")
	fmt.Scanln(&a)
	fmt.Println("請輸入第二個整數:")
	fmt.Scanln(&b)
	
	res1,res2 := utils.Getsumandsub(a,b)
	fmt.Printf("res1=%v,res2=%v",res1,res2)
	
	
	//如果只要調用第二個返回值，則可以運用底線_
	
	_,res3 := utils.Getsumandsub(a,b)
	fmt.Printf("res3=%v",res3)
	
}