//練習
//從1-100隨機取值，取到99後停止

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	
	var a int
	
	for {
	rand.Seed(time.Now().UnixNano())
	a = rand.Intn(100) + 1
		if a == 99 {
		break
		}
	}
	fmt.Println(a)
}