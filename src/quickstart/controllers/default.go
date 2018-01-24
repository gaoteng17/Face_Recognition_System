package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"encoding/json"
	"quickstart/models"
	//"path"
	//"strings"
)

type MainController struct {
	beego.Controller
}




type IdentityController struct{
	beego.Controller
}

type Identity struct{
	School_id string
	Name string 
	Class string
	Options string
	Parameters string
}


func (c *MainController) Get() {
	fmt.Println("dsoivhodsvsd")
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}



func (this *IdentityController) Post() {

	var ob Identity
    err:=json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
    if err==nil{
    	fmt.Println("Unmarshal Success")
    }
    
    fmt.Println(string(this.Ctx.Input.RequestBody))
    fmt.Println(ob)
    fmt.Println(ob.School_id)
    fmt.Println(ob.Name)
    fmt.Println(ob.Class)

    iden:=models.Identity{0,ob.School_id,ob.Name,ob.Class}

    switch ob.Options{
    case "delete":
    	iden.RemoveIdentity()
    case "update":
    	iden.UpdateIdentity(ob.Parameters)
    case "insert":
    	iden.AddIdentity()
    case "search":
    	iden.SerachIdentity()
    default:
    	fmt.Println("options error")
    }
 

    
    this.Data["json"] = map[string]interface{}{"result": true, "msg": "新增成功"}
    this.ServeJSON()
}





type ImgController struct {
	beego.Controller
}

// func (this *ImgController)Get(){

// }

func (this *ImgController)Post(){

	var school_id string
	this.Ctx.Input.Bind(&school_id, "school_id") 
	var name string
	this.Ctx.Input.Bind(&name, "name") 
	var class string
	this.Ctx.Input.Bind(&class, "class") 

	fmt.Println(school_id)
	fmt.Println(name)
	fmt.Println(class)

	f,_,err:=this.GetFile("img")
    if err != nil {
        fmt.Println("getfile error ")
    }

    // fileName := h.Filename 
    // arr:=strings.Split(fileName, ":")
    // if len(arr) > 1 {   
    //   index:=len(arr)-1
    //   fileName=arr[index]
    // }

    f.Close()

    err=this.SaveToFile("img","static/img/"+school_id+class+name+".png")
    if err != nil {
        fmt.Println("savefile error")
        fmt.Println(err)
    }
}

