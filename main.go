package main

import (
	_ "HelloBeego/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	//1连接mysql数据库










	config := beego.AppConfig
	appName := config.String("appname")
	fmt.Println(" 项目应用名称:",appName)
	port,err := config.Int("httpport")
	if err != nil {//err不为nil，表示连接数据库时出现了错误,
		//配置信息解析错误




		panic("项目配置信息解释错误,请稍后重试")

		fmt.Println("应用监听端口:",port)

	}
	beego.Run()
}