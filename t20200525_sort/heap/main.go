package main

import "fmt"

func main()  {
	a:= []int{3,5,4,2,1,9}
	heapSort(&a)
	fmt.Println(a)


}

func heapSort(arr *[]int)  {
	arrlen := len(*arr)
	buildMaxHeap(arr,arrlen)
	for i:=arrlen -1;i>=0;i-- {
		swap(arr,0,i)
		arrlen = arrlen -1
		ajustHead(arr,0,arrlen)
	}
}
func buildMaxHeap(arr *[]int,len int ){
	for  i:= len/2;i >= 0;i--{
		ajustHead(arr,i,len)
	}
}

func  ajustHead(arr *[]int,index int ,length int){
	left  := 2 * index + 1
	right := 2 * index + 2
	largest := index
	if (left < length) && ((*arr)[left] > (*arr)[largest]){
		largest = left
	}

	if (right < length)  && (*arr)[right] > (*arr)[largest]{
		largest = right
	}

	if largest != index {
		swap(arr,index,largest)

		//largest 是否打破顺序
		ajustHead(arr,largest,length)
	}
}

//to do swap
func  swap(arr *[]int,i int ,j int){
	//tmp := arr[i]
	//arr[i] = arr[j]
	//arr[j] =  tmp

	(*arr)[i],(*arr)[j] = (*arr)[j],(*arr)[i]
}