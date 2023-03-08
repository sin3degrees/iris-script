package tengo

import (
	"iris-script/conf"
	"sync"

	"github.com/sin3degrees/tengo/v2"
)

var tengoPool = &TengoPool{}

type TengoVmArray []*tengo.Script
type TengoPool struct {
	m     sync.Mutex
	saved TengoVmArray
}

func init() {
	tengoPool = newTengoPool()
}

func newTengoPool() *TengoPool {
	return &TengoPool{
		saved: make(TengoVmArray, 0, conf.Sysconfig.Tengo.PoolSize),
	}
}

func (p *TengoPool) createTengoVm() *tengo.Script {
	T := tengo.NewScript([]byte{})
	T.EnableFileImport(true)
	return T
}

func (p *TengoPool) Borrow() *tengo.Script {
	p.m.Lock()
	defer p.m.Unlock()
	n := len(p.saved)
	if n == 0 {
		return p.createTengoVm()
	}
	T := p.saved[n-1]
	p.saved = p.saved[0 : n-1]
	return T
}

func (p *TengoPool) Return(T *tengo.Script) {
	p.m.Lock()
	defer p.m.Unlock()
	p.saved = append(p.saved, T)
}

func (p *TengoPool) ShutDown() {
	p.saved = TengoVmArray{}
}
