package main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
	"html/template"
	"strconv"
)


func myroot(w http.ResponseWriter,r *http.Request){
	fmt.Println("hello method",r.Method)
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println(r.Header)
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	w.Header().Set("Andrew", "eric")
	fmt.Fprint(w,"hello astaie:",r.URL.Path)
}

func hello(w http.ResponseWriter,r *http.Request){
	fmt.Println("hello method",r.Method)
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v:=range r.Form  {
		fmt.Println("key:",k)
		fmt.Println("val:",strings.Join(v,""))

	}

	fmt.Fprint(w,"hello astaie")
}

func login(w http.ResponseWriter,r *http.Request){
	fmt.Println("login method",r.Method)
	r.ParseForm()
	if r.Method == "GET"{
		t,_:=template.ParseFiles("html/login.html")
		t.Execute(w,nil)
	}

}


func info(w http.ResponseWriter,r *http.Request){
	fmt.Println("login method",r.Method)
	r.ParseForm()
	if r.Method == "GET"{
		t,_:=template.ParseFiles("html/info.html")
		t.Execute(w,nil)
	}
}



func login_check(w http.ResponseWriter,r *http.Request){
	fmt.Println("login_check method",r.Method)
	r.ParseForm()
	if r.Method == "POST"{
		fmt.Println("username:",r.Form["username"])
		fmt.Println("password:",r.Form["password"])
	}
}

func info_check(w http.ResponseWriter,r *http.Request){
	fmt.Println("info_check method",r.Method)
	r.ParseForm()
	if r.Method == "POST"{

		username:=r.Form.Get("username")
		if(len(r.Form["username"][0])==0){
			fmt.Fprint(w,"用户名不能为空\n")
		}else{
			fmt.Println("账号:",username)
		}


		phone_string:=r.Form.Get("phone")
		if phone_string != ""{
			getint,err:=strconv.Atoi(phone_string)
			if err!=nil{
				fmt.Fprint(w,"电话号码格式错误")
			}else{
				fmt.Println("电话号码：",getint)
			}
		}
		
		fruit:=r.Form.Get("fruit")
		fmt.Println("水果:",fruit)




	}

	if r.Method=="GET"{
		t,_:=template.ParseFiles("html/info.html")
		t.Execute(w,nil)
	}
}


func main()  {
	http.HandleFunc("/",myroot)
	http.HandleFunc("/hello",hello)
	http.HandleFunc("/login",login)
	http.HandleFunc("/login_check",login_check)
	http.HandleFunc("/info",info)
	http.HandleFunc("/info_check",info_check)
	err:=http.ListenAndServe(":9090",nil)
	if err!=nil {
		log.Fatal("ListenAndServe: ",err)
	}
}
