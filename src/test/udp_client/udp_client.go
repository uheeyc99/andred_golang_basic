package main

import (
	"net"
	"fmt"
	"time"
	"strconv"
	"net/http"
	"os"
	"io"
)

func Ask(i int){

	rAddr,err:=net.ResolveUDPAddr("udp4","127.0.0.1:65535")
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
	n,_,err:=conn.ReadFromUDP(buf[0:])
	if(nil !=err){
		fmt.Println(err)
		//fmt.Println("uu")
		return
	}
	fmt.Println("received response: " + string(buf[0:n]))
	fmt.Println(rAddr.IP,rAddr.Port)

}

func get_external(){
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		os.Exit(1)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}

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


func main()  {

	fmt.Println(get_internal())

	t1:=time.Now()
	fmt.Println(t1)
	for i:=0;i<1000;i++{
		time.After(10000)
		Ask(i)
	}
	fmt.Println(time.Now().Sub(t1))
}