package api

import (
	"go_jin_testing/pkg/app"
	"go_jin_testing/pkg/util"
	"go_jin_testing/service/auth_service"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	appGin := app.Gin{C: c}
	valid := validation.Validation{}
	username := c.PostForm("username")
	password := c.PostForm("password")

	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	if !ok {
		appGin.Response(http.StatusBadRequest, 400, nil)
		return
	}

	authService := auth_service.Auth{Username: username, Password: password}
	isExist, err := authService.Check()

	if err != nil {
		appGin.Response(http.StatusInternalServerError, 500, nil)
		return
	}

	if !isExist {
		appGin.Response(http.StatusUnauthorized, 500, nil)
		return
	}

	token, err := util.GenerateToken(username, password)
	if err != nil {
		appGin.Response(http.StatusInternalServerError, 500, nil)
		return
	}

	appGin.Response(http.StatusOK, 200, map[string]string{
		"token": token,
	})
}
