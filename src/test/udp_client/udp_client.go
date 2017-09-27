package main

import (
	"net"
	"fmt"
	"time"
	"strconv"
)

func Ask(i int){

	rAddr,err:=net.ResolveUDPAddr("udp4","10.10.3.226:60000")
	if(nil !=err){
		fmt.Println(err)
		fmt.Println("dd")
		return
	}

	conn,err:=net.DialUDP("udp",nil,rAddr)
	if(nil !=err){
		fmt.Println(err)
		fmt.Println("ss")
		return
	}
	defer conn.Close()

	_,err=conn.Write([]byte(strconv.Itoa(i)))
	if(nil !=err){
		fmt.Println(err)
		fmt.Println("aa")
		return
	}
	var buf [1024]byte
	_,_,err=conn.ReadFromUDP(buf[0:])
	if(nil !=err){
		fmt.Println(err)
		//fmt.Println("uu")
		return
	}
	fmt.Println("received response: " + string(buf[0:n]))
	fmt.Println(rAddr.IP,rAddr.Port)

}

func main()  {
	t1:=time.Now()
	fmt.Println(t1)
	for i:=0;i<1000;i++{
		time.After(10000)
		Ask(i)
	}
	fmt.Println(time.Now().Sub(t1))
}