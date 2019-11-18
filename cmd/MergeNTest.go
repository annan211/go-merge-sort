package main

import (
	"Merging/pipeline"
	"fmt"
	"math/rand"
)
/**
this is a test APP , just merge 20 numbers
 */
func main() {

	result := []<-chan int{}
	for i :=0;i<2;i++{
		source := createChan(10)
		result = append(result,pipeline.InMemorySort(source))
	}

	sortChan := pipeline.MergeN(result...)
	index := 0
	for v := range sortChan {
		index ++
		fmt.Printf("sortChan idnex is %d ,v is %d \n",index,v)
	}
}

func createChan(num int)<- chan int {
	out := make(chan int)
	go func() {
		for i:=0;i<num;i++{
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}
