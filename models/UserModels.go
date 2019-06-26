package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)



func Add_user(user User) (bool){
	o := orm.NewOrm()
	err :=o.Using("default")
	if err != nil {
		fmt.Println("数据库不存在")
		return false
	}
	_ ,err = o.Insert(&user)
	if err != nil {
		fmt.Println("插入失败",err)
		return false
	}
	return true
}
func Query_single_user(name string)  (User){
	o := orm.NewOrm()
	err := o.Using("default")
	user := User{Name:name}
	if err != nil {
		fmt.Println("数据库不存在")
		return user
	}
	err = o.Read(&user,"name")
	if err == orm.ErrNoRows{
		fmt.Println("无此用户")
	}else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	}else {
		fmt.Println(user.Id,user.Name,user.Password)
	}
	return user
}