package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

func run_redis(){


	c,e:=redis.Dial("tcp","qycam.com:50202")
	if e!=nil{
		fmt.Println(e)
		return
	}
	defer c.Close()

	_,err0:=c.Do("AUTH","123456789")
	if err0!=nil {
		fmt.Println(err0)
	}
	//fmt.Println("pass:",v0)



	_,err:=c.Do("set","color1","red")
	if err!=nil {
		fmt.Println(err)
	}
	//fmt.Println("set :",v)

	if err!=nil {
		fmt.Println(err)
	}
	_,err=redis.String(c.Do("get","color1"))
//fmt.Println("get value:",v)



}

func main(){
	t:=time.Now()
	run_redis()
	fmt.Println(time.Since(t))
}