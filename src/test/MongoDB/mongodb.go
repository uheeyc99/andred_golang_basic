package main

import (
	"gopkg.in/mgo.v2"
	"fmt"
)

const URL  = "qycam.com:50203"

func test(){
	session,err:=mgo.Dial(URL)
	if err!=nil{
		fmt.Println(err)
	}
	defer session.Close()
}


func main(){

	test()

}
