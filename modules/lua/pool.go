package lua

import (
	"iris-script/conf"
	"sync"

	lua "github.com/yuin/gopher-lua"
)

var luaPool = &LuaPool{}

func init() {
	luaPool = newLuaPool()
}

type LuaStateArray []*lua.LState
type LuaPool struct {
	m     sync.Mutex
	saved LuaStateArray
}

func newLuaPool() *LuaPool {
	return &LuaPool{
		saved: make(LuaStateArray, 0, conf.Sysconfig.Lua.PoolSize),
	}
}

func (p *LuaPool) createLuaState() *lua.LState {
	L := lua.NewState(lua.Options{
		CallStackSize:       conf.Sysconfig.Lua.CallStackSize,
		RegistrySize:        conf.Sysconfig.Lua.RegistrySize,
		SkipOpenLibs:        true,
		IncludeGoStackTrace: true,
	})
	return L
}

func (p *LuaPool) Borrow() *lua.LState {
	p.m.Lock()
	defer p.m.Unlock()
	n := len(p.saved)
	if n == 0 {
		return p.createLuaState()
	}
	L := p.saved[n-1]
	p.saved = p.saved[0 : n-1]
	return L
}

func (p *LuaPool) Return(L *lua.LState) {
	p.m.Lock()
	defer p.m.Unlock()
	p.saved = append(p.saved, L)
}

func (p *LuaPool) ShutDown() {
	for _, L := range p.saved {
		L.Close()
	}
}
