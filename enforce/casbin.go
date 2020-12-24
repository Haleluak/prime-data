package enforce

import (
	"github.com/casbin/casbin"
)

type Enforcer struct {
	enforcer *casbin.Enforcer
}

func (e *Enforcer) Enforce()  {
	user, path, method := "admin", "/admin/route", "POST"
	result := e.enforcer.Enforce(user, path, method)
}