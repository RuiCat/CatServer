package bytes

import (
	"reflect"
	"unsafe"
)

// BaseRead 基本类型读
func BaseRead(r *Read, v interface{}) bool {
	return baseRead(r, reflect.ValueOf(v))
}

// baseRead 基本类型读实现
func baseRead(r *Read, v reflect.Value) bool {
	// 长度检测
	if len(*r.Byte) == 0 {
		return false
	}
	// 取指针
	var ptr unsafe.Pointer
	if v.CanAddr() {
		ptr = unsafe.Pointer(v.Addr().Pointer())
	} else if v.IsValid() {
		ptr = unsafe.Pointer(v.Pointer())
	} else {
		return false
	}
	// 基础类型储存
	switch v.Kind() {
	case reflect.Bool:
		*(*bool)(ptr) = r.Bool()
	case reflect.Int:
		*(*int)(ptr) = r.Int()
	case reflect.Int8:
		*(*int8)(ptr) = r.Int8()
	case reflect.Int16:
		*(*int16)(ptr) = r.Int16()
	case reflect.Int32:
		*(*int32)(ptr) = r.Int32()
	case reflect.Int64:
		*(*int64)(ptr) = r.Int64()
	case reflect.Uint:
		*(*uint)(ptr) = r.Uint()
	case reflect.Uint8:
		*(*uint8)(ptr) = r.Uint8()
	case reflect.Uint16:
		*(*uint16)(ptr) = r.Uint16()
	case reflect.Uint32:
		*(*uint32)(ptr) = r.Uint32()
	case reflect.Uint64:
		*(*uint64)(ptr) = r.Uint64()
	case reflect.Float32:
		*(*float32)(ptr) = r.Float32()
	case reflect.Float64:
		*(*float64)(ptr) = r.Float64()
	case reflect.Complex64:
		*(*complex64)(ptr) = r.Complex64()
	case reflect.Complex128:
		*(*complex128)(ptr) = r.Complex128()
	// 附加的处理过程
	case reflect.Ptr:
		return baseRead(r, v.Elem())
	case reflect.Map:
		// 创建
		t := v.Type()
		if v.IsNil() {
			v.Set(reflect.MakeMap(t))
		}
		// 参数
		key := reflect.New(t.Key()).Elem()
		value := reflect.New(t.Elem()).Elem()
		count := r.Int()
		// 循环读取
		for i := 0; i < count; i++ {
			// 读取值
			if !baseRead(r, key) {
				return false
			}
			if !baseRead(r, value) {
				return false
			}
			// 设置值
			v.SetMapIndex(key, value)
		}
	case reflect.String:
		*(*string)(ptr) = string(r.Bytes())
	case reflect.Array:
		n := v.Len()
		for i := 0; i < n; i++ {
			if !baseRead(r, v.Index(i)) {
				return false
			}
		}
	case reflect.Slice:
		n := r.Int()
		v.Set(reflect.MakeSlice(v.Type(), n, n))
		for i := 0; i < n; i++ {
			if !baseRead(r, v.Index(i)) {
				return false
			}
		}
	case reflect.Interface:
		k := reflect.Kind(r.Uint())
		switch k {
		case reflect.Array:
			elem, ok := typetoValue(reflect.Kind(r.Uint()))
			if !ok {
				return false
			}
			array := reflect.New(reflect.ArrayOf(r.Int(), elem.Type()))
			if !baseRead(r, array) {
				return false
			}
			v.Set(array)
		case reflect.Slice:
			elem, ok := typetoValue(reflect.Kind(r.Uint()))
			if !ok {
				return false
			}
			slice := reflect.New(reflect.SliceOf(elem.Type()))
			if !baseRead(r, slice) {
				return false
			}
			v.Set(slice)
		case reflect.Map:
			// 参数类型
			key, ok := typetoValue(reflect.Kind(r.Uint()))
			if !ok {
				return false
			}
			elem, ok := typetoValue(reflect.Kind(r.Uint()))
			if !ok {
				return false
			}
			// 构建
			t := reflect.New(reflect.MapOf(key.Type(), elem.Type()))
			if !baseRead(r, t) {
				return false
			}
			v.Set(t)
		default:
			t, ok := typetoValue(k)
			if !ok {
				return false
			}
			if !baseRead(r, t) {
				return false
			}
			v.Set(t)
		}
	case reflect.Struct:
		// 结构体字段数量
		n := v.NumField()
		for i := 0; i < n; i++ {
			// 处理字段
			if !baseRead(r, v.Field(i)) {
				return false
			}
		}
	default:
		return false
	}
	return true
}

// typetoValue 类型到值
func typetoValue(t reflect.Kind) (v reflect.Value, _ bool) {
	switch t {
	case reflect.Bool:
		v = reflect.ValueOf(new(bool)).Elem()
	case reflect.Int:
		v = reflect.ValueOf(new(int)).Elem()
	case reflect.Int8:
		v = reflect.ValueOf(new(int8)).Elem()
	case reflect.Int16:
		v = reflect.ValueOf(new(int16)).Elem()
	case reflect.Int32:
		v = reflect.ValueOf(new(int32)).Elem()
	case reflect.Int64:
		v = reflect.ValueOf(new(int64)).Elem()
	case reflect.Uint:
		v = reflect.ValueOf(new(uint)).Elem()
	case reflect.Uint8:
		v = reflect.ValueOf(new(uint8)).Elem()
	case reflect.Uint16:
		v = reflect.ValueOf(new(uint16)).Elem()
	case reflect.Uint32:
		v = reflect.ValueOf(new(uint32)).Elem()
	case reflect.Uint64:
		v = reflect.ValueOf(new(uint64)).Elem()
	case reflect.Float32:
		v = reflect.ValueOf(new(float32)).Elem()
	case reflect.Float64:
		v = reflect.ValueOf(new(float64)).Elem()
	case reflect.Complex64:
		v = reflect.ValueOf(new(complex64)).Elem()
	case reflect.Complex128:
		v = reflect.ValueOf(new(complex128)).Elem()
	case reflect.String:
		v = reflect.ValueOf(new(string)).Elem()
	case reflect.Interface:
		v = reflect.ValueOf(new(interface{})).Elem()
	default:
		return v, false
	}
	return v, true
}
