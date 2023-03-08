package tengo

import (
	"errors"
	"iris-script/modules/tengo/lib"
	"os"

	"github.com/sin3degrees/tengo/v2"
)

type virtualMachine struct {
	script   *tengo.Script
	compiled *tengo.Compiled
}

func NewVirtualMachine() *virtualMachine {
	exec := &virtualMachine{
		script: tengoPool.Borrow(),
	}
	exec.init()
	return exec
}

func (e *virtualMachine) init() {
	lib.Load(e.script)
}

// Destroy 销毁虚拟机，为了性能考虑，现在只是将之还给虚拟机池。
func (e *virtualMachine) Destroy() {
	if e.script != nil {
		tengoPool.Return(e.script)
	}
}

// LoadString 加载字符串，并编译成字节码
func (e *virtualMachine) LoadString(source string) {
	//script := tengo.NewScript([]byte(source))
	e.script.SetInput([]byte(source))
}

// LoadFile 加载文件，并编译成字节码
func (e *virtualMachine) LoadFile(filePath string) error {
	code, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	//script := tengo.NewScript(code)
	e.script.SetInput(code)
	return nil
}

// Execute 执行已编译的tengo代码
func (e *virtualMachine) Execute() error {
	if e.script == nil {
		return errors.New("no tengo")
	}
	compiled, err := e.script.Run()
	e.compiled = compiled
	return err
}

// ExecuteString 直接执行字符串
func (e *virtualMachine) ExecuteString(source string) error {
	//e.script = tengo.NewScript([]byte(source))
	e.script.SetInput([]byte(source))
	compiled, err := e.script.Run()
	e.compiled = compiled
	return err
}

// ExecuteFile 直接执行tengo文件
func (e *virtualMachine) ExecuteFile(filePath string) error {
	code, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	//e.script = tengo.NewScript(code)
	e.script.SetInput(code)
	e.compiled, err = e.script.Run()
	return err
}

// Register 注册方法或者变量到tengo
func (e *virtualMachine) Register(name string, value interface{}) error {
	err := e.script.Add(name, value)
	return err
}

// 获取tengo中的全局变量
func (e *virtualMachine) Get(name string) interface{} {
	return e.compiled.Get(name)
}
