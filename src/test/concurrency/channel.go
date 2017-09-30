package main

import (
	"fmt"
	"time"
	"sync"
	"runtime"
)

//*******************************************************

func calculate(ch chan int){
	fmt.Println(time.Now().String(),"子线程开始处理")
	time.Sleep(time.Second)
	fmt.Println(time.Now().String(),"子线程处理结束，等待主线程读取结果")
	ch<-1000
	fmt.Println(time.Now().String(),"子线程汇报结束，继续执行其它操作")
}

func hanldeconnection(i int,ch chan int){
	defer func() {
		<-ch  //读出数据，让出channel的buffer
	}()
	fmt.Println(time.Now().String(),"进入子线程",i)
	fmt.Println(time.Now().String(),"子线程",i,"开始处理")
	time.Sleep(time.Second)

}

func get_data(ch chan int){
	fmt.Println(time.Now().String(),"主线程开始调用子线程")
	//此处做一些耗时的准备工作
	<-ch //已就位，等待命令
	fmt.Println(time.Now().String(),"子线程收到信号开始执行")
}

func sub_thread(){


	defer func() {
		w.Done()
	}()

	time.Sleep(time.Millisecond*1200)

}

func say(s string)(){
	fmt.Println(s)
}

//*******************************************************

var w sync.WaitGroup
func test_waitgroup(){
	for i:=0;i<3;i++{
		w.Add(1) //
		go sub_thread()
	}

	fmt.Println(time.Now())
	w.Wait() //等待 w清零
	fmt.Println(time.Now())
}

func test_gosched(){//让出时间片
	runtime.GOMAXPROCS(1)  //强制设置单核执行
	go fmt.Println("a")
	go fmt.Println("b")
	fmt.Println("c")
	runtime.Gosched()
	fmt.Println("d")
	//单核执行顺序是 c d a b，而加了runtime.Gosched()之后则d最后执行
	time.Sleep(time.Second)
}



func test_channel_01(){//无缓冲channel
	// channel
	ch := make(chan int)  //无缓冲channel,读写都会阻塞
	fmt.Println(time.Now().String(),"主线程开始调用子线程")
	go calculate(ch)
	time.Sleep(time.Second*2)  //注释掉试试，无论谁先执行都要等对方
	fmt.Println(time.Now().String(),"主线程尝试读取子线程数据")
	value:= <-ch
	fmt.Println(time.Now().String(),"主线程查看子线程数据",value)
	time.Sleep(time.Second)
}



func test_channel_02(){//有缓冲channel
	//buffered channel
	ch:=make(chan int,2)  //有缓冲channel，只要没写满，就不会写阻塞
	fmt.Println(time.Now().String(),"主线程开始调用子线程")
	go calculate(ch)
	time.Sleep(time.Second*2)  //
	fmt.Println(time.Now().String(),"主线程尝试读取子线程数据")
	value:= <-ch
	fmt.Println(time.Now().String(),"主线程查看子线程数据",value)
	time.Sleep(time.Second)
	//只要channel缓冲区没满，子线程把数据写入channel后，继续执行其它的操作了，无论主线程有没有去读


}

func test_channel_03(){ //利用有缓冲channel控制并发线程数量
	ch:=make(chan int,3)
	for i:=0;i<5;i++{
		ch<-i   //写满3时阻塞, 只允许并行执行3个子线程
		go hanldeconnection(i,ch)
	}
	time.Sleep(time.Second*5)
}



func test_channel_04(){  //利用channel实现类似于线程池的功能
	ch:=make(chan int,1)
	go get_data(ch)  //调用子线程，子线程开始筹备自己的工作

	time.Sleep(time.Second)  //模拟其它操作

	fmt.Println(time.Now().String(),"命令子线程开始执行操作")
	ch<-1

	time.Sleep(time.Second)

}

func test_channel_05(){  //数组

}







func main(){
	test_gosched()
	test_channel_05()

}
