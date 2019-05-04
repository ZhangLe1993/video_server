package task

import (
	"errors"
	"log"
	"testing"
	"time"
)

func TestExecutor(t *testing.T) {
	// 循环网data channel 中发送  数据
	d := func(dc chan interface{}) error {
		for i := 0; i < 30; i++ {
			dc <- i
			log.Printf("Dispatch send: %v \n", i)
		}
		return nil
	}

	e := func(dc chan interface{}) error {
	ownLoop:
		for {
			select {
			case d := <-dc:
				log.Printf("Executor received : %v \n", d)
			default:
				break ownLoop
			}
		}
		return errors.New("Executor")
	}

	executor := NewExecutor(30, false, d, e)
	go executor.start()
	time.Sleep(time.Second * 3)
}
