package main

import "fmt"

//用函數的方式打印出可以自由輸入層數的空心金字塔

func jzt(level int){
	
	for i := 1; i <= level; i++{
		for k := 1; k <= level-i; k++{
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++{
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

//用函數表示九九乘法表，可以輸入各倍數並打印出來
func nine2(j int){
	
	for i := 1; i <= 9; i++{
		fmt.Printf("%v*%v=%v\n",j,i,i*j)
	}
}

func main() {
	
	var a int
	var b int
	fmt.Println("請輸入一個層數：")
	fmt.Scanln(&a)
	jzt(a)
	fmt.Println("請輸入99乘法表的哪一個哪一個倍數")
	fmt.Println(&b)
	nine2(b)
}