package controllers

import (
	"os/exec"
	"github.com/astaxie/beego"
	"fmt"
	//"encoding/json"
	"quickstart/models"
	//"path"
	//"strings"
    "os"
)

//default Controller without using
type MainController struct {
	beego.Controller
}
//test Controller without using
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

//default get method without using
func (c *MainController) Get() {
	fmt.Println("dsoivhodsvsd")
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
    //c.TplName = "index.tpl"
    c.TplName = "index.html"
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
    result:=string(buf)
    //judge result
    if(result[0]=='T'){
        err := os.Rename(file1,file2)
        if err != nil {
            fmt.Println("move tmp file failed")
        }
        return "present"
    }else if(result[0]=='F'){
        err := os.Remove(file1)
        if err != nil {
            fmt.Println("delete tmp file failed")
        }
        return "absent"
    }
    return ""
}

/*function for test*/

// func (this *IdentityController) Post() {

// 	var ob Identity
//     err:=json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
//     if err==nil{
//     	fmt.Println("Unmarshal Success")
//     }
    
//     fmt.Println(string(this.Ctx.Input.RequestBody))
//     fmt.Println(ob)
//     fmt.Println(ob.School_id)
//     fmt.Println(ob.Name)
//     fmt.Println(ob.Class)

//     iden:=models.Identity{0,ob.School_id,ob.Name,ob.Class}
//     result := ""    

//     switch ob.Options{
//     case "delete":
//     	iden.RemoveIdentity()
//     case "update":
//     	iden.UpdateIdentity(ob.Parameters)
//     case "insert":
//     	iden.AddIdentity()
//     case "search":
//     	iden.SerachIdentity()
//     case "checkout":
// 	result=checkout(ob.School_id)
//     default:
//     	fmt.Println("options error")
//     }
 

//     switch ob.Options{
//     case "checkout":
//         this.Data["json"] = map[string]interface{}{"result": result, "msg": "checkout over"}
//     default: 
//         this.Data["json"] = map[string]interface{}{"result": true, "msg": "Success"}
//     }    

//     this.ServeJSON()
// }





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
    var options string
    this.Ctx.Input.Bind(&options, "options") 
    var parameters string 
    this.Ctx.Input.Bind(&parameters,"parameters")

	fmt.Println(school_id)
	fmt.Println(name)
	fmt.Println(class)
    fmt.Println(options)


	// f,_,err:=this.GetFile("img")
 //    if err != nil {
 //        fmt.Println("getfile error ")
 //    }

 //    f.Close()

 //    //save img to temp path
 //    err=this.SaveToFile("img","/root/img/tmp/"+school_id+".jpg")
 //    if err != nil {
 //        fmt.Println("save tmpfile error")
 //        fmt.Println(err)
 //    }

    var result string
    //database table struct
    iden:=models.Identity{0,school_id,name,class}
    //judge options
    switch options{
    case "delete":
        iden.RemoveIdentity()
    case "update":
        iden.UpdateIdentity(parameters)
    case "insert":
        iden.AddIdentity()
    case "search":
        iden.SerachIdentity()
    case "checkout": 
    //result=checkout(school_id)
    result="debug"
    //store results
    rest:=models.Results{0,school_id,class,"","",result}
    rest.AddRecords()
    default:
        fmt.Println("options error")
    }


    //message return to the client
    switch options{
    case "checkout":
        this.Data["json"] = map[string]interface{}{"result": result, "msg": "checkout over"}
    default: 
        this.Data["json"] = map[string]interface{}{"result": true, "msg": "Success"}
    }    

    this.ServeJSON()
}

