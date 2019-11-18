package pipeline

import (
	"bufio"
	"encoding/binary"
	"os"
)

func ReadeResource(filename string) <- chan int  {
	out := make(chan int)
	file,err := os.Open(filename)
	go func() {
		if err == nil {
			defer file.Close()
			reader := bufio.NewReader(file)
			buffer := make([]byte,8)
			n,err := reader.Read(buffer)
			for n > 0 && err == nil{
				out <- int(binary.BigEndian.Uint64(buffer))
				n ,err = reader.Read(buffer)
			}
			close(out)
		}else{
			panic(err)
		}
	}()
	return out
}
