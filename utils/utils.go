package utils

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
type db struct {
	driver  string
	host string
	name string
	port string
	user string
	pass string
	maxidle string
	maxconn string
}
func init(){
	orm.RegisterDriver("mysql", orm.DRMySQL)
	d := db{}
	d.driver =beego.AppConfig.String("driver")
	d.host =beego.AppConfig.String("host")
	d.name =beego.AppConfig.String("name")
	d.port= beego.AppConfig.String("port")
	d.user = beego.AppConfig.String("user")
	d.pass = beego.AppConfig.String("pass")
	d.maxidle= beego.AppConfig.String("maxidle")
	d.maxconn= beego.AppConfig.String("maxconn")
	dbconn := 	d.user+":"+d.pass+"@tcp("+d.host+":"+d.port+")/"+d.name+"?"

	err := orm.RegisterDataBase("default",d.driver,dbconn)
	if err != nil{
		fmt.Println(err)
	}else {
		fmt.Println("测试成功")
	}
	//data, err := sql.Open(d.driver,dbconn)
}

//func Modifiy(sql string) (int64){
//
//
//
//	//fmt.Println(sql)
//	//r,err := DB.Exec(sql)
//	//if err != nil{
//    //      fmt.Println(err)
//    //      return 0
//	//}
//	//count,err:= r.RowsAffected()
//	//if err != nil{
//	//	fmt.Println(err)
//	//	return 0
//	//}
//	//return count
//}
//func Query(sql string) *sql.Row{
//	return DB.QueryRow(sql)
//}
