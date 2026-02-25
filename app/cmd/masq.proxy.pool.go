package cmd

import (
	"sync"
)

var proxyBufferPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 32*1024) // 32KB
	},
}

type poolWrapper struct{}

func (p poolWrapper) Get() []byte  { return proxyBufferPool.Get().([]byte) }
func (p poolWrapper) Put(b []byte) { proxyBufferPool.Put(b) }
