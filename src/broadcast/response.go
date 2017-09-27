package broadcast

import (
	"net"
	"fmt"

)


func Response_Andrew(){

	udpaddr,err:=net.ResolveUDPAddr("udp4",":60000")
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
		fmt.Println("prepare ing ...")
		n,raddr,err:=conn.ReadFromUDP(buf[0:])
		if(nil !=err){
			fmt.Println(err)
			return
		}
		fmt.Println("someone detecting me :"+string(buf[0:n]))
		fmt.Println(raddr.IP,raddr.Port)
		n,err=conn.WriteToUDP([]byte("tks..."),raddr)
		if(nil !=err){
			fmt.Println(err)
			return
		}
	}



}