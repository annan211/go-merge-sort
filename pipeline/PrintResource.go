package pipeline

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
)

func PrintResource(filename string) {
	go func() {
		if file,err := os.Open(filename);err == nil {
			reader := bufio.NewReader(file)
			buffer := make([]byte,8)
			n,err := reader.Read(buffer)
			index := 0
			for n >0 && err == nil {
				index ++
				fmt.Printf("read value is : %d",int(binary.BigEndian.Uint64(buffer)))
				n,err = reader.Read(buffer)
				if index > 100{
					break
				}
			}
		}
	}()

}
