package controller

import (
	"log"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"slq.me/obj"
	"slq.me/service"

	// "slq.me/controller"		//同一个包下可以直接调用其他go文件的public方法
	"fmt"
	"path"
	"time"
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
		g.GET("testtx", testTX)

		g.GET("users", users)
	}
}

func users(c *gin.Context) {
	c.HTML(http.StatusOK, "users.html", nil)
}

func page(c *gin.Context) {
	var p obj.Pagination
	if err:= c.ShouldBind(&p); err != nil {
		log.Println("err",err)
		c.JSON(http.StatusOK, obj.Result{Code: 1,ErrorMessage: err.Error(),Data: nil})
		return
	} 
	// log.Println("page before",p)
	p.PageSize = 3
	if p.PageNum == 0 {
		p.PageNum = 1
	}
	keyword := c.Query("keyword")
	log.Println("page after", p, "keyword", keyword)

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
	//这里如果ShouldBind产生error，必须手动处理。否则会继续无视校验规则绑定成功
	if err:= c.ShouldBind(&o); err != nil {
		log.Println("err",err)
		c.JSON(http.StatusOK, obj.Result{Code: 1,ErrorMessage: err.Error(),Data: nil})
		return
	} 
	log.Println("update", o)
	res := service.UpdateUser(o)
	c.JSON(http.StatusOK, res)
}

func insert(c *gin.Context) {
	var (
		f   *multipart.Form
		err error
		res *obj.Result
		o obj.User
	)
	//ShouldBind绑定不了，可以用PostForm
	// c.ShouldBind(&o)
	// log.Println("111",o)
	// c.ShouldBindJSON(&o)
	// log.Println("112",o)
	// c.ShouldBindQuery(&o)
	// log.Println("113",o)
	// c.ShouldBindTOML(&o)
	// log.Println("114",o)
	// c.ShouldBindXML(&o)
	// log.Println("115",o)
	// c.ShouldBindYAML(&o)
	// log.Println("116",o)
	// c.ShouldBindYAML(&o)
	// log.Println("117",o)
	// 可以把参数放请求头里，用ShouldBindHeader获取
	// 放汉字报错String contains non ISO-8859-1 code point，需要转码
	// 比如headers: { "Content-Type": "multipart/form-data","name":this.user.name }
	// c.ShouldBindHeader(&o)
	// log.Println("118",o)

	
	o.Id,_= strconv.Atoi(c.PostForm("id"))
	o.Name=c.PostForm("name")
	o.Age,_= strconv.Atoi(c.PostForm("age"))
	o.Password=c.PostForm("password")
	
	// 上传多个文件或单个文件
	f, err = c.MultipartForm()
	//fmt.Printf("\nf:%v", f)
	if err != nil {
		log.Println("11111111111",err)
		goto End
	} else {
		
		files:= f.File

		//files：map[file666:[0xc00007e230 0xc00007e050]]，可以直接遍历files['file666']来取得单个文件
		fmt.Printf("\nfiles:%v,%d", files,len(files))
		//这里key是html文件里file的name属性,files_是[]*FileHeader（FileHeader切片）
		for key, files_ := range files {
			for index,f :=range(files_){
				// f := files_[0]				//单文件上传时取第一个文件即可
				fmt.Printf("\n key:%s:%s", key,f.Filename)
				//path.Ext获取文件后缀名，好像前方自带小数点
				dest := fmt.Sprintf("static/upload/%d-%s%s", time.Now().Nanosecond(), key, path.Ext(f.Filename))
				if index==0{
					o.Pic1=dest
				}
				if index==1{
					o.Pic2=dest
				}				
				//fmt.Println(dest)
				if err = c.SaveUploadedFile(f, dest); err != nil {
					log.Println("11111122222",err)
					goto End
				}
			}

		}
	}

	// c.Bind(o)

	res = service.InsertUser(o)
	c.JSON(http.StatusOK, res)
	return
End:
	c.JSON(http.StatusOK, obj.Result{Code: 1, ErrorMessage:"文件上传失败", Data: nil})
}

func testTX(c *gin.Context) {
	service.TestTX()
}
