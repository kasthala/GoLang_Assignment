package main

import (
	"fmt"
	"./sort"
)

func main(){
	list := []int{1,2,3,4,5,6,7,8,9,12,14,165,155,155,125644,648664,654033}
	list = sort.BubbleSort(list)
	begin :=0
	last := len(list) -1
	var search int
	fmt.Println("Enter the number to search: ")
	fmt.Scan(&search)

	for begin <= last{
		x := (begin + last)/2
		if(list[x]< search){
			begin = x +1
		}else{
			last = x -1
		}
	}
	if begin == len(list) || list[begin]!= search{
		fmt.Printf("value not found %d" ,search)
	}else{
		fmt.Printf("value found %d at %d", search, begin)
	}
}

