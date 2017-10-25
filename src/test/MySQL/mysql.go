package main

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

func test(){
	fmt.Println("aaaq")
	db,err:=sql.Open("mysql","root:123456789@tcp(qycam.com:50201)/test?charset=utf8")
	if err !=nil{
		fmt.Printf("Open database error: %s\n",err)
	}
	//defer db.Close()
	err=db.Ping()
	if err!=nil{
		fmt.Println("Ping err",err)
	}


}

func main(){

	test()

}
