package main

import (
	"strings"
	"fmt"
)

var sensitiveWord = make(map[string]interface{})
var Set = make(map[string]interface{})
const InvalidWords = " ,~,!,@,#,$,%,^,&,*,(,),_,-,+,=,?,<,>,.,—,，,。,/,\\,|,《,》,？,;,:,：,',‘,；,“,"
var InvalidWord = make(map[string]interface{}) //无效词汇，不参与敏感词汇判断直接忽略

var a = make(map[string]string)

//生成违禁词集合
func AddSensitiveToMap(set map[string]interface{} ){
	for key := range set {
		str := []rune(key)
		nowMap := sensitiveWord
		for i := 0; i < len(str); i++ {
			if _,ok := nowMap[string(str[i])]; !ok {//如果该key不存在，
				thisMap := make(map[string]interface{})
				thisMap["isEnd"] = false
				nowMap[string(str[i])] = thisMap
				nowMap = thisMap
			}else {
				nowMap = nowMap[string(str[i])].(map[string]interface{})
			}
			if i == len(str)-1 {
				nowMap["isEnd"] = true
			}
		}

	}
}

//敏感词汇转换为*
func ChangeSensitiveWords(txt string,sensitive map[string]interface{}) (word string){
	str := []rune(txt)
	nowMap := sensitive
	start := -1
	tag := -1
	for i := 0; i < len(str); i++ {
		if _, ok:= InvalidWord[(string(str[i]))]; ok {
			continue//如果是无效词汇直接跳过
		}
		if thisMap, ok :=nowMap[string(str[i])].(map[string]interface{}); ok {
			//记录敏感词第一个文字的位置
			tag++
			if  tag == 0 {
				start = i

			}
			//判断是否为敏感词的最后一个文字
			if isEnd, _ := thisMap["isEnd"].(bool);isEnd {
				//将敏感词的第一个文字到最后一个文字全部替换为“*”
				for y := start; y < i+1; y++ {
					str[y] = 42
				}
				//重置标志数据
				nowMap = sensitive
				start = -1
				tag = -1

			}else{//不是最后一个，则将其包含的map赋值给nowMap
				nowMap = nowMap[string(str[i])].(map[string]interface{})
			}

		}else{  //如果敏感词不是全匹配，则终止此敏感词查找。从开始位置的第二个文字继续判断
			if start != -1 {
				i = start + 1
			}
			//重置标志参数
			nowMap = sensitive
			start = -1
			tag = -1
		}
	}

	return string(str)
}
func main() {
	words := strings.Split(InvalidWords,",")
	for _, v := range words {
		InvalidWord[v] = nil
	}
	//Set["你妈逼的"] = nil
	//Set["你妈"] = nil
	Set["狗日的"] = nil
	Set["狗日的傻"] = nil

	AddSensitiveToMap(Set)
	fmt.Println(sensitiveWord["狗"])

	fmt.Println(sensitiveWord)
	text := "文明用语你&* 妈, 逼的你这个狗 日的，怎么这么傻啊。我也是服了，我日,这些话我都说不出口"
	fmt.Println(ChangeSensitiveWords(text,sensitiveWord))
	//
	//

	fmt.Println("........begin test.........")
	var sensitiveWord2 = make(map[string]interface{})
	nowMap := sensitiveWord2

	thisMap := make(map[string]interface{})
	thisMap["isEnd"] = false
	nowMap["A"] = thisMap
	fmt.Printf("2%p\n",nowMap)
	fmt.Printf("%v\n",sensitiveWord2)

	//sensitiveWord2 中只有一个元素,即 A=>thisMap;
	//执行nowMap = thisMap;即:nowMap 指向了sensitiveWord2的元素;
	//因此;对nowMap的任何修改;其实也就是修改了sensitiveWord2
	nowMap = thisMap

	thisMap = make(map[string]interface{})
	thisMap["isEnd"] = false
	nowMap["B"] = thisMap
	fmt.Println("........end test.........")

}