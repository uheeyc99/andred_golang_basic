package main

import (
	"encoding/json"
	"fmt"
	"os"
	"bufio"
)

//json.Marshal()
//json.Unmarshal()

type Person struct {
	Index int
	Name string
	Sex string
	Age int
	Comments [2]string
}


func enc1(){

	p1:= Person{
		Index:1,
		Name:"KongJun",
		Sex:"M",
		Age:23,
	}

	jp,err:=json.Marshal(p1)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("llll:",jp)
	//write
	f,err:=os.Create("bin/jsontext1-1.txt")
	if err!=nil{
		fmt.Println(err)
	}
	defer f.Close()
	f.Write(jp)
}

func dec1(){
	//read
	f1,err:=os.Open("bin/jsontext1-1.txt")
	if err!=nil{
		fmt.Println(err)
	}
	defer f1.Close()
	reader:=bufio.NewReader(f1)
	line,_,_:=reader.ReadLine()

	//modify
	var p1 Person
	json.Unmarshal(line,&p1)
	p1.Comments[0] = "good person"
	p1.Comments[1] = "!!"
	jp,_:=json.Marshal(p1)

	//write
	f2,err:=os.Create("bin/jsontext1-2.txt")
	if err!=nil{
		fmt.Println(err)
	}
	defer f2.Close()
	f2.Write(jp)

	fmt.Println(p1)
	fmt.Println(jp)

}

func test1(){
	enc1()
	dec1()


}

func main(){
	test1()

}
