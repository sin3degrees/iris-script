package utils

import (
	"encoding/json"

	lua "github.com/yuin/gopher-lua"
)

func LuaTableSet(t *lua.LTable, key string, value interface{}) {
	switch value.(type) {
	case nil:
		t.RawSet(lua.LString(key), lua.LNil)
	case bool:
		t.RawSet(lua.LString(key), lua.LBool(value.(bool)))
	case float64:
		t.RawSet(lua.LString(key), lua.LNumber(value.(float64)))
	case float32:
		t.RawSet(lua.LString(key), lua.LNumber(value.(float32)))
	case int:
		t.RawSet(lua.LString(key), lua.LNumber(value.(int)))
	case uint:
		t.RawSet(lua.LString(key), lua.LNumber(value.(uint)))
	case int8:
		t.RawSet(lua.LString(key), lua.LNumber(value.(int8)))
	case uint8:
		t.RawSet(lua.LString(key), lua.LNumber(value.(uint8)))
	case int16:
		t.RawSet(lua.LString(key), lua.LNumber(value.(int16)))
	case uint16:
		t.RawSet(lua.LString(key), lua.LNumber(value.(uint16)))
	case int32:
		t.RawSet(lua.LString(key), lua.LNumber(value.(int32)))
	case uint32:
		t.RawSet(lua.LString(key), lua.LNumber(value.(uint32)))
	case int64:
		t.RawSet(lua.LString(key), lua.LNumber(value.(int64)))
	case uint64:
		t.RawSet(lua.LString(key), lua.LNumber(value.(uint64)))
	case string:
		t.RawSet(lua.LString(key), lua.LString(value.(string)))
	case []byte:
		t.RawSet(lua.LString(key), lua.LString(value.([]byte)))
	default:
		newValue, _ := json.Marshal(value)
		t.RawSet(lua.LString(key), lua.LString(newValue))
	}
}

func LuaArrSet(t *lua.LTable, key int, value interface{}) {
	switch value.(type) {
	case nil:
		t.RawSet(lua.LNumber(key), lua.LNil)
	case bool:
		t.RawSet(lua.LNumber(key), lua.LBool(value.(bool)))
	case float64:
		t.RawSet(lua.LNumber(key), lua.LNumber(value.(float64)))
	case float32:
		t.RawSet(lua.LNumber(key), lua.LNumber(value.(float32)))
	case int:
		t.RawSet(lua.LNumber(key), lua.LNumber(value.(int)))
	case uint:
		t.RawSet(lua.LNumber(key), lua.LNumber(value.(uint)))
	case int8:
		t.RawSet(lua.LNumber(key), lua.LNumber(value.(int8)))
	case uint8:
		t.RawSet(lua.LNumber(key), lua.LNumber(value.(uint8)))
	case int16:
		t.RawSet(lua.LNumber(key), lua.LNumber(value.(int16)))
	case uint16:
		t.RawSet(lua.LNumber(key), lua.LNumber(value.(uint16)))
	case int32:
		t.RawSet(lua.LNumber(key), lua.LNumber(value.(int32)))
	case uint32:
		t.RawSet(lua.LNumber(key), lua.LNumber(value.(uint32)))
	case int64:
		t.RawSet(lua.LNumber(key), lua.LNumber(value.(int64)))
	case uint64:
		t.RawSet(lua.LNumber(key), lua.LNumber(value.(uint64)))
	case string:
		t.RawSet(lua.LNumber(key), lua.LString(value.(string)))
	case []byte:
		t.RawSet(lua.LNumber(key), lua.LString(value.([]byte)))
	default:
		newValue, _ := json.Marshal(value)
		t.RawSet(lua.LNumber(key), lua.LString(newValue))
	}
}

func MapToLuaTable(L *lua.LState, m map[string]interface{}) *lua.LTable {
	t := L.NewTable()
	for k, v := range m {
		switch v.(type) {
		case map[string]interface{}:
			t2 := MapToLuaTable(L, v.(map[string]interface{}))
			t.RawSet(lua.LString(k), t2)
		case []interface{}:
			t2 := ArrToLuaTable(L, v.([]interface{}))
			t.RawSet(lua.LString(k), t2)
		default:
			LuaTableSet(t, k, v)
		}
	}
	return t
}

func ArrToLuaTable(L *lua.LState, arr []interface{}) *lua.LTable {
	t := L.NewTable()
	for k, v := range arr {
		switch v.(type) {
		case map[string]interface{}:
			t2 := MapToLuaTable(L, v.(map[string]interface{}))
			t.RawSet(lua.LNumber(k+1), t2)
		case []interface{}:
			t2 := ArrToLuaTable(L, v.([]interface{}))
			t.RawSet(lua.LNumber(k+1), t2)
		default:
			LuaArrSet(t, (k + 1), v)
		}
	}
	return t
}

func LuaTableToArr(t *lua.LTable) []interface{} {
	len := t.Len()
	arr := make([]interface{}, len)
	t.ForEach(func(l1, l2 lua.LValue) {
		switch l1.(type) {
		case lua.LNumber:
			switch l2.(type) {
			case lua.LBool:
				arr[int(l1.(lua.LNumber))-1] = l2.(lua.LBool)
			case lua.LNumber:
				arr[int(l1.(lua.LNumber))-1] = l2.(lua.LNumber)
			case lua.LString:
				arr[int(l1.(lua.LNumber))-1] = l2.(lua.LString)
			case *lua.LTable:
				len := l2.(*lua.LTable).Len()
				if len > 0 {
					arr2 := LuaTableToArr(l2.(*lua.LTable))
					arr[int(l1.(lua.LNumber))-1] = arr2
				} else {
					m := LuaTableToMap(l2.(*lua.LTable))
					arr[int(l1.(lua.LNumber))-1] = m
				}
			default:
				arr[int(l1.(lua.LNumber))-1] = l2
			}
		}
	})
	return arr
}

func LuaTableToMap(t *lua.LTable) map[string]interface{} {
	m := make(map[string]interface{})
	t.ForEach(func(l1, l2 lua.LValue) {
		switch l1.(type) {
		case lua.LString:
			switch l2.(type) {
			case lua.LBool:
				m[l1.String()] = bool(l2.(lua.LBool))
			case lua.LNumber:
				m[l1.String()] = float64(l2.(lua.LNumber))
			case lua.LString:
				m[l1.String()] = l2.(lua.LString).String()
			case *lua.LTable:
				len := l2.(*lua.LTable).Len()
				if len > 0 {
					arr := LuaTableToArr(l2.(*lua.LTable))
					m[l1.String()] = arr
				} else {
					m2 := LuaTableToMap(l2.(*lua.LTable))
					m[l1.String()] = m2
				}
			default:
				m[l1.String()] = l2
			}
		}
	})
	return m
}
