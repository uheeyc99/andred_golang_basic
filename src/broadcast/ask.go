package broadcast

import (
	"net"
	"fmt"
)

func Ask(){

	rAddr,err:=net.ResolveUDPAddr("udp4","255.255.255.255:60000")
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

	_,err=conn.Write([]byte("hhha"))
	if(nil !=err){
		fmt.Println(err)
		fmt.Println("aa")
		return
	}
	var buf [1024]byte
	n,rAddr,err:=conn.ReadFromUDP(buf[0:])
	if(nil !=err){
		fmt.Println(err)
		fmt.Println("uu")
		return
	}
	fmt.Println("received response: " + string(buf[0:n]))
	fmt.Println(rAddr.IP,rAddr.Port)

}