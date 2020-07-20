package dbsupport

import "reflect"

type DbType interface {
	DataType(ref reflect.Value) string
	TableSql(tablename string) (string, []interface{})
}

var SqlTypeList = map[string]DbType{}

func RegistDb(name string,dbtype DbType){
	SqlTypeList[name] = dbtype
}

func GetDb(name string)(dbtype DbType,ok bool){
	dbtype,ok = SqlTypeList[name]
	return
}