package main

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"fmt"
)

type staffAdditional struct{
	Name string `bson:"name"`
	Age int32 `bson:"age"`
	Description string `bson:"describe"`
	Portrait string
}

type Staff struct {
	Id bson.ObjectId `bson:"_id"`
	NickName string `bson:"nickname"`
	ModifyTime time.Time `bson:"modifytime"`
	Additional staffAdditional `bson:"additionalInfo"`
}





//***********************************************
func AddStaff(staff Staff)(string,error){

	collection:=getStaffCollection()
	defer collection.Database.Session.Close()

	staff.Id=bson.NewObjectId()
	staff.ModifyTime=time.Now()
	err:=collection.Insert(staff)
	if err!=nil{
		fmt.Println("AddStaff",err)
	}
	return staff.Id.Hex(),err

}

func DeleteStaffById(id string) error{
	collection:=getStaffCollection()
	defer collection.Database.Session.Close()

	err:=collection.RemoveId(bson.ObjectIdHex(id))
	if err!=nil{
		fmt.Println("DeleteStaffById:",err)
	}
	return err
}

func DeleteStaffAll(){
	collection:=getStaffCollection()
	defer collection.Database.Session.Close()

	info,err:=collection.RemoveAll(nil)
	//info,err:=collection.RemoveAll(bson.M{"telephone":"10086"})

	if err!=nil{
		fmt.Println("DeleteStaffAll:",err)
	}
	fmt.Println(info,err)
}


func UpdateStaffById(id string,staff Staff)error{
	collection:=getStaffCollection()
	defer collection.Database.Session.Close()
	err:=collection.UpdateId(bson.ObjectIdHex(id),staff)
	if err!=nil{
		fmt.Println("UpdateStaffById:",err)
	}
	return err
}

func PageStaff()([]Staff,error){
	collection:=getStaffCollection()
	defer collection.Database.Session.Close()

	staffs:=[]Staff{}
	err:=collection.Find(nil).All(&staffs)
	if err!=nil{
		fmt.Println("PageStaff:",err)
	}
	return staffs,err
}

func GetStaffById(id string)(Staff,error){
	collection:=getStaffCollection()
	defer collection.Database.Session.Close()

	staff:=new(Staff)
	objectid:=bson.ObjectIdHex(id)
	err:=collection.FindId(objectid).One(staff)
	if err!=nil{
		fmt.Println("GetStaffById:",err)
	}
	return *staff,err
}

func GetStaffByNickname(nick string){
	collection:=getStaffCollection()
	defer collection.Database.Session.Close()

	staff:=new(Staff)

	err:=collection.Find(bson.M{"nickname":nick}).One(&staff)
	if err!=nil{
		fmt.Println("GetStaffByNickname:",err)
	}
	fmt.Println(staff)
}




func test_staff(){

	staff:=Staff{
		NickName:"欧吉",
	}
	staff.NickName="eric"
	staff.Additional.Age=9999

	id,err:=AddStaff(staff)
	if err!=nil{
		fmt.Println(err)
	}else{
		//fmt.Println(id)
		staff,e:=GetStaffById(id)
		if e!=nil{
			fmt.Println(e)
		}
		//staff.Phone="10010"
		UpdateStaffById(id,staff)
		GetStaffById(id)
		//GetStaffByTelephone("10011")
		//DeleteStaffById(id)

	}
	//DeleteStaffAll()
	//staffs,err:=PageStaff()

	//fmt.Println(len(staffs))
	//fmt.Println(staffs[len(staffs)-1:len(staffs)]) //打印最后一组
	w.Done()
}