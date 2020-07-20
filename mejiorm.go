package mejiorm

import (
	"crawl/dbsupport"
	"crawl/log"
	"crawl/session"
	"database/sql"
)

type Engine struct {
	db *sql.DB
	dbtype dbsupport.DbType
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	dial,ok:=dbsupport.GetDb(driver)
	if !ok{
	log.Errorf("暂不支持此数据库",driver)
	return
	}
	e = &Engine{db: db,dbtype:dial}
	log.Info("数据库连接成功")
	return
}

func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		log.Error("数据库关闭失败")
	}
	log.Info("数据库关闭成功")
}

func (e *Engine) NewSession() *session.Session {
	return session.New(e.db,e.dbtype)
}
