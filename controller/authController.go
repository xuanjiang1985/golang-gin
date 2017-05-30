package ctrl

import (
	valid "github.com/asaskevich/govalidator"
	seelog "github.com/cihub/seelog"
	"github.com/flosch/pongo2"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jmoiron/sqlx"
	"golang-gin/csrf"
	"gopkg.in/gin-gonic/gin.v1"
	"log"
	"net/http"
	"strings"
)

type AuthController struct {
}

func init() {
	valid.SetFieldsRequiredByDefault(true)
}

func (ct *AuthController) GetRegister(c *gin.Context) {
	csrfToken := csrf.GetToken(c)
	c.HTML(http.StatusOK, "auth/register.html", pongo2.Context{
		"token": csrfToken,
	})
}

func (ct *AuthController) PostRegister(c *gin.Context) {
	//开启日志
	seelog.ReplaceLogger(logger)
	defer seelog.Flush()

	type Validator struct {
		Name             string `valid:"required~昵称：不能为空,length(4|15)~昵称：4至15个字符之间"`
		Email            string `valid:"required~邮箱：不能为空,email~邮箱：必须是email格式"`
		Password         string `valid:"required~密码：不能为空,length(6|150)~密码：至少6个字符"`
		Confirm_password string `valid:"required~确认密码：不能为空,length(6|150)~确认密码：至少6个字符"`
	}
	data := &Validator{
		Name:             c.PostForm("昵称"),
		Email:            c.PostForm("邮箱"),
		Password:         c.PostForm("密码"),
		Confirm_password: c.PostForm("密码确认"),
	}
	//true or false of validator
	csrfToken := csrf.GetToken(c)
	_, err := valid.ValidateStruct(data)
	if err != nil {
		msg := strings.Trim(err.Error(), ";")
		message := strings.Split(msg, ";")
		log.Println(data)
		c.HTML(200, "auth/register.html", pongo2.Context{
			"token":  csrfToken,
			"errors": message,
			"data":   data,
		})
		return
	}
	c.String(http.StatusOK, "done")
}
