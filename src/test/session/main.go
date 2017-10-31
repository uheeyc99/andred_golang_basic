package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("value:", strings.Join(v, ""))
	}
	fmt.Fprintln(w, "hello nihao")

}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("src/html/login.gtpl")
		if err != nil {
			fmt.Println(err)
			return
		}
		t.Execute(w, nil)

	} else {
		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])
	}
}
func login2(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("src/html/login.gtpl")
		//w.Header().Set("Content-Type", "text/html")
		t.Execute(w, sess.Get("username"))
	} else {
		sess.Set("username", r.Form["username"])
		http.Redirect(w, r, "/", 302)
	}
}

func count(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	createtime := sess.Get("createtime")
	if createtime == nil {
		sess.Set("createtime", time.Now().Unix())
	} else if (createtime.(int64) + 360) < (time.Now().Unix()) {
		globalSessions.SessionDestroy(w, r)
		sess = globalSessions.SessionStart(w, r)
	}
	ct := sess.Get("countnum")
	if ct == nil {
		sess.Set("countnum", 1)
	} else {
		sess.Set("countnum", (ct.(int) + 1))
	}
	t, _ := template.ParseFiles("count.gtpl")
	w.Header().Set("Content-Type", "text/html")
	t.Execute(w, sess.Get("countnum"))
}
func main() {
	//http.HandleFunc("/", sayHelloName);
	http.HandleFunc("/login", login)
	http.HandleFunc("/login2", login2)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatalf("Listen and server", err)
	}
}
var globalSessions *Manager
func init() {
	globalSessions,_= NewSessionManager("memory", "goSessionid", 3600)
	//go globalSessions.GC()
	fmt.Println("fd")
}
