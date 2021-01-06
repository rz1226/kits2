package kits2

import (
	"sync"
	"sync/atomic"
)

/*
cq := kits.NewCircleQueue(100)
cq.Put("abc")
lists, newestId := cq.GetSeveral(10)

*/

// 只保留最近写入的部分数据
type CircleQueue struct {
	currentID    uint64 //永远的自增长
	size         uint32
	dataNodeList []dataNode
}

type dataNode struct {
	id      uint64
	content string
	mu      *sync.RWMutex
}

func NewCircleQueue(size uint32) *CircleQueue {
	cq := &CircleQueue{}
	cq.currentID = 0
	cq.size = minQuantity(size)
	cq.dataNodeList = initDataNodes(size)
	return cq
}

func initDataNodes(size uint32) []dataNode {
	dataNodeList := make([]dataNode, size)
	for k, _ := range dataNodeList {
		ele := &(dataNodeList[k])
		ele.id = 0
		ele.mu = &sync.RWMutex{}
	}
	return dataNodeList
}

// 把任何类型的数据放入队列
func (c *CircleQueue) Put(val string) uint64 {
	// func AddUint64(addr *uint64, delta uint64) (new uint64)
	nextID := atomic.AddUint64(&c.currentID, 1)
	positionInList := c.getPos(nextID)
	dataNode := &(c.dataNodeList[positionInList])
	dataNode.mu.Lock()
	defer dataNode.mu.Unlock()
	dataNode.id = nextID
	dataNode.content = val
	return nextID
}

func (c *CircleQueue) getPos(ID uint64) uint64 {
	return ID & uint64((c.size - 1)) //实际上就是取模操作
}

// 从队列中取出count个数的数据，返回当初放进去的数据列表，以及最新的id
func (c *CircleQueue) Get(count int) ([]string, uint64) {
	res := make([]string, count)
	currentID := atomic.LoadUint64(&c.currentID)
	for i := 0; i < count-1; i++ {
		pos := c.getPos(currentID - uint64(i))
		dataNode := &(c.dataNodeList[pos])
		dataNode.mu.RLock()
		//如果不通过下面的判断，说明取到了上一圈的数据
		if dataNode.id == currentID-uint64(i) {
			res[i] = dataNode.content
			dataNode.mu.RUnlock()
		} else {
			dataNode.mu.RUnlock()
			break
		}
	}
	return res, currentID
}

//  round 到最近的2的倍数
func minQuantity(v uint32) uint32 {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v++
	return v
}
