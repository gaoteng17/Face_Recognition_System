package controllers

import (
	"os/exec"
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

func checkout(School_id string) (string) {
    file1 := "/root/img/tmp/"+School_id+".jpg"
    file2 := "/root/img/"+School_id+"/1.jpg"

    fmt.Println("debug1")
    cmd := exec.Command("python", "/root/Face_Recognition_System/ServerTest/ServerTest.py", file1, file2)
   // cmd := exec.Command("ls")
    fmt.Println("debug2")
    buf, err := cmd.Output()
    if err != nil {
        fmt.Println("debug3")
        fmt.Println(err)
    }
    fmt.Println(string(buf))
    return string(buf)
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
    result := ""    

    switch ob.Options{
    case "delete":
    	iden.RemoveIdentity()
    case "update":
    	iden.UpdateIdentity(ob.Parameters)
    case "insert":
    	iden.AddIdentity()
    case "search":
    	iden.SerachIdentity()
    case "checkout":
	result=checkout(ob.School_id)
    default:
    	fmt.Println("options error")
    }
 

    switch ob.Options{
    case "checkout":
        this.Data["json"] = map[string]interface{}{"result": result, "msg": "checkout over"}
    default: 
        this.Data["json"] = map[string]interface{}{"result": true, "msg": "Success"}
    }    

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

    err=this.SaveToFile("img","/root/img/tmp/"+school_id+".jpg")
    if err != nil {
        fmt.Println("save tmpfile error")
        fmt.Println(err)
    }

    
}

