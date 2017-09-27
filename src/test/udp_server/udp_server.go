package main

import (
	"net"
	"fmt"
	"strconv"
	"time"
)

var udp_port int

func udp_response(){
	udp_port=60000
	udpaddr,err:=net.ResolveUDPAddr("udp4",":"+strconv.Itoa(udp_port))
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

		conn.WriteToUDP([]byte("tks..."),raddr)

	}

}

func main(){
	udp_response()
}