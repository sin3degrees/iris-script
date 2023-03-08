package controllers

import (
	"iris-script/models"
	"iris-script/service"
	"log"

	"github.com/kataras/iris/v12"
)

type ScriptController struct {
	Ctx     iris.Context
	Service service.ScriptService
}

func NewScriptController() *ScriptController {
	return &ScriptController{Service: service.NewScriptService()}
}

func (g *ScriptController) PostJs() (result models.Result) {
	var m map[string]interface{}
	err := g.Ctx.ReadJSON(&m)
	if err != nil {
		log.Println("ReadJSON Error:", err)
	}
	if m["file"] == "" || m["file"] == nil {
		result.Code = -1
		result.Msg = "参数缺失 file"
		return
	}
	return g.Service.DoJs(m)
}

func (g *ScriptController) PostLua() (result models.Result) {
	var m map[string]interface{}
	err := g.Ctx.ReadJSON(&m)
	if err != nil {
		log.Println("ReadJSON Error:", err)
	}
	if m["file"] == "" || m["file"] == nil {
		result.Code = -1
		result.Msg = "参数缺失 file"
		return
	}
	return g.Service.DoLua(m)
}

func (g *ScriptController) PostPython() (result models.Result) {
	var m map[string]interface{}
	err := g.Ctx.ReadJSON(&m)
	if err != nil {
		log.Println("ReadJSON Error:", err)
	}
	if m["file"] == "" || m["file"] == nil {
		result.Code = -1
		result.Msg = "参数缺失 file"
		return
	}
	return g.Service.DoPython(m)
}

func (g *ScriptController) PostRuby() (result models.Result) {
	var m map[string]interface{}
	err := g.Ctx.ReadJSON(&m)
	if err != nil {
		log.Println("ReadJSON Error:", err)
	}
	if m["file"] == "" || m["file"] == nil {
		result.Code = -1
		result.Msg = "参数缺失 file"
		return
	}
	return g.Service.DoRuby(m)
}

func (g *ScriptController) PostTengo() (result models.Result) {
	var m map[string]interface{}
	err := g.Ctx.ReadJSON(&m)
	if err != nil {
		log.Println("ReadJSON Error:", err)
	}
	if m["file"] == "" || m["file"] == nil {
		result.Code = -1
		result.Msg = "参数缺失 file"
		return
	}
	return g.Service.DoTengo(m)
}
