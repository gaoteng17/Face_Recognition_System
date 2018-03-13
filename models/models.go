package models

import(
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type Results struct{
	Id int
	School_id string
	Class_Name string
	Class_Time string
	Class_Teacher string
	Class_State string
}

func (rest *Results)AddRecords(){
	o:=orm.NewOrm()

	p,err:=o.Raw("insert into records(Id,School_id,Class_Name,Class_Time,Class_Teacher,Class_State) VALUES(?,?,?,?,?,?);").Prepare()
	if err!=nil{
		fmt.Println("Records Insert Failed")
	}
	_,err=p.Exec(rest.Id,rest.School_id,rest.Class_Name,rest.Class_Time,rest.Class_Teacher,rest.Class_State)
	if err!=nil{
		fmt.Println("Records Insert Failed")
	}
	p.Close()
}

type Identity struct{
	Id int
	School_id string
	Name string 
	Class string
}

func (iden *Identity)AddIdentity(){
	o:=orm.NewOrm()

	// id,err:=o.Insert(&iden)

	// fmt.Println(id)
	// fmt.Println(err)
	p,err:=o.Raw("insert into identity(Id,School_id,Name,Class) VALUES(?,?,?,?);").Prepare()
	if err!=nil{
		fmt.Println("Identity Insert Failed")
	}
	_,err=p.Exec(iden.Id,iden.School_id,iden.Name,iden.Class)
	if err!=nil{
		fmt.Println("Identity Insert Failed")
	}
	p.Close()
}

func (iden *Identity)UpdateIdentity(value string){

	if value==""{
		fmt.Println("parameters error")
		return 
	}

	o:=orm.NewOrm()
	p,err:=o.Raw("UPDATE identity SET ?=? WHERE ?=?").Prepare()
	if err!=nil{
		fmt.Println("Identity Update Failed")
	}

	switch value{
	case "School_id":
		_,err=p.Exec("School_id",iden.School_id,"School_id",value)
	case "Name":
		_,err=p.Exec("Name",iden.Name,"Name",value)
	case "Class":
		_,err=p.Exec("Class",iden.Class,"Class",value)
	default:
		fmt.Println("Identity Update Failed")
	}

	if err!=nil{
		fmt.Println("Identity Update Failed")
	}
	p.Close()
}

func (iden *Identity)RemoveIdentity(){
	o:=orm.NewOrm()
	_,err:=o.Raw("DELETE FROM identity WHERE School_id=?",iden.School_id).Exec()
	if err!=nil{
		fmt.Println("Identity Delete Failed")
	}
}

func (iden *Identity)SerachIdentity(){
	o:=orm.NewOrm()
	_,err:=o.Raw("SELECT School_id FROM identity WHERE School_id=?",iden.School_id).Exec()
	if err!=nil{
		fmt.Println("Identity Search Failed")
	}
}


func init(){
	orm.RegisterDataBase("default", "mysql", "root:lj1512510237@/mysql?charset=utf8")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//orm.RegisterModel(new(Identity))
	//orm.RunSyncdb("default", false, true)
}

