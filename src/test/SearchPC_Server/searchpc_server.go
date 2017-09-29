package main

import (
	"net"
	"os"
	"strconv"
	"fmt"
)

var listen_port int = 65530

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


func handleConnection(c net.Conn){
	defer c.Close()
	buffer := make([]byte,1024)
	_,err:=c.Read(buffer)
	if err!=nil {
		fmt.Println(err)
		return
	}
	c.Write([]byte(get_internal_ip()))

}

func server(){
	l,err:=net.Listen("tcp",":"+strconv.Itoa(listen_port))
	if err != nil{
		fmt.Println(err)
		return
	}
	defer l.Close()

	for{
		c,err:= l.Accept()
		if(err!=nil){
			fmt.Println(err)
			return
		}
		handleConnection(c)
	}

}


func main(){

	server()

}
