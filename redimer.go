package redimer

import (
	"reflect"
	"unsafe"
)

// Redim return slice or array dimmensioned with specified dimensions.
func Redim(v interface{}, dims ...int) interface{} {
	in := reflect.ValueOf(v)
	rt := reflect.Indirect(in).Type()
	k := rt.Kind()
	for rt.Kind() == reflect.Array || rt.Kind() == reflect.Slice {
		rt = rt.Elem()
	}
	for i := 0; i < len(dims)-1; i++ {
		rt = reflect.ArrayOf(dims[i], rt)
	}
	var rv reflect.Value
	if k == reflect.Slice {
		rt = reflect.SliceOf(rt)
		rv = reflect.NewAt(rt, unsafe.Pointer(in.Pointer())).Elem()
		rv = rv.Slice(0, dims[len(dims)-1])
	} else {
		rt = reflect.ArrayOf(dims[len(dims)-1], rt)
		rv = reflect.NewAt(rt, unsafe.Pointer(in.Pointer()))
	}
	if rt.Kind() == reflect.Array {
		return rv.Elem().Interface()
	}
	return rv.Interface()
}
