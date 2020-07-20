package session

import (
	"crawl/log"
	"crawl/schema"
	"fmt"
	"reflect"
	"strings"
)

func(this *Session)Model(value interface{})*Session{
	if this.reftable==nil || reflect.TypeOf(value) != reflect.TypeOf(this.reftable.Model){
		this.reftable = schema.Parse(value,this.Dbtype)
	}
	return this
}

func(this *Session)RefTable()*schema.Schema{
	if this.reftable==nil{
		log.Error("数据表未设置")
	}
	return this.reftable
}

func(this *Session)CreatTable()error{
	table := this.RefTable()
	var columns []string
	for _,field :=range table.Fields{
		columns = append(columns,fmt.Sprintf("%s %s %s",field.Name,field.Type,field.Tag))
	}
	desc:= strings.Join(columns,",")
	_,err:=this.Raw(fmt.Sprintf("CREATE TABLE %s (%s);",table.Name,desc)).Exec()
	return err
}

func(this *Session)DropTable()error{
	_,err:=this.Raw(fmt.Sprintf("DROP TABLE IF EXISTS %s",this.RefTable().Name)).Exec()
	return err
}

func(this *Session)HasTable()bool{
	sql,values:=this.Dbtype.TableSql(this.RefTable().Name)
	row:= this.Raw(sql,values...).QueryRow()
	var tmp string
	_ =row.Scan(&tmp)
	return	tmp==this.RefTable().Name
}