package main 

import(
	"fmt"
	"github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)

type User struct{
	Id int
	Name string
}

func init(){
	    // set default database
    orm.RegisterDataBase("default", "mysql", "username:password@tcp(127.0.0.1:3306)/db_name?charset=utf8", 30)

    // register model
    orm.RegisterModel(new(User))

    // create table
    orm.RunSyncdb("default", false, true)
}

func main(){
	o:=orm.NewOrm()
	user:=User(Name:"slene")

	id,err:=o.Insert(&user)
	fmt.Printf("ID:%d,ERR:%v\n",id,err)

	// update
    user.Name = "astaxie"
    num, err := o.Update(&user)
    fmt.Printf("NUM: %d, ERR: %v\n", num, err)

    // read one
    u := User{Id: user.Id}
    err := o.Read(&u)
    fmt.Printf("ERR: %v\n", err)

    // delete
    num, err = o.Delete(&u)
    fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}