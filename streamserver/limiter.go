package main

import "log"

type ConnLimiter struct {
	concurrentConn int
	buckert        chan int
}

func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: cc,
		buckert:        make(chan int, cc),
	}
}

func (cl *ConnLimiter) GetConn() bool {
	if len(cl.buckert) >= cl.concurrentConn {
		log.Print("Reached the rate limitation.")
		return false
	}

	cl.buckert <- 1
	return true
}

func (cl *ConnLimiter) ReleaseConn() {
	c := <-cl.buckert
	log.Printf("NewConnction coming: %d", c)
}
