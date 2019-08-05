package main

import "log"

// ConnLimiter 连接限流器
type ConnLimiter struct {
	concurrentConn int
	buckert        chan int
}

// NewConnLimiter 创建型的连接限流器
func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: cc,
		buckert:        make(chan int, cc),
	}
}

// GetConn 获取一个连接
func (cl *ConnLimiter) GetConn() bool {
	if len(cl.buckert) >= cl.concurrentConn {
		log.Print("Reached the rate limitation.")
		return false
	}

	cl.buckert <- 1
	return true
}

// ReleaseConn 释放一个连接
func (cl *ConnLimiter) ReleaseConn() {
	c := <-cl.buckert
	log.Printf("NewConnction coming: %d", c)
}
