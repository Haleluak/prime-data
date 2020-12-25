package app

import (
	"database/sql"
	sqladapter "github.com/Blank-Xu/sql-adapter"
	"github.com/casbin/casbin/v2"
	"log"
)

func InitCasbin() *casbin.Enforcer {
	db, err := sql.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		panic(err)
	}

	a, err := sqladapter.NewAdapter(db, "sqlite3", "casbin_rules")
	if err != nil {
		panic(err)
	}

	e, err := casbin.NewEnforcer("enforce/model.conf", a)
	if err != nil {
		panic(err)
	}

	if err = e.LoadPolicy(); err != nil {
		log.Println("LoadPolicy failed, err: ", err)
	}

	return e
}
