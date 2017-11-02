package main

import (
	"net/http"
	"fmt"
	"net/url"
	"strings"
	"compress/gzip"
	"io"
	"io/ioutil"
	"encoding/json"
)

var auth_token string = "my token"

func andrew_get(url string){
	resp,err:=http.Get("http://127.0.0.1/andrew/hello")
	if err!=nil{
		fmt.Println("error:",err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header)
	var p [10240]byte
	n,err:=resp.Body.Read(p[0:])
	fmt.Println(string(p[0:n]))
}



func andrew_portform(addr string){

	resp,err:=http.PostForm(addr,
		url.Values{"username":{"eric123"},"password":{"12345678"}})
	if err!=nil{
		fmt.Println("error:",err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Status:",resp.StatusCode)
	fmt.Println("Header:",resp.Header)
	var p [228]byte
	n,err:=resp.Body.Read(p[0:])
	auth_token = string(p[0:n])

	fmt.Println("auth_token:",auth_token,n)

	type Sta struct {
		Status string `json:"status_code"`
		Status_code string `json:"status"`
	}
	type pp struct{
		Username string `json:"username"`
		Pass string `json:"passwd"`
		AA Sta
	}
	var a pp
	json.Unmarshal(p[0:n],&a)
	fmt.Println(a.Username,a.Pass)
}


func andrew_post(url string){
	resp,err:=http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader("username=eric"))

	if err!=nil{
		fmt.Println("error:",err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header)
	var p [10240]byte
	n,err:=resp.Body.Read(p[0:])
	fmt.Println(string(p[0:n]))





}


func test_get(url string){

	client:= &http.Client{}
	request,err:=http.NewRequest("GET",url,nil)
	//request.Header.Add("Authorization","Bearer "+ auth_token)
	request.Header.Add("Authorization",auth_token)
	request.Header.Add("Content-Type","application/x-www-form-urlencoded")


	response,err:=client.Do(request)
	if err!=nil{
		fmt.Println("response:",err)
		return
	}
	defer response.Body.Close()

	fmt.Println("status:",response.StatusCode)
	fmt.Println("header:",response.Header)

	if response.StatusCode == 200{
		var body string
		switch response.Header.Get("Content-Type") {
			case "gzip":
				reader, _ := gzip.NewReader(response.Body)
				for {
					buf := make([]byte, 1024)
					n, err := reader.Read(buf)
					if err != nil && err != io.EOF {
						panic(err)
					}
					if n == 0 {
						break
					}
					body += string(buf)
				}
			default:
				bodyByte, _ := ioutil.ReadAll(response.Body)
				body = string(bodyByte)

		}


		fmt.Println("body:",body)

	}








}
func main(){
	//andrew_portform()
	//test_get()
	//test_get("http://127.0.0.1:9090/test")
	//test_get("http://127.0.0.1:9090/index")
	//andrew_portform("http://127.0.0.1:9090/login_check")
	andrew_portform("http://127.0.0.1:9090/postform")

	//test_get("http://127.0.0.1/andrew/ssss")
	//test_get("http://127.0.0.1:9090/ssss")
	//test_get("http://127.0.0.1/baidu")
	//andrew_get("http://127.0.0.1/go/hello")
}
