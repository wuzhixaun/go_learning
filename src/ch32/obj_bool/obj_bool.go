package obj_bool

//+--------------+
//|              |
//|              +------------------------>
//|              |
//+--------------+                                       X X XX XX XXXXXX
//XXX                 XXX
//XXX                       X       XXXXXXXX
//XX                          X     XXX       XXX
//XX                           X   XX             XX
//XX                             X X                 X
//X                             XXX                  X
//X                              XX                   X
//X                              X                   XX
//X                                                  X
//X                                                XX
//X                                               XX
//X                                              XX
//X                                            XX
//X                                        XXX
//X                                     XXXX
//XX                                XXXXX
//X                           XXXXXX
//XXXXXXXXXXXXXXXXXXXX XXXXXXX
//XXXX
//XXXXXXXXXX

import (
	"errors"
	"time"
)

type ResuableObj struct {
}

type ObjectPool struct {
	bufChan chan *ResuableObj
}

func NewObjPool(numSize int) *ObjectPool {
	objpool := ObjectPool{}

	objpool.bufChan = make(chan *ResuableObj, numSize)

	for i := 0; i < numSize; i++ {
		objpool.bufChan <- &ResuableObj{}
	}

	return &objpool
}

/*
 * 从线程池中获取某一个对象
 */
func (p *ObjectPool) GetObRes(timeOut time.Duration) (*ResuableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeOut): // 超时控制
		return nil, errors.New("timeout")
	}
}

// 释放对象
func (p *ObjectPool) ReleaseObject(obj *ResuableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("size is out")
	}
}
