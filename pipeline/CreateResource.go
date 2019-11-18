package pipeline

import (
	"bufio"
	"encoding/binary"
	"math/rand"
	"os"
)

func CreateResource(filename string,numCount int){
	file,err := os.Create(filename)
	defer file.Close()
	if err == nil {
	   writer := bufio.NewWriter(file)
	   resource := RandomResource(numCount)
	   buffer := make([]byte,8)
	   for v := range resource{
	   	 binary.BigEndian.PutUint64(buffer,uint64(v))
	   	 writer.Write(buffer)
	   }
	   writer.Flush()
	}else{
		panic(err)
	}
}

func RandomResource(numCount int)<- chan int{
	out := make(chan int)
	go func() {
		for i := 0;i<numCount;i++{
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}
