package detect

import (
	"net"
	"fmt"
	"strconv"
	"sync"
	"time"
	"bytes"
	"strings"
	"os"
	"bufio"
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
	fmt.Println("detect: ",ip_start,ip_end,p_start,p_end)
	ch1:=make(chan  int,500)

	port_start_int,_:=strconv.Atoi(port_start)
	port_end_int,_:=strconv.Atoi(port_end)
	for ip_int := StringIpToInt(ip_start); ip_int <= StringIpToInt(ip_end); ip_int++ {
		fmt.Println("detecting " +IpIntToString(ip_int) + "...")
		for port_int:=port_start_int;port_int<=port_end_int;port_int++ {
			wait.Add(1)
			ch1 <- ip_int
			go detect_one(IpIntToString(ip_int), strconv.Itoa(port_int), ch1)
		}

	}
}

var conn_timeout time.Duration
var h_start,h_end,p_start,p_end string
func init(){
	h_start="10.10.1.1"
	h_end ="10.10.5.254"
	p_start="1"
	p_end="65535"
}

func Detect_Port(){

	file,err:=os.Open("in.conf")
	if(err!=nil) {
		fmt.Println("where is in.conf")
	}else{
		defer file.Close()

		br:=bufio.NewReader(file)
		var str []byte
		str,_,_ = br.ReadLine()
		h_start =string(str)
		str,_,_ = br.ReadLine()
		h_end =string(str)
		str,_,_ = br.ReadLine()
		p_start =string(str)
		str,_,_ = br.ReadLine()
		p_end =string(str)
	}

	conn_timeout = 10000  //10s
	detect(h_start,h_end,p_start,p_end)
	wait.Wait()

}
