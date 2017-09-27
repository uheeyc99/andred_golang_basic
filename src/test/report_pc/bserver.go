package main

import (
	"net"
	"strconv"
	"fmt"
	"time"
	"os"
)

var local_port int = 65535

func get_internal_ip()(str string) {
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

func bserver()  {
	udpaddr,err:=net.ResolveUDPAddr("udp4",":"+strconv.Itoa(local_port))
	if(nil !=err){
		fmt.Println(err)
		return
	}
	conn,err:=net.ListenUDP("udp",udpaddr)
	if(nil !=err){
		fmt.Println(err)
		return
	}
	var buf [1024]byte
	for{
		n,raddr,err:=conn.ReadFromUDP(buf[0:])
		if(nil !=err){
			fmt.Println(err)
			return
		}
		fmt.Println(time.Now(),raddr.IP.String(),strconv.Itoa(raddr.Port),"detecting me :"+string(buf[0:n]))

		send_udp(raddr.IP.String(),string(buf[0:n]),get_internal_ip())
	}
}

func send_udp(ip string,port string,str string){
	fmt.Println(ip,port,str)
	rAddr,err:=net.ResolveUDPAddr("udp4",ip+":"+port)
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

	_,err=conn.Write([]byte(str))
	if(nil !=err){
		fmt.Println(err)
		fmt.Println("aa")
		return
	}
	//time.Sleep(time.Second*5)
}



func main()  {

	bserver()
}
