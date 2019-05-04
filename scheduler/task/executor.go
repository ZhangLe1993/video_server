package task

type Executor struct {
	Boss  chan string
	Error chan string
	Data  chan interface{}
	// 大小
	size int
	// 是否持久，长期存活
	persistent bool
	//分发器
	Dispatch func(d chan interface{}) error
	//执行器
	Execute func(d chan interface{}) error
}

func NewExecutor(s int, lived bool, d func(dc chan interface{}) error, e func(dc chan interface{}) error) *Executor {
	return &Executor{
		Boss:       make(chan string, 1),
		Error:      make(chan string, 1),
		Data:       make(chan interface{}, s),
		persistent: lived,
		size:       s,
		Dispatch:   d,
		Execute:    e,
	}
}

func (executor *Executor) startDispatch() {
	defer func() {
		// 如果不是长期存活的数据要回收资源
		if !executor.persistent {
			close(executor.Boss)
			close(executor.Data)
			close(executor.Error)
		}
	}()

	for {
		//select 随机执行一个可运行的 case。如果没有 case 可运行，它将阻塞，直到有 case 可运行
		select {
		//
		case boss := <-executor.Boss:
			{
				// 如果就绪准备分发状态
				if boss == READY_TO_DISPATCH {
					err := executor.Dispatch(executor.Data)
					if err != nil {
						executor.Error <- CLOSE
					} else {
						executor.Boss <- READY_TO_EXECUTE
					}
				}
				// 如果就绪准备执行状态
				if boss == READY_TO_EXECUTE {
					err := executor.Execute(executor.Data)
					if err != nil {
						executor.Error <- CLOSE
					} else {
						executor.Boss <- READY_TO_DISPATCH
					}
				}
			}
		case err := <-executor.Error:
			{
				if err == CLOSE {
					return
				}
			}

		default:

		}
	}
}

func (executor *Executor) Start() {
	executor.Boss <- READY_TO_DISPATCH
	executor.startDispatch()
}
