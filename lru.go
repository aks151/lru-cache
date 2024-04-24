package main


type node struct {
	key int 
	val int
	next *node
	prev *node
}

type LRUCache struct {
	mp map[int]*node
	cap int
}

func NewLRUCache(c int) *LRUCache {
	return &LRUCache{
		cap: c,
		mp: make(map[int]*node, c),
	}
}

func (lru *LRUCache) addNode() {

}

func (lru *LRUCache) deleteNode() {
	
}

func (lru *LRUCache) putData(key_, val_ int) {

}

func (lru *LRUCache) getData(key_ int) int {
	return -1
}

func (lru *LRUCache) deleteData() {

}