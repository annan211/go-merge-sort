package pipeline

import (
	"sort"
)

func InMemorySort(inputs <- chan int)<- chan int {
	out := make(chan int)
	go func() {
		array := []int{}
		for v := range inputs{
			array = append(array,v)
		}
		sort.Ints(array)
		for _,v := range array {
			//fmt.Printf("InMemorySort v is %d \n",v)
			out <- v
		}
		close(out)
	}()
	return out
}
