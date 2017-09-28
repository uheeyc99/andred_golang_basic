package main

import (

	"os"
	"fmt"
	"path/filepath"
)

var size int64 = 0
func AndrewwalkFunc(path string,info os.FileInfo,err error)error{
	fmt.Println(path,info.Size())
	size +=info.Size()
	return nil
}


func main(){

	filepath.Walk(".",AndrewwalkFunc)
	fmt.Println(size,size/1024/1024)

}
