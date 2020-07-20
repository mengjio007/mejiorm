package schema

import (
	"crawl/dbsupport"
	"go/ast"
	"reflect"
)

type Field struct {
	Name string
	Type string
	Tag string
}
type Schema struct {
	Model interface{}
	Name string
	Fields []*Field
	FieldNames []string
	filedMap map[string]*Field
}

func Parse(dest interface{},d dbsupport.DbType) *Schema{
	modeltype := reflect.Indirect(reflect.ValueOf(dest)).Type()	//indirect获取指针指向的实例
	schema:=&Schema{
		Model:     dest,
		Name:      modeltype.Name(),
		filedMap:  make(map[string]*Field),
	}
	for i:=0;i<modeltype.NumField();i++{
		p:=modeltype.Field(i)
		if !p.Anonymous && ast.IsExported(p.Name){
			field:=&Field{
				Name: p.Name,
				Type: d.DataType(reflect.Indirect(reflect.New(p.Type))),
			}
			if v,ok := p.Tag.Lookup("mejiorm");ok{
				field.Tag = v
			}
			schema.Fields = append(schema.Fields,field)
			schema.FieldNames = append(schema.FieldNames,p.Name)
			schema.filedMap[p.Name]=field
		}
	}
	return schema
}