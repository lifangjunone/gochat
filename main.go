/**
 * Created by lock
 * Date: 2019-08-09
 * Time: 10:56
 */
package main

import (
	"flag"
	"fmt"
	"gochat/api"
	"gochat/connect"
	"gochat/logic"
	"gochat/site"
	"gochat/task"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 申明一个命令行参数，用来命令行输入需要运行的模块
	var module string
	// 通过flag包来创建并解析命令行参数, 申明一个字符串变量module
	// 变量名： module， 默认值： “”  变量说明/描述： assign run module
	flag.StringVar(&module, "module", "", "assign run module")
	// 解析命令行参数，这使用命令行变量前这个方法必须先调用
	flag.Parse()
	// 字符串拼接： fmt.Sprintf("start run %s module", module)
	fmt.Println(fmt.Sprintf("start run %s module", module))
	// switch 通过命令行参数 module 的值来启动对应的服务
	switch module {
	case "logic":
		logic.New().Run()
	case "connect_websocket":
		connect.New().Run()
	case "connect_tcp":
		connect.New().RunTcp()
	case "task":
		task.New().Run()
	case "api":
		api.New().Run()
	case "site":
		site.New().Run()
	default:
		fmt.Println("exiting,module param error!")
		return
	}
	fmt.Println(fmt.Sprintf("run %s module done!", module))
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
	fmt.Println("Server exiting")
}
