package main

import (
	"net/http"
	"fmt"
	"net/url"
	"strings"
	"compress/gzip"
	"io"
	"io/ioutil"
)

var auth_token string = "my token"

func andrew_get(){
	resp,err:=http.Get("http://127.0.0.1:9090/login")
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

func andrew_portform(){

	resp,err:=http.PostForm("http://127.0.0.1:9090/login",
		url.Values{"username":{"eric"},"password":{"12345678"}})
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
	fmt.Println("auth_token:",auth_token)
}
func andrew_post(){
	resp,err:=http.Post("http://127.0.0.1:9090/login",
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


func test_get(){

	client:= &http.Client{}
	request,err:=http.NewRequest("GET","http://127.0.0.1:9090/resource",nil)
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
	andrew_portform()
	test_get()
}
