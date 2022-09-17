package main

import (
	"fmt"
	"html/template"
	"log"
	"mime/multipart"
	"net/http"
	"path"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	mainGin()
}

type UserInfo struct {
	Name   string
	Gender bool
	Age    int
}

func midware1(c *gin.Context) {
	log.Println("m1 start>>")
	c.Set("person",UserInfo{"柯南",true,6})
	start:=time.Now()
	//next调用后续的 处理函数/中间件
	c.Next()
	// c.Abort()
	// c.JSON(http.StatusOK,gin.H{"return at":"m1"})
	// return
	// c.Redirect(http.StatusPermanentRedirect,"/b2")
	log.Println("<<m1 end",time.Since(start))
}


func midware2(c *gin.Context) {
	log.Println("m2 start>>")
	userInfo,ok:=c.Get("person")
	if ok || !ok{
		fmt.Printf("%v",userInfo)
	}
	// c.Next()
	log.Println("<<m2 end")
}

func handler1(c *gin.Context) {
	log.Println("h1 start>>")
	c.JSON(http.StatusOK, gin.H{"msg": "test middleware"})
	log.Println("<<h1 end")
}
func mainGin() {
	r := gin.Default()

	//据说可以限制文件大小，不起作用
	// r.MaxMultipartMemory=1 << 20
	r.MaxMultipartMemory = 1

	//为静态文件做映射，通过 项目url地址/cssjs 访问./static下的静态文件
	r.Static("/cssjs", "./static")

	//设置自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
		"sqr": func(i int) int {
			return i * i
		},
	})

	// 读取模板文件1
	// r.LoadHTMLGlob("template/**/*")
	// r.LoadHTMLGlob("./template/**/*")
	//template/**/*似乎不读取template里的文件，template/tmpl转义.html没读到
	// 读取模板文件2
	r.LoadHTMLFiles(
		"./template/tmpl转义.html",
		"./template/404.html",
		"./template/tmpl基本语法.go.tmpl",
		"./template/test1/block.tmpl",
		"./template/test1/useblock.tmpl",
		"./template/upload.html",
		"./template/test1/js.html")


	// r.Use(midware1,midware2)

	// 访问/m1输出如下，
	// m1 start>>
	// m2 start>>
	// h1 start>>
	// <<h1 end
	// <<m2 end
	// <<m1 end
	r.GET("/m1", midware1, midware2,  handler1)
	r.GET("/m2", midware2, midware1,  handler1)

	g1 := r.Group("/g1")
	// g1.Use(midware1)
	{
		// get /g1/1
		g1.GET("/1", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "g11"})
		})
		// post /g1/2
		g1.POST("/2", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "g12"})
		})
		// get /g1/g2/1
		g2 := g1.Group("/g2")
		// g2.Use(midware2,midware2)
		{
			g2.GET("/3", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"msg": "g23"})
			})
		}
	}

	//处理url请求
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	//请求转发
	r.GET("/b5", func(c *gin.Context) {
		c.Request.URL.Path = "/b2"
		r.HandleContext(c)
	})
	//重定向
	r.GET("/b4", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "https://www.google.com")
	})
	//重定向
	r.GET("/b3", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/b2")
	})
	r.GET("/b2", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", nil)
	})
	r.Any("/b1", func(c *gin.Context) {
		type user struct {
			Name   string `binding:"required,min=2,max=4"`
			Age    int8   `form:"age" json:"age" binding:"required,gte=1,lte=9"`
			Weight float32
			Email  string `binding:"required,email"`
			Email2 string `binding:"required,eqfield=Email"`
			Gender string `binding:"required,oneof=m f"`
		}
		var (
			u   user
			f   *multipart.Form
			err error
		)
		if err = c.ShouldBind(&u); err != nil {
			goto End
		} else {
			log.Println(u)
		}

		// 上传多个文件或单个文件
		f, err = c.MultipartForm()
		fmt.Printf("\nf:%v", f)
		if err != nil {
			goto End
		} else {
			files := f.File
			fmt.Printf("\nfiles:%v", files)
			//这里key是html文件里file的name属性,file_是[]*FileHeader（FileHeader切片）
			for key, file_ := range files {
				f := file_[0]
				fmt.Printf("\n key:%s:%s", key, f.Filename)
				//path.Ext获取文件后缀名
				dest := fmt.Sprintf("./upload/%d-%s.%s", time.Now().Unix(), key, path.Ext(f.Filename))
				fmt.Println(dest)
				if err = c.SaveUploadedFile(f, dest); err != nil {
					goto End
				}
			}
		}

		// 上传单个文件
		// f, err = c.FormFile("file1")
		// if err != nil {
		// 	goto End
		// } else {
		// 	dest := path.Join("./upload", f.Filename)
		// 	fmt.Println(dest)
		// 	if err = c.SaveUploadedFile(f, dest); err != nil {
		// 		goto End
		// 	}
		// }
	End:
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"err": err.Error()})
		} else {
			c.JSON(http.StatusOK, u)
		}

	})

	r.Any("/p1", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			a1 := c.Query("a")
			a2 := c.DefaultQuery("a", "Default A")
			a3, ok := c.GetQuery("a")
			log.Println("\na1=", a1, "\na2=", a2, "\na3=", a3, "\nok=", ok)
			log.Println("\na1=", len(a1), "\na2=", len(a2), "\na3=", len(a3))
			log.Println("\na1=", a1 == "", "\na3=", a3 == "")
		case "POST":
			a1 := c.PostForm("a")
			a2 := c.DefaultPostForm("a", "Default A")
			a3, ok := c.GetPostForm("a")
			log.Println("\na1=", a1, "\na2=", a2, "\na3=", a3, "\nok=", ok)
			log.Println("\na1=", len(a1), "\na2=", len(a2), "\na3=", len(a3))
			log.Println("\na1=", a1 == "", "\na3=", a3 == "")
		default:
			log.Println("c.Request.Method=", c.Request.Method)
		}
		c.JSON(http.StatusOK, nil)
	})

	///p2//update/a  ok
	///p2/3/update//  ng
	r.GET("/p2/:id/update/:name", func(c *gin.Context) {
		a1 := c.Param("id")
		a2 := c.Param("name")

		log.Println("\na1=", a1, "\na2=", a2)
		log.Println("\na1=", len(a1), "\na2=", len(a2))
		log.Println("\na1=", a1 == "", "\na2=", a2 == "")

		c.JSON(http.StatusOK, nil)
	})

	r.GET("/a1", func(c *gin.Context) {
		// 参数1：响应码，参数2：模板名，参数3：要传递给页面的参数
		c.HTML(http.StatusOK, "useblock.tmpl", nil)
	})
	r.GET("/a2", func(c *gin.Context) {
		c.HTML(http.StatusOK, "tmpl转义.html", gin.H{
			"msg1": "<script>alert('msg1')</script>",
			"msg2": "<script>alert('msg2')</script>",
		})
	})

	//返回json方法1，key首字母可以小写
	r.GET("/j1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "kitty",
			"age":  3,
		})
	})

	//返回json方法2，key首字母必须大写
	r.GET("/j2", func(c *gin.Context) {
		// 首字母大写才能显示在页面上，
		type user struct {
			//`json:"usernane"`指定json返回的key 是 username
			Name string `json:"usernane"`
			//不显示首字母小写的字段
			age    int8
			Weight float32
		}
		u := user{"kitty2", 4, 1.23}
		//页面显示{"usernane":"kitty2","Weight":1.23}
		c.JSON(http.StatusOK, u)
	})
	r.Run(":8080")
}

