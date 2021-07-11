package orm

import (
	"database/sql"
	"wfw/orm/session"
)

type Engine struct {
	db *sql.DB
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		return
	}
	if err = db.Ping(); err != nil {
		return
	}
	e = &Engine{db: db}
	return
}

func (engine *Engine) Close()(err error) {
	return  engine.db.Close()
}

func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db)
}