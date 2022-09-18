package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"slq.me/obj"
	"slq.me/service"
	// "slq.me/controller"		//同一个包下可以直接调用其他go文件的public方法
	// "fmt"
)

func SetUserRouter(r *gin.Engine) {
	g := r.Group("/u")
	g.Use(LoginMidware)
	{
		g.GET("page", page)
		g.GET("oneByid", oneByid)
		g.GET("delete", delete)
		g.POST("update", update)
		g.POST("insert", insert)

		g.GET("users", users)
	}
}

func users(c *gin.Context) {
	c.HTML(http.StatusOK, "users.html", nil)
}

func page(c *gin.Context) {
	var p obj.Pagination
	c.ShouldBind(&p)
	// log.Println("page before",p)
	p.PageSize=3
	if p.PageNum==0{
		p.PageNum=1
	}
	keyword:=c.Query("keyword")
	log.Println("page after",p,"keyword",keyword)

	res := service.PageUser(p, keyword)
	c.JSON(http.StatusOK, res)
}

func oneByid(c *gin.Context) {
	id := c.Query("id")
	res := service.OneUser(id)
	c.JSON(http.StatusOK, res)
}

func delete(c *gin.Context) {
	id := c.Query("id")
	res := service.DeleteUser(id)
	c.JSON(http.StatusOK, res)
}

func update(c *gin.Context) {
	var o obj.User
	c.ShouldBind(&o)
	// c.BindJSON(&o)
	log.Println("update", o)
	res := service.UpdateUser(o)
	c.JSON(http.StatusOK, res)
}

func insert(c *gin.Context) {
	var o obj.User
	// c.Bind(o)
	c.ShouldBind(&o)
	res := service.InsertUser(o)
	c.JSON(http.StatusOK, res)
}
