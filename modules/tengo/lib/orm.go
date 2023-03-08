package lib

import (
	"encoding/json"
	"fmt"
	"iris-script/datasource"
	"reflect"
	"time"

	"github.com/sin3degrees/tengo/v2"
	"github.com/sin3degrees/tengo/v2/stdlib"
)

var ormModule = map[string]interface{}{
	"Begin":    stdlib.FuncAR(begin),
	"Rollback": stdlib.FuncAR(rollback),
	"Commit":   stdlib.FuncAR(commit),
	"Err":      stdlib.FuncARS(err),
	"Query":    query,
	"Exec":     execSQL,
}
var gdb = datasource.GetDB()

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

func query(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) < 1 {
		return nil, tengo.ErrWrongNumArguments
	}
	y1, ok := args[0].(*tengo.String)
	if !ok {
		return nil, tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string",
			Found:    args[0].TypeName(),
		}

	}
	sql := y1.Value
	arr := []interface{}{}
	for i := 1; i < len(args); i++ {
		v1, ok1 := args[i].(*tengo.Int)
		v2, ok2 := args[i].(*tengo.Float)
		v3, ok3 := args[i].(*tengo.String)
		if ok1 {
			arr = append(arr, v1.Value)
		} else if ok2 {
			arr = append(arr, v2.Value)
		} else if ok3 {
			arr = append(arr, v3.Value)
		} else {
			return nil, tengo.ErrInvalidArgumentType{
				Name:     fmt.Sprintf("args[%d]", i),
				Expected: "string or int or float",
				Found:    args[i].TypeName(),
			}
		}
	}
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
	data := []tengo.Object{}
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
		tmp, _ := tengo.FromInterface(mRow)
		data = append(data, tmp)
	}
	result["data"] = data
	return tengo.FromInterface(result)
}

func execSQL(args ...tengo.Object) (ret tengo.Object, err error) {
	if len(args) < 1 {
		return nil, tengo.ErrWrongNumArguments
	}
	y1, ok := args[0].(*tengo.String)
	if !ok {
		return nil, tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string",
			Found:    args[0].TypeName(),
		}

	}
	sql := y1.Value
	arr := []interface{}{}
	for i := 1; i < len(args); i++ {
		v1, ok1 := args[i].(*tengo.Int)
		v2, ok2 := args[i].(*tengo.Float)
		v3, ok3 := args[i].(*tengo.String)
		if ok1 {
			arr = append(arr, v1.Value)
		} else if ok2 {
			arr = append(arr, v2.Value)
		} else if ok3 {
			arr = append(arr, v3.Value)
		} else {
			return nil, tengo.ErrInvalidArgumentType{
				Name:     fmt.Sprintf("args[%d]", i),
				Expected: "string or int or float",
				Found:    args[i].TypeName(),
			}
		}
	}

	gdb.Exec(sql, arr...)
	err = gdb.Error
	if err != nil {
		return tengo.FalseValue, err
	}
	return tengo.TrueValue, nil
}
