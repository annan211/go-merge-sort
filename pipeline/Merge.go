package pipeline

import "fmt"

func Merge(n1,n2 <- chan int )<- chan int {
   out := make(chan int)
   index := 0
   go func() {
   		v1,ok1 := <- n1
   		v2,ok2 := <- n2
   		for (ok1) || (ok2) {
			if !ok1 || (ok1 && v1 <= v2) {
				out <- v2
				//fmt.Printf("Merge v is %d \n",v2)
				v2,ok2 = <- n2
			}else{
				out <- v1
				//fmt.Printf("Merge v is %d \n",v1)
				v1,ok1 = <- n1
			}
			index ++
		}
   		close(out)
   		fmt.Printf("merge done***********************************************")
   }()
   return out
}
