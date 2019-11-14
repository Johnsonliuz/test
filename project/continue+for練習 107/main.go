//某人有100000元，過路費每次要收5%，少於五萬則每次收1000，請問可以過幾次

package main

import "fmt"

func main() {

	var a int = 100000 //money
	var b int
	
	for{
		if a > 50000 {
			a = a - a * 5 / 100
			b++
		} else if a <= 50000 && a >= 1000{
			a = a - 1000
			b++
		} else if a <1000 {
			break
		}
	}
	fmt.Printf("共通過了%v次\n",b)
	fmt.Printf("還剩%v元\n",a)
}