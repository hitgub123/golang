package controller

import (
	"net/http"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func GetBaseRouter() *gin.Engine {
	r := gin.Default()
	//修改{{xx}}成{[xx]}，否则与vue冲突
	r.Delims("{[","]}")
	// 创建基于cookie的存储引擎，secret1 参数是用于加密的密钥
	store := cookie.NewStore([]byte("secret1"))
	// 设置session中间件，参数go-session，指的是session的名字，也是cookie的名字
	// store是前面创建的存储引擎，我们可以替换成其他存储引擎
	r.Use(sessions.Sessions("go-session", store))

	r.LoadHTMLGlob("./template/**/*")
	r.Static("/static", "./static")

	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	SetLoginRouter(r)
	return r
}
