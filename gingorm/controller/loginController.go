package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"slq.me/obj"
	"slq.me/service"
	"reflect"
	// "fmt"
)

func SetLoginRouter(r *gin.Engine) {

	r.GET("login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.GET("logout", func(c *gin.Context) {
		session := sessions.Default(c)
		// name:=session.Get("name")
		// log.Println("session before name =", name)
		session.Delete("name")
		session.Save()
		// name=session.Get("name")
		// log.Println("session after name =", name)
		c.JSON(http.StatusOK, nil)
	})

	r.POST("login", func(c *gin.Context) {
		var u obj.User
		c.ShouldBind(&u)
		res := service.OneUser2(u.Name, u.Password)
		//reflect.TypeOf( res.Data)=**obj.User
		log.Println("Ok res1 =", res.Data,reflect.TypeOf( res.Data))
		//修改OneUser2后，只需判断res.Data是不是nil即可
		//类型转换
		//value=reflect.TypeOf( res.Data)
		if value, ok := res.Data.(**obj.User); ok {
			// value.Id报错，reflect.TypeOf(* value)=*obj.User
			log.Println("Ok value =", value,reflect.TypeOf( value),(*value).Id,(*value).Name )
			if (*value).Id != 0 {
				session := sessions.Default(c)
				session.Set("name", (*value).Name)
				session.Save()
			}
		}else{
			log.Println("Ok ok =", ok)
		}
		c.JSON(http.StatusOK, res)
	})
}

func LoginMidware(c *gin.Context) {
	log.Println("LoginMidware检查是否登录")
	// 初始化session对象
	session := sessions.Default(c)
	// 通过session.Get读取session值
	// session是键值对格式数据，因此需要通过key查询数据
	name := session.Get("name")
	log.Println("session name =", name)
	if name == nil {
		// 设置session数据
		// session.Set("hello", "world")
		// 删除session数据
		// session.Delete("tizi365")
		// 保存session数据
		// session.Save()
		// 删除整个session
		// session.Clear()
		c.Redirect(http.StatusPermanentRedirect, "/login")
		c.Abort()
	}

}
