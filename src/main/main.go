package main

import (
	"net"
	"fmt"
	"strconv"
	"sync"
	"time"
	"bytes"
	"strings"

	"os"
)
func StringIpToInt(ipstring string) int {
	ipSegs := strings.Split(ipstring, ".")
	var ipInt int = 0
	var pos uint = 24
	for _, ipSeg := range ipSegs {
		tempInt, _ := strconv.Atoi(ipSeg)
		tempInt = tempInt << pos
		ipInt = ipInt | tempInt
		pos -= 8
	}
	return ipInt
}

func IpIntToString(ipInt int) string {
	ipSegs := make([]string, 4)
	var len int = len(ipSegs)
	buffer := bytes.NewBufferString("")
	for i := 0; i < len; i++ {
		tempInt := ipInt & 0xFF
		ipSegs[len-i-1] = strconv.Itoa(tempInt)
		ipInt = ipInt >> 8
	}
	for i := 0; i < len; i++ {
		buffer.WriteString(ipSegs[i])
		if i < len-1 {
			buffer.WriteString(".")
		}
	}
	return buffer.String()
}
var wait sync.WaitGroup

func detect_one(ip string,port string,ch_x chan int)(err1 error){
	t1:=time.Now()
	defer func() {
		<- ch_x

		//fmt.Println(ip+":"+port +"成功不连接  "+time.Since(t1).String())
	}()
	defer wait.Done()



	//fmt.Println(ip+":"+port)
	conn,err:=net.DialTimeout("tcp",ip+":"+port,conn_timeout*1000000)
	if(err!=nil){

		if(strings.Contains(err.Error(),"timeout") ==false){
			//fmt.Println(err)
		}
		return err
	}
	fmt.Println(ip+":"+port +" detected  "+time.Since(t1).String())


	defer conn.Close()
	return
}

func detect(ip_start string,ip_end string,port_start string,port_end string){

	ch1:=make(chan  int,500)

	port_start_int,_:=strconv.Atoi(port_start)
	port_end_int,_:=strconv.Atoi(port_end)
	for port_int:=port_start_int;port_int<=port_end_int;port_int++ {

		for ip_int := StringIpToInt(ip_start); ip_int <= StringIpToInt(ip_end); ip_int++ {
			wait.Add(1)
			ch1 <- ip_int
			go detect_one(IpIntToString(ip_int), strconv.Itoa(port_int), ch1)
		}

	}
}

var conn_timeout time.Duration

func main(){

	conn_timeout = 10000  //10s
	detect("10.10.1.1","10.10.1.254","1","600")
	//detect(os.Args[1],os.Args[2],os.Args[3],os.Args[4])
	wait.Wait()

	//fmt.Println(time.Now())

}
