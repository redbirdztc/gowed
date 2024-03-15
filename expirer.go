package gowed

import (
	"reflect"
	"time"
)

type Expirer interface {
	IsExpired() bool
}

type ExpiryKeeper interface {
	TimeoutAt(time.Time)
	Timeout() time.Time
}

func HasExpiredData(v any) bool {
	if v == nil {
		return false
	}

	tp := reflect.TypeOf(v)
	expirableType := reflect.TypeOf((*Expirer)(nil)).Elem()

	if tp.Implements(expirableType) {
		return v.(Expirer).IsExpired()
	}

	switch tp.Kind() {

	case reflect.Array:
		fallthrough
	case reflect.Slice:
		slice := reflect.ValueOf(v)

		for i := 0; i < slice.Len(); i++ {
			if HasExpiredData(slice.Index(i).Interface()) {
				return true
			}
		}
		return false
	case reflect.Struct:
		// check fields in struct is Expirable
		for i := 0; i < tp.NumField(); i++ {
			if HasExpiredData(reflect.ValueOf(v).Field(i).Interface()) {
				return true
			}
		}
		return false
	case reflect.Map:
		// check fields in map is Expirable
		for _, key := range reflect.ValueOf(v).MapKeys() {
			if HasExpiredData(reflect.ValueOf(v).MapIndex(key).Interface()) {
				return true
			}
		}
		return false
	case reflect.Ptr:
		value := reflect.ValueOf(v)
		return HasExpiredData(value.Elem().Interface())
	default:
		return false
	}
}
