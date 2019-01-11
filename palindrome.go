package main

import "fmt"


func main(){
	var num , temp, rem, rev int
	fmt.Print("Enter the number: ")
	fmt.Scan(&num)
	temp = num
	for{
		rem = num%10
		rev = rev*10 + rem
		num = num/10
		if(num==0){
			break
		}
	}

	if(temp==rev){
		fmt.Printf("%d is  a palindrome", temp)
	}else{
		fmt.Printf("%d is not a palindrome", temp)
	}
}