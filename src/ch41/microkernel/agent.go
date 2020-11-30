package microkernel

import "context"

type Event struct {
	name    string
	content string
}

type EventReceiver interface {
	OnEvent(evt Event)
}

type Collector interface {
	Init(receiver EventReceiver) error
	Start(ctx context.Context) error // Context可以cancel 一组协程
	Stop() error
	Destory() error
}

type Agent struct {
	collectors map[string]Collector
	evtBuf     chan Event
	cancel     context.CancelFunc
	ctx        context.Context
	state      int // 1激活，0
}

// Agent绑定这个方法
func (agt *Agent) OnEvent(evt Event) {
	agt.evtBuf <- evt // 将evt 放进chan 管道
}

// 注册

func (agt *Agent) RegisterCollector(name string, collector Collector) error {
	//if agt.state !=
	return nil

}
