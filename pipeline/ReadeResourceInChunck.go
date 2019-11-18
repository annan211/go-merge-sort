package pipeline

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
)

func ReadeResourceInChunck(filename string,filesize int,chunckCount int)<- chan int  {
     result := []<-chan int {}
     for i:= 0 ;i< chunckCount;i++{
     	fmt.Printf("this is %d chunck \n ",i)
     	readResult := chunckRead(i,filename,filesize,chunckCount)
     	fmt.Printf("readResult len is %d ",len(readResult))
        sortResult := InMemorySort(readResult)
        result = append(result,sortResult)
	 }
     return MergeN(result...)
}


func chunckRead(index int,filename string,filesize int,chunckCount int)<- chan int  {
	out := make(chan int)
	go func() {
		chunckSize := filesize / chunckCount
		file,err := os.Open(filename)
		if err == nil {
			file.Seek(int64(index*chunckSize),0)
			reader := bufio.NewReader(file)
			buffer := make([]byte,8)
			n,err := reader.Read(buffer)
			readBytes := n
			for n > 0 && err == nil {
				//fmt.Printf("chunckRead index is %d v is %d \n",index,int(binary.BigEndian.Uint64(buffer)))
				out <- int(binary.BigEndian.Uint64(buffer))
				n,err = reader.Read(buffer)
				readBytes += n
				if readBytes >= chunckSize || err != nil {
					break
				}
			}
			close(out)
		}else{
			panic(err)
		}
	}()
	return out
}
