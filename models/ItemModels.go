package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Add_item(item Item)  (bool){
	o := orm.NewOrm()
	err :=o.Using("default")
	if err != nil {
		fmt.Println("数据库不存在")
		return false
	}
	_ ,err = o.Insert(&item)
	if err != nil {
		fmt.Println("插入失败",err)
		return false
	}
	return true
}
func Query_item(str string ) ([]Item){
	var ii []Item
	o := orm.NewOrm()
	err := o.Using("default")
	if err != nil{
		fmt.Println("数据库不存在")
		return nil
	}
	qs := o.QueryTable("item")
	count,err :=qs.Count()
	if count==0 || err != nil{
		return nil
	}
	_,err=qs.Limit(5).Filter("name__icontains", str).All(&ii)
	if err != nil{
		fmt.Println("查询失败")
	}
	return ii
}
func Item_List(limit int,offset int)  ([]Item){
	var ii []Item
	o := orm.NewOrm()
	err := o.Using("default")
	if err != nil {
		fmt.Println("数据库不存在")
		return nil
	}
	qs := o.QueryTable("item")
	_ ,err= qs.Count()
	if err != nil {
		return nil
	}
	_,err=qs.Limit(limit,offset).All(&ii)
	if err != nil{
		fmt.Println("查询失败")
	}
	return ii
}
