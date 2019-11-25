package main

import "fmt"

//請編寫一個函數swap(n1,n2 int),可以交換n1,n2的值

func swap(n1,n2 int)(a int,b int){
	var n3 int
	n3 = n1 + n2
	n2 = n3 - n2
	n1 = n3 - n2
	
	a = n1
	b = n2
	
	return a,b
}


//在編寫一個函數swap2(n3,n4 int),交換n3,n4的值


func main() {
	res1,res2 := swap(1,2)
	fmt.Printf("res1=%v,res2=%v",res1,res2)
}



