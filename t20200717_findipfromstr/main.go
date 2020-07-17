package main

import (
	"strconv"
	"fmt"
)

var ipList   = make([]string,10)
func DFS(s string,tmp string,num int)  {
	if num==3 && isValid(s){
		ipList = append(ipList, tmp + s)
		return
	}

	for i:=1;i<=3 && i<len(s);i++{
		subs := s[0:i]
		if isValid(subs){
			DFS(s[i:], tmp + subs+".", num+1);
		}
	}
}

func isValid(s string) (isOK bool)  {
	if len(s) >1 && s[0] == '0'{
		return  false
	}

	a,err:= strconv.Atoi(s)
	if err !=nil{
		return  false
	}

	return  a <=255 && a>=0
}
func main()  {
	DFS("25525511135","",0)
	fmt.Println(ipList)
}