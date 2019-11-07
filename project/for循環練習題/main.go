//練習
//請打印出1~100所有9的倍數的整數各數及總和

package main

import "fmt"

func main() {
	var i int = 1
	var sum int = 0
	var count int = 0
	var max int = 0
	for ; i <= max; i++ {
		if (i % 9) == 0 {
		count++
		sum += i
		}
		fmt.Printf("count=%v, sum=%v\n",count,sum)
	}
}