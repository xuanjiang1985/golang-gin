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
		Name             string `valid:"required,length(3|15)"`
		Email            string `valid:"required,email"`
		Password         string `valid:"required,length(6|150)"`
		Confirm_password string `valid:"required,length(6|150)"`
	}
	data := &Validator{
		Name:             c.PostForm("昵称"),
		Email:            c.PostForm("邮箱"),
		Password:         c.PostForm("密码"),
		Confirm_password: c.PostForm("密码确认"),
	}

	result, err := valid.ValidateStruct(data)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(result)
	c.String(http.StatusOK, "done")
}
