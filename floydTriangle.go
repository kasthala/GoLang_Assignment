package main

import "fmt"


func main(){
	var i = 0
	var x int = 1
	fmt.Println("enter rows n: ")
	fmt.Scan(&i)
	for index := 1; index <= i; index++ {
		for index1 := 1; index1 < index ; index1++ {
			fmt.Printf("%d ", x)
			x++
		}
		fmt.Println("")	
	}
}