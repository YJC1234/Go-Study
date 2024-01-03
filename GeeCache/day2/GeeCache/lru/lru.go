package lru

import "container/list"

//LRU:最近最久未使用算法
//维持一个双端队列，将使用的元素移动到链表头，插入元素时如果内存不够，去除末尾的元素

// Cache is a LRU cache(not safe for concurrent access)
type Cache struct {
	maxBytes int64 //为0时无上限
	nBytes   int64
	ll       *list.List
	cache    map[string]*list.Element

	OnEvicted func(key string, value Value)
}

// Entry cache条目
type entry struct {
	key   string
	value Value
}

// Value 用Len计算占了多少字节
type Value interface {
	Len() int
}

// New 创建新的Cache
func New(maxBytes int64, onEvicted func(key string, value Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		nBytes:    0,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

// Get 获取key对应的value,并移动到链表头
func (c *Cache) Get(key string) (Value, bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry) //类型断言，Value为entry指针类型
		return kv.value, true
	}
	return nil, false
}

// 删除尾节点
func (c *Cache) removeOld() {
	ele := c.ll.Back()
	c.ll.Remove(ele)
	kv := ele.Value.(*entry)
	delete(c.cache, kv.key)
	c.nBytes -= int64(len(kv.key) + kv.value.Len())
	if c.OnEvicted != nil {
		c.OnEvicted(kv.key, kv.value)
	}
}

// Add 添加key,value，如果存在则更新。
func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		kv.value = value
		c.nBytes += int64(len(key) + value.Len())
	} else {
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.nBytes += int64(len(key) + value.Len())
	}
	if c.maxBytes != 0 && c.maxBytes < c.nBytes { //remove尾元素
		c.removeOld()
	}
}
