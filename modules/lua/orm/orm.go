package orm

import (
	"encoding/json"
	"iris-script/datasource"
	"iris-script/modules/lua/utils"
	"reflect"
	"time"

	lua "github.com/yuin/gopher-lua"
)

var gdb = datasource.GetDB()

func Loader(L *lua.LState) int {
	// register functions to the table
	mod := L.SetFuncs(L.NewTable(), exports)
	// register other stuff
	L.SetField(mod, "name", lua.LString("orm"))

	// returns the module
	L.Push(mod)
	return 1
}

var exports = map[string]lua.LGFunction{
	"begin":    begin,
	"rollback": rollback,
	"commit":   commit,
	"err":      err,
	"query":    query,
	"exec":     exec,
}

func begin(L *lua.LState) int {
	gdb.Begin()
	return 0
}

func rollback(L *lua.LState) int {
	gdb.Rollback()
	return 0
}

func commit(L *lua.LState) int {
	gdb.Commit()
	return 0
}

func err(L *lua.LState) int {
	L.Push(lua.LString(gdb.Error.Error()))
	return 1
}

func query(L *lua.LState) int {
	sql := L.ToString(1)
	tParam := L.ToTable(2)
	var arr = []interface{}{}
	if tParam != nil {
		arr = utils.LuaTableToArr(tParam)
	}
	rows, err := gdb.Raw(sql, arr...).Rows()
	if err != nil {
		L.Push(lua.LNil)
		return 1
	}

	colTypes, _ := rows.ColumnTypes()                 // 列信息
	var rowParam = make([]interface{}, len(colTypes)) // 传入到 rows.Scan 的参数 数组
	var rowValue = make([]interface{}, len(colTypes)) // 接收数据一行列的数组

	for i, colType := range colTypes {
		rowValue[i] = reflect.New(colType.ScanType())           // 跟据数据库参数类型，创建默认值 和类型
		rowParam[i] = reflect.ValueOf(&rowValue[i]).Interface() // 跟据接收的数据的类型反射出值的地址
	}

	t := L.NewTable()
	// 遍历
	for rows.Next() {
		rows.Scan(rowParam...) // 赋值到 rowValue 中
		tRow := L.NewTable()
		for i, colType := range colTypes {
			switch rowValue[i].(type) {
			case nil:
				tRow.RawSet(lua.LString(colType.Name()), lua.LNil)
			case bool:
				tRow.RawSet(lua.LString(colType.Name()), lua.LBool(rowValue[i].(bool)))
			case float64:
				tRow.RawSet(lua.LString(colType.Name()), lua.LNumber(rowValue[i].(float64)))
			case float32:
				tRow.RawSet(lua.LString(colType.Name()), lua.LNumber(rowValue[i].(float32)))
			case int:
				tRow.RawSet(lua.LString(colType.Name()), lua.LNumber(rowValue[i].(int)))
			case uint:
				tRow.RawSet(lua.LString(colType.Name()), lua.LNumber(rowValue[i].(uint)))
			case int8:
				tRow.RawSet(lua.LString(colType.Name()), lua.LNumber(rowValue[i].(int8)))
			case uint8:
				tRow.RawSet(lua.LString(colType.Name()), lua.LNumber(rowValue[i].(uint8)))
			case int16:
				tRow.RawSet(lua.LString(colType.Name()), lua.LNumber(rowValue[i].(int16)))
			case uint16:
				tRow.RawSet(lua.LString(colType.Name()), lua.LNumber(rowValue[i].(uint16)))
			case int32:
				tRow.RawSet(lua.LString(colType.Name()), lua.LNumber(rowValue[i].(int32)))
			case uint32:
				tRow.RawSet(lua.LString(colType.Name()), lua.LNumber(rowValue[i].(uint32)))
			case int64:
				tRow.RawSet(lua.LString(colType.Name()), lua.LNumber(rowValue[i].(int64)))
			case uint64:
				tRow.RawSet(lua.LString(colType.Name()), lua.LNumber(rowValue[i].(uint64)))
			case string:
				tRow.RawSet(lua.LString(colType.Name()), lua.LString(rowValue[i].(string)))
			case []byte:
				tRow.RawSet(lua.LString(colType.Name()), lua.LString(rowValue[i].([]byte)))
			case time.Time:
				tRow.RawSet(lua.LString(colType.Name()), lua.LString(rowValue[i].(time.Time).String()))
			default:
				newValue, _ := json.Marshal(rowValue[i])
				tRow.RawSet(lua.LString(colType.Name()), lua.LString(newValue))
			}
		}
		t.Append(tRow)
	}
	L.Push(t)
	return 1
}

func exec(L *lua.LState) int {
	sql := L.ToString(1)
	tParam := L.ToTable(2)
	arr := utils.LuaTableToArr(tParam)
	gdb.Exec(sql, arr...)
	err := gdb.Error
	if err != nil {
		L.Push(lua.LFalse)
	} else {
		L.Push(lua.LTrue)
	}
	return 1
}
