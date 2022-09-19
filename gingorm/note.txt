项目介绍：
    代码地址：https://github.com/hitgub123/golang/tree/master/gingorm
    使用技术：
        go(gin gorm)，postgresql
        vue(无脚手架)，axios
        gin-contrib/sessions
        gin-contrib/sessions/cookie
    大部分功能前后端分离，mvc架构，实现了增删改查，数据校验，多文件上传

使用"github.com/gin-contrib/sessions"和"github.com/gin-contrib/sessions/cookie"，
Set，Delete等方法修改session后必须session.Save()才生效

结构体字段是int时，前端ajax post请求传参也要是int，传string会无法接收

gorm的事务处理：
    func TestTX() {
        tx := GetDB().Begin()
        defer func() {
            if err := recover(); err != nil {
                tx.Rollback()
            } else {
                tx.Commit()
            }
        }()
        tx.Create(&obj.User{Name: "Giraffe"})
        zero := 0
        log.Println("InsertUser>>", 1/zero)
        tx.Create(&obj.User{Name: "Lion"})

    }

gorm打印sql文
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:logger.Default.LogMode(logger.Info),	
	})
    

axios（多个或者单个）上传文件，代码见static\html\user.html
    html部分
        文件1<input type='file' name='file11' ref='file1'/><br />
        文件2<input type='file' name='file22' ref='file2'/><br />
    js部分，这里用vue的refs取的文件，也可以用其他方法
          let params = new FormData();
          //append到同一个key，会自动创建数组。这里设置什么key，后台接收的map就用什么key
          params.append('file666', this.$refs.file1.files[0])        
          params.append('file666', this.$refs.file1.files[0])
          const config = {
            headers: { "Content-Type": "multipart/form-data" }
          }
          axios.post("/u/insert", params,config)
    后台部分，使用go
        f, err = c.MultipartForm()
        if err != nil {
            goto End
        } else {         
            files:= f.File
            //files：map[file666:[0xc00007e230 0xc00007e050]]，可以直接遍历files['file666']来取得单个文件
            //这里key是html文件里file的name属性（即file666）,files_[]*FileHeader（FileHeader切片）
            for key, files_ := range files {
                for index,f :=range(files_){
                    // f := files_[0]       //单文件上传时取第一个文件即可
                    //path.Ext获取文件后缀名，好像前方自带小数点
                    dest := fmt.Sprintf("static/upload/%d-%s%s", time.Now().Nanosecond(), key, path.Ext(f.Filename))
                    if err = c.SaveUploadedFile(f, dest); err != nil {
                        goto End
                    }
                }
            }
        }

vue拼接字符串和变量：<img style='height:100px' :src='"/"+value.pic' />

为了方便学习比对，user的insert方法提交了文件，update方法没提交文件。
代码见 static\html\user.html的js部分 和 controller\userController.go的update和insert方法

文件上传的multipart/form-data表单，shouldbind似乎无法绑定参数，可以用PostForm绑定

同一个包下可以直接调用其他go文件的public方法

在结构体中对field添加binding属性，
    结构体可以写在dao层等其他地方，
    使用Query或者PostForm不会进行校验，
    如果提交的表单是MultipartForm，不会进行校验，
    使用ShouldBind/Bind进行检验。如果产生error，必须手动处理。否则会继续无视校验规则绑定成功，代码如下：
	if err:= c.ShouldBind(&o); err != nil {
		log.Println("err",err)
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} 
    可以把参数放请求头里，用ShouldBindHeader获取，
        比如headers: { "Content-Type": "multipart/form-data","name":this.user.name }
	放汉字报错String contains non ISO-8859-1 code point，需要转码

重定向遇到的method的bug：
    访问项目时，使用中间件查看session里有没有user字段，
    如果没有就重定向到login，login的get/post和logout没有使用中间件。
    代码是c.Redirect(http.StatusPermanentRedirect , "/login")。
    发现get方法访问就get到此页面(正常)，post访问就post到此页面(异常)，
    因为post用来提交登录信息。需要全都设置为get方法。
    一开始修改c.Request.Method="GET"，发现method是改了，依旧是POST请求。
    解决方法：
        重定向代码改成c.Redirect(http.StatusMovedPermanently , "/login")
        参考https://gin-gonic.com/zh-cn/docs/examples/redirects/
        和http://psychedelicnekopunch.com/archives/1848
    推荐使用301(StatusMovedPermanently)或302(StatusFound)

    

    