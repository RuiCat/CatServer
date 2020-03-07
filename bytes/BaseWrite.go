package bytes

import (
	"reflect"
)

// BaseWrite 基础类型写
func BaseWrite(w *Write, v interface{}) bool {
	return baseWrite(w, reflect.ValueOf(v))
}

// baseWrite 基础类型写实现
func baseWrite(w *Write, v reflect.Value) bool {
	in := v.Interface()
	switch v.Kind() {
	case reflect.Bool:
		w.Bool(in.(bool))
	case reflect.Int:
		w.Int(in.(int))
	case reflect.Int8:
		w.Int8(in.(int8))
	case reflect.Int16:
		w.Int16(in.(int16))
	case reflect.Int32:
		w.Int32(in.(int32))
	case reflect.Int64:
		w.Int64(in.(int64))
	case reflect.Uint:
		w.Uint(in.(uint))
	case reflect.Uint8:
		w.Uint8(in.(uint8))
	case reflect.Uint16:
		w.Uint16(in.(uint16))
	case reflect.Uint32:
		w.Uint32(in.(uint32))
	case reflect.Uint64:
		w.Uint64(in.(uint64))
	case reflect.Float32:
		w.Float32(in.(float32))
	case reflect.Float64:
		w.Float64(in.(float64))
	case reflect.Complex64:
		w.Complex64(in.(complex64))
	case reflect.Complex128:
		w.Complex128(in.(complex128))
	case reflect.String:
		w.Bytes([]byte(in.(string)))
	// 处理其他类型
	case reflect.Ptr:
		// 取指针指向元素
		return baseWrite(w, v.Elem())
	case reflect.Map:
		// map 数量
		w.Int(v.Len())
		// 内容
		iter := v.MapRange()
		for iter.Next() {
			if !baseWrite(w, iter.Key()) {
				return false
			}
			if !baseWrite(w, iter.Value()) {
				return false
			}
		}
	case reflect.Array, reflect.Slice:
		// 数组数量
		count := v.Len()
		w.Int(count)
		for i := 0; i < count; i++ {
			if !baseWrite(w, v.Index(i)) {
				return false
			}
		}
	case reflect.Interface:
		// 接口指向元素
		e := v.Elem()
		// 写入元素类型
		w.Uint(uint(e.Kind()))
		// 写入元素数据
		if !baseWrite(w, e) {
			return false
		}
	case reflect.Struct:
		n := v.NumField()
		for i := 0; i < n; i++ {
			// 处理字段
			if !baseWrite(w, v.Field(i)) {
				return false
			}
		}
	default:
		return false
	}
	return true
}
