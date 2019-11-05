// 練習
// while break

package main

import "fmt"

func main() {

	var i int = 1
	for {
		fmt.Println("do hi",i) //do..while 先執行在判斷
		//迭代變量已經寫在底下
		if i > 10 {
		break
		}
		fmt.Println("hi",i) //while 先判斷在執行
		i++
	}
}