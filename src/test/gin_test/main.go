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
curl http://127.0.0.1:8006/simple/server/get?ggg=5
curl -d "" http://127.0.0.1:8006/simple/server/post
curl --request PUT http://127.0.0.1:8006/simple/server/put
curl --request DELETE http://127.0.0.1:8006/simple/server/delete
http://blog.csdn.net/moxiaomomo/article/details/51153779

 */


func test(){

	gin.SetMode(gin.DebugMode)
	router:=gin.Default()

	router.Use(Middleware)
	router.LoadHTMLGlob("templates/*")
	router.POST("/upload", PostFileHandler)
	router.POST("/uploads", PostFilesHandler)
	router.POST("/simple/server/form_post", PostHandler2)

	router.GET("/simple/server/get",GetHandler)
	router.POST("/simple/server/post", PostHandler)
	router.PUT("/simple/server/put", PutHandler)
	router.DELETE("/simple/server/delete", DeleteHandler)
	router.Any("/test",TestHandler)
	//监听端口
	router.Run(":9090")

}
func Middleware(c *gin.Context) {
	fmt.Println("this is a middleware!")

}
func TestHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html",gin.H{
		"title": "Users",
	})

	return
}

func GetHandler(c *gin.Context) {
	value, exist := c.GetQuery("key")
	if !exist {
		value = "the key is not exist!"
	}
	fmt.Println(c.Request.Header)
	c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("get success! %s\n", value)))
	return
}

func PostHandler2(c *gin.Context) {

		user := c.PostForm("username")
		pass := c.DefaultPostForm("password", "999999")
		//pass := c.PostForm("password")
		c.JSON(http.StatusOK, gin.H{
			"passwd": pass,
			"username":    user,
		})

}

func PostHandler(c *gin.Context) {
	type JsonHolder struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	holder := JsonHolder{Id: 1, Name: "my name"}
	//若返回json数据，可以直接使用gin封装好的JSON方法
	c.JSON(http.StatusOK, holder)
	return
}

func PostFileHandler(c *gin.Context) {
	//curl -X POST http://127.0.0.1:9090/upload
	// -F "andrewfile=@/Users/eric/Desktop/2.jpg" -H "Content-Type: multipart/form-data"

	file,fileheader,err:=c.Request.FormFile("andrewfile")
	if err !=nil{
		fmt.Println("err1:",err)
		c.String(http.StatusBadRequest,"Bad resuest!")
		return
	}
	filename:=fileheader.Filename
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
	c.Data(http.StatusOK, "text/plain", []byte("put success!\n"))
	return
}
func DeleteHandler(c *gin.Context) {
	c.Data(http.StatusOK, "text/plain", []byte("delete success!\n"))
	return
}

func main(){

	test()

}
