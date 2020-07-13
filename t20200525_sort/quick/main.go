package main

import "fmt"

func main()  {
	var a   []int = []int{3,6,1,2,5,4}
	fmt.Println(a)
	quickSort(a,0,len(a)-1)


	fmt.Println(a)
}


func quickSort(a []int ,left,right int ){
	tmp := a[left]
	p := left
	i,j  := left,right
	for i<=j  {

		for  j >= p  && a[j] >=  tmp{
				j--
		}
		if j> p {
			a[p],a[j]= a[j],a[p]
			p = j
		}

		for i <= p && a[i] <= tmp{
			i++
		}

		if i <p {
			a[p],a[i] = a[i],a[p]
			p = i
		}
	}
	//a[p] = tmp
	if p-left > 1{
		quickSort(a,left,p-1)
	}
	if right -p  > 1{
		quickSort(a,p+1,right)
	}
}


