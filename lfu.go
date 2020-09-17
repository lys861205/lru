package lru

import (
  "fmt"
)

type Elem struct {
  head, tail *Node
  items int
}

type LFUCache struct {
	head, tail *Node
	size int
	cacheItem []*Node
	kv map[interface{}]*Node
  freq map[int]*Elem
}

func NewElem()*Elem {
    e := new(Elem) 
    e.head = new(Node)
    e.tail = new(Node)
    e.head.prev = nil
    e.head.next = e.tail
    e.tail.prev = e.head 
    e.tail.next = nil
    e.items = 0
    return e
}

func NewLFUCache(n int)*LFUCache{
	c := new(LFUCache)
	c.size = n
	c.cacheItem = make([]*Node, n)
	for i := 0; i < n ; i++{
		c.cacheItem[i] = &Node{Freq:0}
	}
	c.kv = make(map[interface{}]*Node)
  c.freq = make(map[int]*Elem)
	return c
}

func (cache* LFUCache)Get(key interface{})(value interface{}, err error){
	node, ok := cache.kv[key]
	if !ok {
		value = nil
		err = NotExist
		return 
	}
  err = nil
	value = node.Value
  cache.detach(node)
  cache.attach(node)
	return
}

func (cache* LFUCache) MinFreq() (node *Node) {
  for i := 1; i > 0; i++ {
    ele, ok := cache.freq[i]
    if !ok || ele.items == 0{
      continue
    }
    node = ele.tail.prev
    break
  }
  return
}

func (cache* LFUCache)Set(key interface{}, value interface{})bool {
	node, ok := cache.kv[key]
	if !ok {
		l := len(cache.cacheItem)
		if l==0 {
			node = cache.MinFreq()
			cache.detach(node)
			delete(cache.kv, node.Key)
      node.Freq = 0
		} else {
			node = cache.cacheItem[l-1]
			cache.cacheItem = cache.cacheItem[:l-1]
		}
		cache.kv[key] = node
	} else {
		cache.detach(node)
	}
	if node == nil {
		// fmt.Println("node is empty")
		return false
	}
	node.Key   = key
	node.Value = value
	cache.attach(node)
  return true
}

func (cache* LFUCache)Debug() {
  fmt.Println("-------------Debug----------------")
  fmt.Println(cache.freq)
  for k, v := range cache.freq {
    cur := v.head.next
    for cur != v.tail {
      fmt.Println(k, cur.Key, cur.Value, v.items)
      cur = cur.next
    }
  }
  fmt.Println("-------------End Debug----------------")
}

func (cache* LFUCache)detach(node* Node){
	if node == nil {
		return 
	}
  ele, ok := cache.freq[node.Freq]
  if ok {
    ele.items--
  }
	node.prev.next = node.next 
	node.next.prev = node.prev
}

func (cache* LFUCache)attach(node* Node) {
	if node == nil{
		return 
	}
  node.Freq++
  ele, ok := cache.freq[node.Freq]
  if !ok {
    ele = NewElem()
    cache.freq[node.Freq] = ele
  }
  n := ele.head.next
  node.prev = ele.head
  node.next = ele.head.next
  ele.head.next = node
  n.prev = node
  ele.items++
}