func mainNoGin() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/a1", a1)
	http.HandleFunc("/a2", a2)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("HTTP server failed,err:", err)
		return
	}

}

func a2(w http.ResponseWriter, r *http.Request) {
	file := "./template/tmpl转义.html"
	tmpl, err := template.New("tmpl转义.html").
		Funcs(template.FuncMap{"safe": func(s string) template.HTML {
			return template.HTML(s)
		}}).
		ParseFiles(file)
	if err != nil {
		log.Println("create template failed, err:", err)
		return
	}
	Res := map[string]interface{}{
		"msg1": "<script>alert('msg1')</script>",
		"msg2": "<script>alert('msg2')</script>",
	}
	tmpl.Execute(w, Res)
}
func a1(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("./template/**/*.*")
	// tmpl,err:=template.ParseGlob("./template/test*/*")
	// tmpl,err:=template.ParseGlob("./template/*/*/*")
	// tmpl,err:=template.ParseGlob("./template/**/*.tmpl")
	if err != nil {
		log.Println("create template failed, err1:", err)
		return
	}
	err = tmpl.ExecuteTemplate(w, "useblock.tmpl", nil)
	// err=tmpl.Execute(w,"useblock.tmpl")
	if err != nil {
		log.Println("create template failed, err2:", err)
		return
	}
}

func mysqr(a int) (int, error) {
	return a * a, nil
}
func hello(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	file := "./template/tmpl基本语法.go.tmpl"
	jsfile := "./template/js.html"
	// tmpl, err := template.ParseFiles(file,jsfile)
	// 如果要自定义函数，使用下面的链式代码。否则使用上一行即可
	tmpl, err := template.New("tmpl基本语法.go.tmpl").Funcs(template.FuncMap{"sqr": mysqr}).ParseFiles(file, jsfile)
	// 下面这行代码不报错，但是页面不会显示任何内容。貌似New的参数只能是不带路径的文件名
	// tmpl, err := template.New(file).Funcs(template.FuncMap{"sqr": mysqr}).ParseFiles(file,jsfile)

	if err != nil {
		log.Println("create template failed, err:", err)
		return
	}
	// 利用给定数据渲染模板，并将结果写入w
	user := UserInfo{
		Name:   "枯藤",
		Gender: true,
		Age:    18,
	}
	user2 := UserInfo{
		Name:   "枯藤2",
		Gender: true,
		Age:    182,
	}
	usermap := map[string]interface{}{
		"Name":   "枯藤map",
		"Gender": false,
		"Age":    22,
	}
	arr := []UserInfo{user, user2}
	arr2 := []int{}
	Res := map[string]interface{}{
		"user":    user,
		"usermap": usermap,
		"arr":     arr,
		"arr2":    arr2,
	}
	tmpl.Execute(w, Res)
}


