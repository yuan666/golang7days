package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type Hash func(data []byte) uint32

type Map struct {
	hash     Hash
	replicas int
	keys     []int
	hashMap  map[int]string
}

func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}

	sort.Ints(m.keys)
}

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

/***
测试结果
[root@localhost consistenthash]# go test  -v
=== RUN   TestHashing
--- PASS: TestHashing (0.00s)
PASS
ok  	Day4/geecache/consistenthash	0.002s
+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
一致性哈希(consistent hashing)的原理以及为什么要使用一致性哈希。
实现一致性哈希代码，添加相应的测试用例

	一致性哈希算法是啥？为什么要使用一致性哈希算法？这和分布式有什么关系？

一致性哈希算法将 key 映射到 2^32 的空间中，将这个数字首尾相连，形成一个环。
	计算节点/机器(通常使用节点的名称、编号和 IP 地址)的哈希值，放置在环上。
	计算 key 的哈希值，放置在环上，顺时针寻找到的第一个节点，就是应选取的节点/机器。

一致性哈希算法，在新增/删除节点时，只需要重新定位该节点附近的一小部分数据，而不需要重新定位所有的节点
	数据倾斜问题，引入了虚拟节点。
*/
