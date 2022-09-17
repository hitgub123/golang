package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
)

func main() {
	mainGin()
}

func mainGin() {
	r := gin.Default()

	//为静态文件做映射，通过 项目url地址/cssjs 访问./static下的静态文件
	r.Static("/cssjs","./static")

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
		"./template/tmpl基本语法.go.tmpl",
		"./template/test1/block.tmpl",
		"./template/test1/useblock.tmpl",
		"./template/test1/js.html")

	//处理url请求
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
		c.JSON(http.StatusOK,gin.H{
			"name":"kitty",
			"age":3,
		})
	})

	//返回json方法2，key首字母必须大写
	r.GET("/j2", func(c *gin.Context) {
		// 首字母大写才能显示在页面上，
		type user struct{
			//`json:"usernane"`指定json返回的key 是 username
			Name string	`json:"usernane"`
			//不显示首字母小写的字段
			age int8
			Weight float32
		}
		u:=user{"kitty2",4,1.23}
		//页面显示{"usernane":"kitty2","Weight":1.23}
		c.JSON(http.StatusOK,u)
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

type UserInfo struct {
	Name   string
	Gender bool
	Age    int
}
