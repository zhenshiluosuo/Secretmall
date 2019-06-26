package models

import "github.com/astaxie/beego/orm"

type User struct {
	Id         int `orm:"auto"`
	Name       string
	Password   string
	Permission int
}

type Cert struct {
	Id     int     `orm:"auto"`
	UserId *User   `orm:"rel(fk)"`
	ItemId *Item   `orm:"rel(fk)"`
}
type Item struct {
	Id          int `orm:"auto"`
	Name        string
	Price       float32
	Amount      int
	descirption string `orm:"size(50)"`
}

func init() {
	orm.RegisterModel(new(User), new(Cert), new(Item))
}
