package orm

import (
	"encoding/json"
	"iris-script/datasource"
	"reflect"
	"time"

	"github.com/dop251/goja"
)

var gdb = datasource.GetDB()

func Load(vm *goja.Runtime) {
	mod := map[string]interface{}{}
	mod["begin"] = begin
	mod["rollback"] = rollback
	mod["commit"] = commit
	mod["err"] = err
	mod["query"] = query
	mod["exec"] = exec
	vm.Set("Orm", mod)
}

func begin() {
	gdb.Begin()
}

func rollback() {
	gdb.Rollback()
}

func commit() {
	gdb.Commit()
}

func err() string {
	return gdb.Error.Error()
}

func query(args ...interface{}) string {
	sql := args[0].(string)
	arr := args[1:]
	rows, err := gdb.Raw(sql, arr...).Rows()
	result := map[string]interface{}{
		"data": nil,
		"flag": true,
		"err":  "",
	}
	if err != nil {
		result["data"] = nil
		result["flag"] = false
		result["err"] = err.Error()
	}
	colTypes, _ := rows.ColumnTypes()                 // 列信息
	var rowParam = make([]interface{}, len(colTypes)) // 传入到 rows.Scan 的参数 数组
	var rowValue = make([]interface{}, len(colTypes)) // 接收数据一行列的数组

	for i, colType := range colTypes {
		rowValue[i] = reflect.New(colType.ScanType())           // 跟据数据库参数类型，创建默认值 和类型
		rowParam[i] = reflect.ValueOf(&rowValue[i]).Interface() // 跟据接收的数据的类型反射出值的地址
	}
	data := []map[string]interface{}{}
	// 遍历
	for rows.Next() {
		rows.Scan(rowParam...) // 赋值到 rowValue 中
		mRow := map[string]interface{}{}
		for i, colType := range colTypes {
			switch rowValue[i].(type) {
			case nil:
				mRow[colType.Name()] = nil
			case bool:
				mRow[colType.Name()] = rowValue[i].(bool)
			case float64:
				mRow[colType.Name()] = float64(rowValue[i].(float64))
			case float32:
				mRow[colType.Name()] = rowValue[i].(float32)
			case int:
				mRow[colType.Name()] = float64(rowValue[i].(int))
			case uint:
				mRow[colType.Name()] = rowValue[i].(uint)
			case int8:
				mRow[colType.Name()] = rowValue[i].(int8)
			case uint8:
				mRow[colType.Name()] = rowValue[i].(uint8)
			case int16:
				mRow[colType.Name()] = rowValue[i].(int16)
			case uint16:
				mRow[colType.Name()] = rowValue[i].(uint16)
			case int32:
				mRow[colType.Name()] = rowValue[i].(int32)
			case uint32:
				mRow[colType.Name()] = rowValue[i].(uint32)
			case int64:
				mRow[colType.Name()] = rowValue[i].(int64)
			case uint64:
				mRow[colType.Name()] = rowValue[i].(uint64)
			case string:
				mRow[colType.Name()] = rowValue[i].(string)
			case []byte:
				mRow[colType.Name()] = string(rowValue[i].([]byte))
			case time.Time:
				mRow[colType.Name()] = rowValue[i].(time.Time).String()
			default:
				newValue, _ := json.Marshal(rowValue[i])
				mRow[colType.Name()] = string(newValue)
			}
		}
		data = append(data, mRow)
	}
	result["data"] = data
	byteJson, _ := json.Marshal(result)
	return string(byteJson)
}

func exec(args ...interface{}) bool {
	sql := args[0].(string)
	arr := args[1:]
	gdb.Exec(sql, arr...)
	err := gdb.Error
	if err != nil {
		return false
	}
	return true
}
