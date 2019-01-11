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
	minValue := list[0]
	maxValue := list[len(list)-1]
	var search int
	fmt.Println("Enter the number to search: ")
	fmt.Scan(&search)
	for {
        if search < minValue {
			fmt.Printf("value not found %d" ,search)
			break
        }
 
        if search > maxValue {
           	fmt.Printf("value not found %d" ,search)
			break
        }
 
        var sIndex int
        if last == begin {
            sIndex = last
        } else {
            size := last - begin
            offset := int(float64(size-1) * (float64(search-minValue) / float64(maxValue-minValue)))
            sIndex = begin + offset
        }
 
        if list[sIndex] == search {
            for sIndex > 0 && list[sIndex-1] == search {
                sIndex--
            }
			fmt.Printf("value found %d at %d", search, sIndex)
			break
        }
 
        if list[sIndex] > search {
            last = sIndex - 1
            maxValue = list[last]
        } else {
            begin = sIndex + 1
            minValue = list[begin]
        }
    }
}

