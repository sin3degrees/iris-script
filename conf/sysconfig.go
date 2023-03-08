package conf

import (
	"io/ioutil"

	jsoniter "github.com/json-iterator/go"
)

var Sysconfig = &sysconfig{}

func init() {
	//指定对应的json配置文件
	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic("Sys config read err")
	}
	err = jsoniter.Unmarshal(b, Sysconfig)
	if err != nil {
		panic(err)
	}

}

type sysconfig struct {
	Common common `json:"Common"`
	DB     db     `json:"DB"`
	Js     js     `json:"Js"`
	Lua    lua    `json:"Lua"`
	Python python `json:"Python"`
	Ruby   ruby   `json:"Ruby"`
	Tengo  tengo  `json:"Tengo"`
}

type common struct {
	Port      string `json:"Port"`
	AuthCheck bool   `json:"AuthCheck"`
}

type db struct {
	Type     string `json:"Type"`
	UserName string `json:"UserName"`
	Password string `json:"Password"`
	Ip       string `json:"Ip"`
	Port     string `json:"Port"`
	DBName   string `json:"DBName"`
	MaxLife  int    `json:"MaxLife"`
	MaxIdle  int    `json:"MaxIdle"`
	MaxOpen  int    `json:"MaxOpen"`
}

type js struct {
	Path     string `json:"Path"`
	PoolSize int    `json:"PoolSize"`
}

type lua struct {
	Path          string `json:"Path"`
	PoolSize      int    `json:"PoolSize"`
	CallStackSize int    `json:"CallStackSize"`
	RegistrySize  int    `json:"RegistrySize"`
}

type python struct {
	Path string `json:"Path"`
}

type ruby struct {
	Path string `json:"Path"`
}

type tengo struct {
	Path     string `json:"Path"`
	PoolSize int    `json:"PoolSize"`
}
