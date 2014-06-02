package main

import (
    "fmt"
)


func main() {
	ans := make([]int, 0)
	for i := 1; i < 100; i++ {
		if i % 5 == 0 && i % 3 == 0{
			ans = append(ans, i)	
		}		
	}
	fmt.Println(ans)
}




