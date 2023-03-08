package lua

import (
	"fmt"
	"iris-script/conf"
	"iris-script/modules/lua/orm"
	"log"
	"os"
	"path/filepath"

	"github.com/tengattack/gluacrypto"
	libs "github.com/vadv/gopher-lua-libs"
	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

type TableMap map[string]interface{}
type virtualMachine struct {
	L *lua.LState
	F *lua.LFunction
}

func NewVirtualMachine() *virtualMachine {
	vm := &virtualMachine{
		L: luaPool.Borrow(),
	}
	vm.init()
	return vm
}

// GetRunPath 获取程序执行目录
func GetRunPath() string {
	path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return path
}

func (e *virtualMachine) init() {

	e.L.OpenLibs()

	libs.Preload(e.L)

	gluacrypto.Preload(e.L)

	//lua_debugger.Preload(e.L)
	e.L.PreloadModule("orm", orm.Loader)

	e.RegisterFunction("GetLuaPath", func(vm *lua.LState) int {
		// 绝对路径
		//e.L.Push(lua.LString(GetRunPath() + "/script"))
		e.L.Push(lua.LString(conf.Sysconfig.Lua.Path))
		return 1
	})
}

// Destroy 销毁虚拟机，为了性能考虑，现在只是将之还给虚拟机池。
func (e *virtualMachine) Destroy() {
	if e.L != nil {
		luaPool.Return(e.L)
	}
}

// LoadString 加载字符串，并编译成字节码
func (e *virtualMachine) LoadString(source string) error {
	var lFunc *lua.LFunction
	var err error
	if lFunc, err = e.L.LoadString(source); err != nil {
		return err
	}

	e.F = lFunc

	return nil
}

// LoadFile 加载文件，并编译成字节码
func (e *virtualMachine) LoadFile(filePath string) error {
	var lFunc *lua.LFunction
	var err error
	if lFunc, err = e.L.LoadFile(filePath); err != nil {
		return err
	}

	e.F = lFunc

	return nil
}

// Execute 执行已编译的lua代码
func (e *virtualMachine) Execute() error {
	if err := e.doCompiledFile(); err != nil {
		return err
	}
	return nil
}

// ExecuteString 直接执行字符串
func (e *virtualMachine) ExecuteString(source string) error {
	if err := e.L.DoString(source); err != nil {
		return err
	}
	return nil
}

// ExecuteFile 直接执行lua文件
func (e *virtualMachine) ExecuteFile(filePath string) error {
	if err := e.L.DoFile(filePath); err != nil {
		return err
	}
	return nil
}

// CallFunction 调用lua当中的方法
func (e *virtualMachine) CallFunction(name string, args ...interface{}) {
	var lArgs []lua.LValue
	for _, arg := range args {
		lArgs = append(lArgs, e.ConvertToLValue(arg))
	}

	if err := e.L.CallByParam(lua.P{
		Fn:      e.L.GetGlobal(name),
		NRet:    1,    // 指定返回值数量
		Protect: true, // 如果出现异常，是panic还是返回err
	}, lArgs...); err != nil { // 传递输入参数：10
		panic(err)
	}
}

func (e *virtualMachine) PCall(f string, args ...interface{}) {
	e.L.Push(e.L.GetGlobal(f))
	for _, arg := range args {
		val := e.ConvertToLValue(arg)
		e.L.Push(val)
	}
	if err := e.L.PCall(len(args), -1, nil); err != nil {
		log.Println("lua pcall err:", err)
	}
}

func (e *virtualMachine) PCall2(f string, args ...lua.LValue) {
	e.L.Push(e.L.GetGlobal(f))
	for _, arg := range args {
		e.L.Push(arg)
	}
	if err := e.L.PCall(len(args), -1, nil); err != nil {
		log.Println("lua pcall2 err:", err)
	}
}

func (e *virtualMachine) PCall3(f lua.LValue, args ...lua.LValue) {
	e.L.Push(f)
	for _, arg := range args {
		e.L.Push(arg)
	}
	if err := e.L.PCall(len(args), -1, nil); err != nil {
		log.Println("lua pcall3 err:", err)
	}
}

// RegisterFunction 注册一个全局的方法到lua
func (e *virtualMachine) RegisterFunction(name string, fn lua.LGFunction) {
	e.L.SetGlobal(name, e.L.NewFunction(fn))
}

// RegisterModule 注册一个模块到lua
func (e *virtualMachine) RegisterModule(name string, mod lua.LGFunction) {
	e.L.Push(e.L.NewFunction(mod))
	e.L.Push(lua.LString(name))
	e.L.Call(1, 0)
}

// BindStruct 绑定一个struct到lua，可以双向操作。
func (e *virtualMachine) BindStruct(name string, data interface{}) {
	e.L.SetGlobal(name, luar.New(e.L, data))
}

// GetLuaTableToStruct 从lua读取一个table到go的struct
func (e *virtualMachine) GetLuaTableToStruct(name string, out interface{}) error {
	return gluamapper.Map(e.L.GetGlobal(name).(*lua.LTable), &out)
}

// 执行已经编译的字节码
func (e *virtualMachine) doCompiledFile() error {
	e.L.Push(e.F)
	return e.L.PCall(0, lua.MultRet, nil)
}

// ConvertToLValue 将go的值转换为LValue
func (e *virtualMachine) ConvertToLValue(val interface{}) lua.LValue {
	if val == nil {
		return lua.LNil
	}
	switch v := val.(type) {
	case lua.LValue:
		return v
	case bool:
		return lua.LBool(v)
	case float32:
		return lua.LNumber(v)
	case float64:
		return lua.LNumber(v)
	case int:
		return lua.LNumber(v)
	case int8:
		return lua.LNumber(v)
	case int16:
		return lua.LNumber(v)
	case int32:
		return lua.LNumber(v)
	case int64:
		return lua.LNumber(v)
	case uint8:
		return lua.LNumber(v)
	case uint16:
		return lua.LNumber(v)
	case uint32:
		return lua.LNumber(v)
	case uint64:
		return lua.LNumber(v)
	case string:
		return lua.LString(v)
	case []byte:
		ud := e.L.NewUserData()
		ud.Value = v
		return ud
	case map[string]interface{}:
		return e.ConvertToLTable(v)
	case []interface{}:
		lt := e.L.NewTable()
		for k, v := range v {
			lt.RawSetInt(k+1, e.ConvertToLValue(v))
		}
		return lt
	default:
		return nil
	}
}

// ConvertFromLValue 将LValue转换为go的值
func (e *virtualMachine) ConvertFromLValue(lv lua.LValue) interface{} {
	switch v := lv.(type) {
	case *lua.LNilType:
		return nil
	case *lua.LUserData:
		return v.Value
	case lua.LBool:
		return bool(v)
	case lua.LString:
		return string(v)
	case lua.LNumber:
		f64i := float64(v)
		I64i := int64(v)
		if f64i == float64(I64i) {
			return I64i
		}
		return f64i
	case *lua.LTable:
		maxn := v.MaxN()
		if maxn == 0 {
			// table
			ret := make(map[string]interface{})
			v.ForEach(func(key, value lua.LValue) {
				keyStr := fmt.Sprint(e.ConvertFromLValue(key))
				ret[keyStr] = e.ConvertFromLValue(value)
			})
			return ret
		} else {
			// array
			ret := make([]interface{}, 0, maxn)
			for i := 1; i <= maxn; i++ {
				ret = append(ret, e.ConvertFromLValue(v.RawGetInt(i)))
			}
			return ret
		}
	default:
		log.Println("error lua type ", lv)
		return nil
	}
}

// ConvertToLTable 将go的map转换成LTable
func (e *virtualMachine) ConvertToLTable(data map[string]interface{}) *lua.LTable {
	lt := e.L.NewTable()

	for k, v := range data {
		lt.RawSetString(k, e.ConvertToLValue(v))
	}

	return lt
}

// ConvertFromLTable 将LTable转换成map。
func (e *virtualMachine) ConvertFromLTable(lv *lua.LTable) map[string]interface{} {
	returnData, _ := e.ConvertFromLValue(lv).(map[string]interface{})
	return returnData
}
