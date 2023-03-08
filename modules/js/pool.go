package js

import (
	"iris-script/conf"
	"sync"

	"github.com/dop251/goja"
)

var jsPool = &JsPool{}

func init() {
	jsPool = newJsPool()
}

type JsVmArray []*goja.Runtime
type JsPool struct {
	m     sync.Mutex
	saved JsVmArray
}

func newJsPool() *JsPool {
	return &JsPool{
		saved: make(JsVmArray, 0, conf.Sysconfig.Js.PoolSize),
	}
}

func (p *JsPool) createJsVm() *goja.Runtime {
	R := goja.New()
	return R
}

func (p *JsPool) Borrow() *goja.Runtime {
	p.m.Lock()
	defer p.m.Unlock()
	n := len(p.saved)
	if n == 0 {
		return p.createJsVm()
	}
	R := p.saved[n-1]
	p.saved = p.saved[0 : n-1]
	return R
}

func (p *JsPool) Return(R *goja.Runtime) {
	p.m.Lock()
	defer p.m.Unlock()
	p.saved = append(p.saved, R)
}

func (p *JsPool) ShutDown() {
	for _, R := range p.saved {
		R.ClearInterrupt()
	}
}
