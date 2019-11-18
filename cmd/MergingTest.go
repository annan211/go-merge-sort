package main

import (
	"Merging/pipeline"
	"fmt"
)

func main() {

	var filenameIn string = "large.in"
	var numCount int = 100000000
	var filenameOut string = "large.out"

	/**
	 1 随机生成函数
	 2 将随机数写入文件
	 */
	fmt.Println("create file begin")
	pipeline.CreateResource(filenameIn,numCount)
	fmt.Println("create file end")

	/**
	1 分块读取文件
	2 读取的数据 在内部排序
	3 排序之后的chan进行合并
	4 数组chan 进行merge
	5 写入文件
	6 打印文件确认排序结果
	*/
	//resource := pipeline.ReadeResource(filename)
	//sortResource := pipeline.InMemorySort(resource)

	filesize := numCount * 8
	chunckCount := 4
	fmt.Println("read file in chunck begin")
	resource := pipeline.ReadeResourceInChunck(filenameIn,filesize,chunckCount)
	fmt.Println("read file in chunck end")

	//for v := range resource {
	//	fmt.Printf("test v is %d \n",v)
	//}

	fmt.Println("write file begin")
	pipeline.WriteResource(resource,filenameOut)
	fmt.Println("write file end")

	pipeline.PrintResource(filenameOut)

}
