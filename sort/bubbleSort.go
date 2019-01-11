package sort


//BubbleSort xyz
func BubbleSort(list []int) []int{
    n := len(list)
    sort := true
    for sort {
        sort = false
        for i := 1; i < n; i++ {
            if list[i-1] > list[i] {
                list[i], list[i-1] = list[i-1], list[i]
                sort = true
            }
        }
	}
	return list
}