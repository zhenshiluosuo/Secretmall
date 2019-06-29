package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

/*
   输入：传入物品
   功能：添加物品
   输出：添加结果
*/
func Add_item(item Item) bool {
	o := orm.NewOrm()
	err := o.Using("default")
	if err != nil {
		fmt.Println("数据库不存在")
		return false
	}
	_, err = o.Insert(&item)
	if err != nil {
		fmt.Println("插入失败", err)
		return false
	}
	return true
}

/*
   输入：传入字符串，查询相似物品
   功能：查询物品
   输出：查询物品集合
*/
func Query_item(str string) []Item {
	var ii []Item
	o := orm.NewOrm()
	err := o.Using("default")
	if err != nil {
		fmt.Println("数据库不存在")
		return nil
	}
	qs := o.QueryTable("item")
	count, err := qs.Count()
	if count == 0 || err != nil {
		return nil
	}
	_, err = qs.Limit(5).Filter("name__icontains", str).All(&ii)
	if err != nil {
		fmt.Println("查询失败")
	}
	return ii
}

/*
   输入：limit限制返回条数，offset数据库偏移量
   功能：列出物品列表
   输出：物品集合
*/
func Item_List(limit int, offset int) []Item {
	var ii []Item
	o := orm.NewOrm()
	err := o.Using("default")
	if err != nil {
		fmt.Println("数据库不存在")
		return nil
	}
	qs := o.QueryTable("item")
	_, err = qs.Count()
	if err != nil {
		return nil
	}
	_, err = qs.Limit(limit, offset).All(&ii)
	if err != nil {
		fmt.Println("查询失败")
	}
	return ii
}
