package pipeline

func MergeN(inputs ... <- chan int)<- chan int {
	//for i:= 0 ;i<len(inputs);i++{
	//	for v:= range inputs[i]{
	//		fmt.Printf("inputs[%d] v is %d \n",i,v)
	//	}
	//}
	//fmt.Printf("inputs length is %d \n",len(inputs))

	if len(inputs)==1{
		return inputs[0]
	}
	m := len(inputs)/2
	return Merge(MergeN(inputs[:m]...),
		MergeN(inputs[m:]...))
}
