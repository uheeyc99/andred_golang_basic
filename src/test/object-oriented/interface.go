package main

import "fmt"

type Human struct{
	name string
	age int
	phone string

}

type Student struct{
	name string
	age int
	phone string

}


func (h *Human)set_name(name string){
	h.name = name
}
func (s *Human)set_age(age int){
	h.age=age
}


var h Human
func main(){

	h.set_name("KK")
	fmt.Println(h.name)
}
