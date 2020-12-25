package enforce

import (
	"github.com/casbin/casbin/v2"
)

type Enforcer struct {
	enforcer *casbin.Enforcer
}

func CasbinEnforce(e *casbin.Enforcer, sub, obj, act string) (bool, error) {
	return e.Enforce(sub, obj, act)
}