package main

import (
	"net"
	"strconv"
	"fmt"
)

var try_port int = 65530

func try_pc(ip_str string){
	conn,err:=net.Dial("tcp",ip_str+":"+strconv.Itoa(try_port))
	if err!=nil{
		fmt.Println(conn)
		return
	}
	defer conn.Close()

	_,err=conn.Write([]byte("eric"))
	if err!=nil{
		fmt.Println(conn)
		return
	}

	buffer:=make([]byte,1024)
	n,err:=conn.Read(buffer)
	if err!=nil{
		fmt.Println(conn)
		return
	}else{
		fmt.Println(string(buffer[0:n]))
	}


}

func main()  {
	try_pc("127.0.0.1")
}
