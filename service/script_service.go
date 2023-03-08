package service

import (
	"encoding/json"
	"iris-script/conf"
	"iris-script/models"
	"iris-script/modules/js"
	"iris-script/modules/lua"
	"iris-script/modules/tengo"

	T "github.com/sin3degrees/tengo/v2"
)

type ScriptService interface {
	DoJs(m map[string]interface{}) (result models.Result)
	DoLua(m map[string]interface{}) (result models.Result)
	DoPython(m map[string]interface{}) (result models.Result)
	DoRuby(m map[string]interface{}) (result models.Result)
	DoTengo(m map[string]interface{}) (result models.Result)
}

type scriptService struct{}

func NewScriptService() ScriptService {
	return &scriptService{}
}

func (s scriptService) DoJs(m map[string]interface{}) (result models.Result) {
	if m["file"] == nil {
		result.SetErr(-1, "file参数为空")
		return
	}
	vm := js.NewVirtualMachine()
	defer vm.Destroy()
	inData, err := json.Marshal(m)
	if err != nil {
		result.SetErr(-1, err.Error())
		return
	}
	vm.Register("inData", string(inData))
	err = vm.LoadFile(conf.Sysconfig.Js.Path + m["file"].(string) + ".js")
	if err != nil {
		result.SetErr(-1, err.Error())
		return
	}
	err = vm.Execute()
	if err != nil {
		result.SetErr(-1, err.Error())
		return
	}
	sJson := vm.Get("outData").String()
	outMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(sJson), &outMap)
	if err != nil {
		result.SetErr(-1, err.Error())
	}
	result.Code = int(outMap["code"].(float64))
	result.Msg = outMap["msg"].(string)
	result.Data = outMap["data"]
	return
}

func (s scriptService) DoLua(m map[string]interface{}) (result models.Result) {
	if m["file"] == nil {
		result.SetErr(-1, "file参数为空")
		return
	}
	vm := lua.NewVirtualMachine()
	defer vm.Destroy()
	err := vm.LoadFile(conf.Sysconfig.Lua.Path + m["file"].(string) + ".lua")
	if err != nil {
		result.SetErr(-1, err.Error())
		return
	}
	inData := vm.ConvertToLTable(m)
	outData := vm.L.NewTable()
	vm.L.SetGlobal("inData", inData)
	vm.L.SetGlobal("outData", outData)
	if err := vm.Execute(); err != nil {
		result.SetErr(-1, err.Error())
	} else {
		outMap := vm.ConvertFromLTable(outData)
		result.Code = int(outMap["code"].(int64))
		result.Msg = outMap["msg"].(string)
		result.Data = outMap["data"]
	}
	return
}

func (s scriptService) DoPython(m map[string]interface{}) (result models.Result) {
	return
}

func (s scriptService) DoRuby(m map[string]interface{}) (result models.Result) {
	return
}

func (s scriptService) DoTengo(m map[string]interface{}) (result models.Result) {
	if m["file"] == nil {
		result.SetErr(-1, "file参数为空")
		return
	}
	vm := tengo.NewVirtualMachine()
	defer vm.Destroy()
	vm.Register("inData", m)
	vm.Register("outData", map[string]interface{}{})
	err := vm.LoadFile(conf.Sysconfig.Tengo.Path + m["file"].(string) + ".tengo")
	if err != nil {
		result.SetErr(-1, err.Error())
		return
	}
	err = vm.Execute()
	if err != nil {
		result.SetErr(-1, err.Error())
		return
	}
	outMap := vm.Get("outData").(*T.Variable).Map()
	result.Code = int(outMap["code"].(int64))
	result.Msg = outMap["msg"].(string)
	result.Data = outMap["data"]
	return
}
