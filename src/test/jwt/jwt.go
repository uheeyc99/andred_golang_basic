package main

import (
	"fmt"
	"time"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"html/template"
	"encoding/json"
)

var mySignKey= "Andrew"


type AndrewToken struct {
	AndrewToken string `json:"TokenAndrew"`
}

func AndrewMakeToken(audience string)(string,error){
	token:=jwt.New(jwt.SigningMethodHS256)
	claims:=jwt.StandardClaims{
		Audience:audience,
		ExpiresAt:int64(time.Now().Unix()+1000),
		Id:"id",
		IssuedAt:int64(time.Now().Unix()),
		Issuer:"IamServer",
		NotBefore:int64(time.Now().Unix()-1000),
		Subject:"admin",
	}
	token.Claims=claims
	fmt.Println("claims",token.Claims,token.Header,token.Signature)
	SignedToken,err:=token.SignedString([]byte(mySignKey))
	if err!=nil{
		fmt.Println(err)
		return "",err
	}
	return SignedToken,nil
}

func AndrewParseToken(token string)(string,error){
	t,err:=jwt.Parse(token,func(*jwt.Token)(interface{},error){
		return []byte(mySignKey),nil
	})
	if err !=nil {
		fmt.Println("parase with claims failed:",err)
		return "",err
	}

	c,_:=t.Claims.(jwt.MapClaims)
	fmt.Println(c["aud"],c["jti"])
	return c["jti"].(string),nil
}

func loginHandler(w http.ResponseWriter, r *http.Request){

	fmt.Println("login_check method",r.Method)
	fmt.Println("Header",r.Header)


	if r.Method == "GET"{
		t,_:=template.ParseFiles("html/login.html")
		t.Execute(w,nil)
	}

	if r.Method == "POST"{

		r.ParseForm()
		w.Header().Set("Content-Type", "application/json")
		username:=r.Form.Get("username")
		passwd:=r.Form.Get("password")
		fmt.Println(username,passwd)
		reason:=checkUserByName(username,passwd)
		if reason==""{
			tokenString,_:=AndrewMakeToken(username)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(tokenString))
			return
		}

		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(reason))
		return
	}

}

func checkUserByName(user string,passwd string)(string){
	if(user != "eric"){
		return "Not Found User"
	}
	if(passwd !="12345678"){
		return "passwd error"
	}
	return ""
}

func resourceHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("resource method",r.Method)
	fmt.Println("Header",r.Header)
	r.ParseForm()
	s,err:=AndrewParseToken(r.Header.Get("Authorization"))
	if nil != err{
		w.Write([]byte(err.Error()))
	}
	w.Write([]byte("hello " + s))

}

func AndrewJsonResponse(response interface{}, w http.ResponseWriter) {

	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(string(json))
	w.Write(json)
}

func start_server(){

	http.HandleFunc("/login",loginHandler)
	http.HandleFunc("/resource",resourceHandler)
	http.ListenAndServe(":9090", nil)
}

func main(){
	start_server()
}