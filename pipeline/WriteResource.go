package pipeline

import (
	"bufio"
	"encoding/binary"
	"os"
)

func WriteResource(input <- chan int , filename string){
	if file,err := os.Create(filename);err == nil {
		writer := bufio.NewWriter(file)
		buffer := make([]byte,8)
		for v := range input {
			binary.BigEndian.PutUint64(buffer,uint64(v))
			writer.Write(buffer)
		}
		writer.Flush()
	}else{
		panic(err)
	}
}
