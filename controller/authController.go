package ctrl

import (
	//seelog "github.com/cihub/seelog"
	"github.com/flosch/pongo2"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jmoiron/sqlx"
	"golang-gin/csrf"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

type AuthController struct {
}

func (ct *AuthController) GetRegister(c *gin.Context) {
	csrfToken := csrf.GetToken(c)
	c.HTML(http.StatusOK, "auth/register.html", pongo2.Context{
		"token": csrfToken,
	})
}
