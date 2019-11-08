//練習
package main

import fmt

func main() {
	var a int = 300
	var ptr *int = a //錯在哪？如何修正

}

func main() {
	var a int = 300
	var ptr float32 = &a //錯在哪？如何修正
}

func main() {
	var a int = 300
	var b int = 400
	var ptr int = &a 
	ptr = 100 a =
	ptr = &b
	ptr = 200 b=
	fmt.Printf(a=%d,b=%d,ptr=%d,a,b,ptr)
	//輸出什麼內容
}



//解答-------------------









/*
package main

import "fmt"

func main() {
	var a int = 1
	var ptr *int = &a 

	fmt.Printf("a=%d ptr=%v",i,ptr)

}

func main() {
	var a int = 300
	var ptr *int = &a  //¿ù¦b­þ¦p¦ó­×¥¿

	fmt.Printf("a=%d,ptr=%v",a,ptr)

}*/