package models

import "github.com/astaxie/beego/orm"

type User struct {
	Id         int `orm:"auto"`
	Name       string
	Password   string
	Permission int
}

type Cert struct {
	Id     int `orm:"auto"`
	UserId int `orm:"column(user_id)"`
	ItemId int `orm:"column(item_id)"`
	Amount int
}
type Item struct {
	Id          int `orm:"auto"`
	Name        string
	Price       float32
	Amount      int
	Descirption string
}

func init() {
	orm.RegisterModel(new(User), new(Cert), new(Item))
}
