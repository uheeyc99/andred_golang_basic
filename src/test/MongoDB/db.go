package main

import (
	"gopkg.in/mgo.v2"
	"time"
	"fmt"
)

var (
	mgoSession *mgo.Session
	dataBase ="test"
	staffCollection="staff"
)

var dbInfo =mgo.DialInfo{
	Addrs: []string{"qycam.com:50203"},
	Direct:false,
	Timeout:time.Second*5,
	Database:dataBase,
	Username:"andrew",
	Password:"123456789",
	PoolLimit:4096,

}

var h = 1

func getSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession,err = mgo.DialWithInfo(&dbInfo)
		if err!=nil{
			fmt.Println("getSession:",h,err)
			panic(err)
		}
	}

	return mgoSession.Clone()
}

func getStaffCollection() *mgo.Collection{
	session:=getSession()
	collection:=session.DB(dataBase).C(staffCollection)
	return collection

}