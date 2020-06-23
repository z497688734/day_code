package main

import (
	"sync"
	"container/list"
	"fmt"
)


type AtomicInt int64
type LRUCache struct {
	mutex sync.Mutex
	size  int
	cacheList  *list.List
	cache      map[interface{}]*list.Element
	hits, gets  AtomicInt
}

type Entry struct {
	k interface{}
	v interface{}
}
type EntryImpA struct {
	k int
	v string
}

func NewLruCache(_size int)  LRUCache {
	lru := LRUCache{
		size:_size,
		cacheList:list.New(),
		cache:make(map[interface{}] *list.Element,_size),
	}
	return  lru
}
func(lru LRUCache) Get(key string) (interface{},bool)  {
	if  ele,isOk := lru.cache[key];isOk {
		lru.cacheList.MoveToFront(ele)
		return ele.Value.(*Entry).v, true
	}

	return  nil,false
}

func (lru LRUCache)Set(key string,value interface{})  {
	if  ele,isOk := lru.cache[key];isOk{
		lru.cacheList.MoveToFront(ele)
		ele.Value = value
		return
	}

	ele := lru.cacheList.PushFront(&Entry{k:key,v:value})
	lru.cache[key]  = ele
	if lru.cacheList.Len() > lru.size{
		ele := lru.cacheList.Back()
		lru.cacheList.Remove(ele)
		delete(lru.cache,ele.Value.(*Entry).k)
	}
}

func main()  {
	lru :=  NewLruCache(3)
	lru.Set("1","1")
	lru.Set("2","2")

	a,_ := lru.Get("1")
	fmt.Println(a)

	fmt.Println(lru.cacheList.Len())

}
