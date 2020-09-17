package lru

/* import( */
	// "fmt"
/* ) */

type Node struct {
	Key	  interface{}
	Value interface{}
	next, prev *Node
}

type Cache struct {
	head, tail *Node
	size int
	cacheItem []*Node
	kv map[interface{}]*Node
}

func New(n int)*Cache{
	c := new(Cache)
	c.size = n
	c.head = new(Node)
	c.tail = new(Node)
	c.head.prev = nil
	c.head.next = c.tail
	c.tail.prev = c.head
	c.tail.next = nil
	c.cacheItem = make([]*Node, n)
	for i := 0; i < n ; i++{
		c.cacheItem[i] = &Node{}
	}
	c.kv = make(map[interface{}]*Node)
	return c
}
func (cache* Cache)Set(key interface{}, value interface{})bool {
	node, ok := cache.kv[key]
	if !ok {
		l := len(cache.cacheItem)
		if l==0 {
			node = cache.tail.prev
			cache.detach(node)
			delete(cache.kv, node.Key)
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

   /*  for k, v := range cache.kv{ */
	//     fmt.Println("key:",k, "value:", v)
	// }
	// fmt.Println("==============================")
    return true
}

func (cache* Cache)Get(key interface{})(value interface{}, err error){
	node, ok := cache.kv[key]
	if !ok {
		value = nil
		err = NotExist
		return 
	}
	cache.detach(node)
	cache.attach(node)
	value = node.Value
	err = nil
	return
}

func (cache* Cache)detach(node* Node){
	if node == nil {
		return 
	}
	node.prev.next = node.next 
	node.next.prev = node.prev
}

func (cache* Cache)attach(node* Node){
	if node == nil{
		return 
	}
	n := cache.head.next
	node.prev = cache.head
	node.next = cache.head.next
	cache.head.next = node
	n.prev = node
}






