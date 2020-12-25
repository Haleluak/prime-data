package app

import (
	"database/sql"
	"github.com/casbin/casbin/v2"
	sqladapter "github.com/Blank-Xu/sql-adapter"
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

	return e
}
