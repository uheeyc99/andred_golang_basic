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
	"bytes"
)

type andrew_token struct {
	Token string
}

var auth_token string = "my token"

type login_info struct {
	Username string
	Password string
}



func andrew_get(){
	resp,err:=http.Get("http://qycam.com:50219/info2.php")
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

func andrew_portform1(){

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

	fmt.Println("auth_token:",auth_token,n)
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





func getToken2() {
	user:=login_info{
		Username:"eric",
		Password:"12345678",
	}
	j,_:=json.Marshal(user)

	req:= bytes.NewBuffer(j)

	client := &http.Client{}
	request, err := http.NewRequest("POST", "http://127.0.0.1:9090/login", req)
	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)

	if err != nil {
		fmt.Println("response:", err)
		return
	}
	defer response.Body.Close()

	fmt.Println("status:", response.StatusCode)
	fmt.Println("header:", response.Header)

	if response.StatusCode == 200 {
		var bs string
		ct :=response.Header.Get("Content-Type")

		if strings.Contains(ct,"gzip") {
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
				bs += string(buf)
			}
			auth_token = bs
		}



		if strings.Contains(ct,"application/json"){
			var t andrew_token
			bodyByte, _ := ioutil.ReadAll(response.Body)
			json.Unmarshal(bodyByte,&t)
			auth_token= t.Token
		}

		if  strings.Contains(ct,"text/plain"){
			bodyByte, _ := ioutil.ReadAll(response.Body)
			auth_token = string(bodyByte)
		}


	}

}

func getToken1(){

	response,err:=http.PostForm("http://127.0.0.1:9090/login",
		url.Values{"username":{"eric"},"password":{"12345678"}})
	if err!=nil{
		fmt.Println("error:",err)
		return
	}
	defer response.Body.Close()
	fmt.Println("Status:",response.StatusCode)
	fmt.Println("Header:",response.Header)
	if response.StatusCode == 200 {
		var p [1024]byte
		n,_:=response.Body.Read(p[0:])
		ct :=response.Header.Get("Content-Type")

		if strings.Contains(ct,"application/json"){
			var t andrew_token
			json.Unmarshal(p[0:n],&t)
			auth_token= t.Token
		}

		if  strings.Contains(ct,"text/plain"){
			auth_token = string(p[0:n])
		}


	}

}


func main(){
	andrew_get()
	getToken1()
	fmt.Println("auth_Token:",auth_token)
	getToken2()
	fmt.Println("auth_Token:",auth_token)


}
