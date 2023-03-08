package utils

func ArrToIntf[T any](arr []T) []interface{} {
	arr2 := make([]interface{}, len(arr))
	for i := 0; i < len(arr); i++ {
		arr2[i] = arr[i]
	}
	return arr2
}
