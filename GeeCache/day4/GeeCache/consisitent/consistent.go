package consisitent

import (
	"hash/crc32"
	"sort"
	"strconv"
)

// Hash 哈希函数
type Hash func(data []byte) uint32

// Map map
type Map struct {
	hash     Hash           //计算哈希值的函数
	replicas int            //虚拟节点倍数
	keys     []int          //环，sorted
	hashMap  map[int]string //虚拟节点哈希值->真实节点名称
}

// New 构造
func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hashMap:  make(map[int]string),
		hash:     fn,
	}
	if fn == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// Add 添加，keys:真实节点名称
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.hashMap[hash] = key
			m.keys = append(m.keys, hash)
		}
	}
	sort.Ints(m.keys)
}

// Get 从key计算出哈希值，选择最近的节点
func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}
	hash := int(m.hash([]byte(key)))
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})
	return m.hashMap[m.keys[idx%len(m.keys)]]

}
