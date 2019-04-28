package main

import "log"

//流控机制，使用go语言的channel机制代替加解锁机制
type ConnLimiter struct {
	conCurrentConn int
	bucket         chan int
}

func InitConnLimiter(conCurrent int) *ConnLimiter {
	return &ConnLimiter{conCurrent, make(chan int, conCurrent)}
}

func (connLimiter *ConnLimiter) GetConn() bool {
	if len(connLimiter.bucket) >= connLimiter.conCurrentConn {
		log.Printf("连接数超过限制")
		return false
	}
	connLimiter.bucket <- 1
	return true
}

func (connLimiter *ConnLimiter) ReleaseConn() {
	conn := <-connLimiter.bucket
	log.Printf("释放连接：%d", conn)
}
