package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"prime-data/models"
	"prime-data/pkg/errors"
	gohttp "prime-data/pkg/http"
	"prime-data/schema"
	"prime-data/services"
)

type Auth struct {
	service services.IAuthService
}

func NewAuthAPI(service services.IAuthService) *Auth {
	return &Auth{service: service}
}

func (a *Auth) Login(c *gin.Context) gohttp.Response{
	var item schema.LoginBodyParam
	if err := c.BindJSON(&item); err != nil {
		return gohttp.Response{
			Error: errors.InvalidParams.New(),
		}
	}

	ctx := c.Request.Context()
	user, err := models.QueryUser(ctx, models.Client,item.Username, item.Password)
	if err != nil{
		return gohttp.Response{
			Error: errors.InvalidParams.New(),
		}
	}

	if user.ID == 0 {
		return gohttp.Response{
			Error: errors.ErrorAuth.New(),
		}
	}

	tokenInfo, err := a.service.Login(ctx, user)
	if err != nil {
		fmt.Print(err)
		return gohttp.Response{
			Error: err,
		}
	}

	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  tokenInfo,
	}
}

func (a *Auth) Hello(c *gin.Context) gohttp.Response{
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  "hahahaha",
	}
}