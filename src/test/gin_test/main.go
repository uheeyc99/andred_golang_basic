package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"os"
	"log"
	"io"
)

/*

http://blog.csdn.net/moxiaomomo/article/details/51153779
https://studygolang.com/articles/2488
http://blog.csdn.net/lijunwyf/article/details/50238005
 */


func test(){

	gin.SetMode(gin.DebugMode)
	router:=gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/assets","./assets")
	router.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK,"pong")
	})


	v0:=router.Group("/v0")
	v1:=router.Group("/v1")
	v2:=router.Group("/v2")

	{


		v0.GET("/get",Middleware,GetHandler)
		v0.POST("/post", PostHandler)
		v0.POST("/postform", PostFormHandler)

		v0.PUT("/put/:id", PutHandler)
		v0.DELETE("/delete/:id/:video", DeleteHandler)

		v0.GET("/upload",Middleware,TestHandler)
		v0.POST("/upload", PostFileHandler)
		v0.POST("/uploads", PostFilesHandler)
	}



	v1.Use(Middleware)
	{
		v1.GET("/login",LoginHTMLHandler)
		v1.POST("/login",LoginPostHandler)
		v1.GET("/hello", func(c *gin.Context) {
			c.String(http.StatusOK,"hello,I am v1")
		})
	}

	{
		v2.GET("/hello", Middleware,func(c *gin.Context) {
			c.String(http.StatusOK,"hello,I am v2")
			//c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
		})
	}




	router.Run(":9090")

}
func Middleware(c *gin.Context) {
	//中间件最大的作用，莫过于用于一些记录log，错误handler，还有就是对部分接口的鉴权
	fmt.Println("this is a middleware!")
	c.Request.Header.Add("userid","9999")
	cCp:=c.Copy()
	fmt.Println(cCp.Request.URL.Path)
	//c.HTML(http.StatusOK, "upload.html",gin.H{
	//	"title": "Users",
	//})
	//c.Abort()

}
func CheckUser(c *gin.Context) {
	//中间件最大的作用，莫过于用于一些记录log，错误handler，还有就是对部分接口的鉴权
	//check token ,check cookie等操作

}
func TestHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html",gin.H{
		"title": "Users",
	})
	return
}

func LoginHTMLHandler(c *gin.Context) {

	cs,err:=c.Cookie("andrewLoginCookie")
	c.Request.Cookies()
	if err!=nil{
		fmt.Println("get cookie:",err)
		c.HTML(http.StatusOK, "login.html",gin.H{
			"title": "Users",
		})
		return
	} else {
		c.String(http.StatusOK,"you already loged in")
	}
	fmt.Println(cs)

	return
}
func LoginPostHandler(c *gin.Context) {
	user:=c.PostForm("username")
	//pass:=c.PostForm("password")
	fmt.Println(user)
	cookie:=&http.Cookie{
		Name:"andrewLoginCookie",
		Value:"userid"+user,
		Path:"/",
		HttpOnly:true,
	}
	http.SetCookie(c.Writer,cookie)
	c.String(http.StatusOK,"Login successful")

	return
}

func GetHandler(c *gin.Context) {
	//curl http://127.0.0.1:9090/get?key=5
	fmt.Println(c.Request.Header)
	value, exist := c.GetQuery("key")
	if !exist {
		value = "the key is not exist!"
	}
	fmt.Println(c.Request.Header)
	c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("get success! %s\n", value)))

	return
}



func PostHandler(c *gin.Context) {

	type JsonHolder struct {
		Username   string    `json:"username"`
		Password string `json:"password"`
	}
	holder := JsonHolder{Username: "eric", Password: "999999"}
	//若返回json数据，可以直接使用gin封装好的JSON方法
	c.JSON(http.StatusOK, holder)
	return
}


func PostFormHandler(c *gin.Context) {

	fmt.Println(c.Request.Header.Get("Content-Length"),c.Request.Header.Get("Content-Type"))
	user := c.PostForm("username")
	pass := c.DefaultPostForm("password", "999999") //如果没有则选择默认值

	fmt.Println(user,pass)

	c.JSON(http.StatusOK, gin.H{
		"passwd": pass,
		"username":    user,
	})

}



func PostFileHandler(c *gin.Context) {
	//curl -X POST http://127.0.0.1:9090/upload
	// -F "andrewfile=@/Users/eric/Desktop/2.jpg" -H "Content-Type: multipart/form-data"

	file,fileHeader,err:=c.Request.FormFile("andrewfile")
	if err !=nil{
		fmt.Println("err1:",err)
		c.String(http.StatusBadRequest,"Bad resuest!")
		return
	}
	filename:=fileHeader.Filename
	fmt.Println("filename:",filename)

	out,err:=os.Create("upload"+"/"+filename)
	if err != nil{
		log.Fatal(err)
	}
	defer out.Close()
	_,err = io.Copy(out,file)
	if err != nil{
		log.Fatal(err)
	}
	c.String(http.StatusCreated,"upload successful")

	return
}

func PostFilesHandler(c *gin.Context) {
	//curl -X POST http://127.0.0.1:9090/uploads
	// -F "andrewfile=@/Users/eric/Desktop/2.jpg" -F "andrewfile=@/Users/eric/Desktop/2.pdf"
	// -H "Content-Type: multipart/form-data"

	err:=c.Request.ParseMultipartForm(2000)
	if err !=nil{
		fmt.Println("err1:",err)
		c.String(http.StatusBadRequest,"Bad resuest!")
		return
	}

	formdata:=c.Request.MultipartForm
	files:=formdata.File["andrewfile"]
	for i,_:=range files{
		f,e:=files[i].Open()
		if e!=nil{
			log.Fatal(e)
		}
		defer f.Close()

		out,e:=os.Create("upload/"+files[i].Filename)
		defer out.Close()
		fmt.Println("copying:",out.Name())
		_,e=io.Copy(out,f)

	}
	c.String(http.StatusCreated,"upload successful !!")
	return
}


func PutHandler(c *gin.Context) {
	id:=c.Param("id")
	fmt.Println(id)
	c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("put %s success!\n",id)))
	return
}
func DeleteHandler(c *gin.Context) {
	//http://127.0.0.1:9090/simple/server/id/video?time=3333
	id:=c.Param("id")
	video:=c.Param("video")
	t,_:=c.GetQuery("time")
	fmt.Println(id,video,t)
	fmt.Println(c.Request.URL.Path)
	c.Data(http.StatusOK, "text/plain", []byte("delete success!\n"))
	return
}

func main(){

	test()

}
