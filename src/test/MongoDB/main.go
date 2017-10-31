package main

import (
	"sync"

	"time"
	"fmt"
)
var w sync.WaitGroup
func main(){
	t1:=time.Now()
	for i:=0;i<1;i++{
		go w.Add(1)
		test_staff()
	}
	w.Wait()
	td:=time.Since(t1)
	fmt.Println(td)
}