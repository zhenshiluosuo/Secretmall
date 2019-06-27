package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

/*
   输入：
   功能：清空购物车
   输出：结果true为成功，false为失败
*/
//func Clean_cert()(bool){
//	o := orm.NewOrm()
//	err := o.Using("default")
//	if err != nil{
//		fmt.Println("数据库不存在")
//		return false
//	}
//}

/*
   输入：表单id、用户id、物品id、购买物品件数
   功能：添加购物车
   输出：添加结果true为添加成功
                false为添加失败
*/

func Add_to_cert(id int, userid int, itemid int, num int) bool {
	o := orm.NewOrm()
	err := o.Using("default")
	if err != nil {
		fmt.Println("数据库不存在")
		return false
	}

	//检测数据库中有没有已存数据
	cert := Cert{UserId: userid, ItemId: itemid}
	var it Item
	err_item := o.QueryTable("Item").Filter("Id", itemid).One(&it)
	_ = o.QueryTable("Cert").Filter("userid", cert.UserId).Filter("itemid", cert.ItemId).One(&cert)
	count := cert.Amount + num

	fmt.Println(cert.Id)
	if err_item != nil {
		fmt.Println("此商品不存在")
		return false
	}
	if count > it.Amount {
		fmt.Println("库存不够了")
		return false
	} else if cert.Id == 0 {
		//订单不存在直接添加
		cert.Id = id
		cert.Amount = count
		rid, err := o.Insert(&cert)
		fmt.Println(rid, err)
		return true
	} else if cert.Id != 0 {
		cert.Amount = count
		rid, err := o.Update(&cert)
		fmt.Println(rid, err)
		return true
	}
	return false
}

/*
   输入：用户id
   功能：查询用户购物车列表
   输出：物品数组
*/
func Check_cert(user User) []Item {
	o := orm.NewOrm()
	var ii []Cert
	var items []Item
	var u User
	err := o.Using("default")

	if err != nil {
		fmt.Println("数据库不存在")
		return nil
	}
	//用户检测
	err = o.QueryTable("User").Filter("Name", user.Name).One(&u)
	if err != nil {
		fmt.Println("用户不存在")
		return nil
	}
	_, err = o.QueryTable("Cert").Filter("userid", u.Id).All(&ii)
	if err != nil {
		fmt.Println("不买东西要什么购物车")
	}
	var it Item
	for _, v := range ii {
		it.Id = v.ItemId
		_ = o.Read(&it)
		items = append(items, it)
	}
	return items
}
