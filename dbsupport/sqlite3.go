package dbsupport

import (
	"fmt"
	"reflect"
	"time"
)

type sqlite3 struct{}

var _ DbType = (*sqlite3)(nil)

func init() {
	RegistDb("sqlite3", &sqlite3{})
}

func(this *sqlite3)DataType(ref reflect.Value) string{
	switch ref.Kind() {
	case reflect.Bool:
		return "bool"
	case reflect.Int,reflect.Int16,reflect.Int8,reflect.Int32,reflect.Uint,reflect.Uint8,
	reflect.Uint16,reflect.Uint32:
		return "interger"
	case reflect.Int64,reflect.Uint64:
		return "bigint"
	case reflect.Float32,reflect.Float64:
		return "real"
	case reflect.String:
		return "text"
	case reflect.Array,reflect.Slice:
		return "blob"
	case reflect.Struct:
		if _,ok := ref.Interface().(time.Time);ok{
			return "datetime"
		}
	}
	panic(fmt.Sprintf("invalid sql type %s (%s)", ref.Type().Name(), ref.Kind()))
}

func(this *sqlite3)TableSql(tablename string) (string, []interface{}) {
	args := []interface{}{tablename}
	return "SELECT name FROM sqlite_master WHERE type='table' and name = ?", args
}