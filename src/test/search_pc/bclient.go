package main

import (
	"net"
	"fmt"
	"strconv"
	"os"
	"time"
)


var  peer_port int = 65535
var  local_port int = 65530
func get_internal()(str string) {
	var ip_string string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops:" + err.Error())
		os.Exit(1)
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				//os.Stdout.WriteString(ipnet.IP.String() + "\n")
				ip_string = ipnet.IP.String()
			}
		}
	}
	return ip_string
}

func Ask(){

	rAddr,err:=net.ResolveUDPAddr("udp4","255.255.255.255:65535")
	if(nil !=err){
		fmt.Println(err)
		fmt.Println("dd")
		return
	}

	conn1,err:=net.DialUDP("udp",nil,rAddr)
	if(nil !=err){
		fmt.Println(err)
		fmt.Println("ss")
		return
	}
	defer conn1.Close()

	_,err=conn1.Write([]byte(strconv.Itoa(local_port)))
	if(nil !=err){
		fmt.Println(err)
		fmt.Println("aa")
		return
	}



	udpaddr,err:=net.ResolveUDPAddr("udp4",":"+strconv.Itoa(local_port))
	if(nil !=err){
		fmt.Println(err)
		return
	}
	conn2,err:=net.ListenUDP("udp",udpaddr)
	if(nil !=err){
		fmt.Println(err)
		return
	}
	defer conn2.Close()
	var buf [1024]byte
	t1:=time.Now()
	for i:=0;i<5;i++{
		conn2.SetReadDeadline(time.Now().Add( time.Second))
		//time.Now().Add(time.Duration(10) * time.Second)
		n,raddr,err:=conn2.ReadFromUDP(buf[0:])
		if(err !=nil){
			//fmt.Println(err)
			//continue
			break
		}
		fmt.Println(time.Now().Sub(t1).String()+":" + string(buf[0:n])+" from "+ raddr.IP.String(),strconv.Itoa(raddr.Port))
	}


}



func main()  {

		for i:=0;i<1;i++{

			Ask()
		}


}