package geeorm

import (
	"database/sql"
	"mylog"
	"session"
)

type Engine struct {
	db *sql.DB
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		mylog.Error(err)
		return
	}

	if err = db.Ping(); err != nil {
		mylog.Error(err)
		return
	}

	e = &Engine{
		db: db,
	}
	mylog.Info("Connect database success")
	return
}

func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		mylog.Error("Failed to close database")
	}
	mylog.Info("Close database success")
}

func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db)
}
