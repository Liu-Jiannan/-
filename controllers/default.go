package controllers

import (
	"HelloBeego/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//name := c.GetString("name")
	//age,err := c.GetInt("age")


	//获取get类型请求的参数

    name := c.Ctx.Input.Query("name")
    age := c.Ctx.Input.Query("age")
    sex := c.Ctx.Input.Query("sex")
    fmt.Println(name,age,sex)

    //以admin,18为正确数据进行验证
	if name != "admin" || age != "18" {
		c.Ctx.ResponseWriter.Write([]byte("数据验证错误"))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("数据验证正确"))


	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

/**
该post方法是处理post类型的请求的时候要调用的方法
 */

func (c *MainController) Post(){
	fmt.Println("post类型的请求")
	user := c.Ctx.Request.FormValue("user")
    fmt.Println("用户名为:",user)
	psd := c.Ctx.Request.FormValue("psd")
	fmt.Println("密码是:",psd)

	//与固定值比较
	if user != "admin" || psd != "123456"{
		//失败页面
		c.Ctx.ResponseWriter.Write([]byte("对不起，数据不正确"))
		return
	}

	c.Ctx.ResponseWriter.Write([]byte("恭喜你，数据正确"))

	//qequest 请求，response 响应
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"



    //以下是便捷式解析 (JSON解析) . 在models文件夹下建立JSON接口文件,与前端连接

	body := c.Ctx.Request.Body
	dataByes,err := ioutil.ReadAll(body)
	if err != nil{
		c.Ctx.WriteString("数据接收失败，请重试")
		return
	}

	var person models.Person
	err = json.Unmarshal(dataByes,&person)
	if err!= nil {
		c.Ctx.WriteString("数据解析失败，请重试")
		return
	}
	fmt.Println("用户名",person.Name,",年龄",person.Age)
	c.Ctx.WriteString("用户名是:"+person.Name)
}