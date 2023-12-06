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
	// 内部业务逻辑层
	case "logic":
		logic.New().Run()
	// websocket 服务
	case "connect_websocket":
		connect.New().Run()
	// tcp 长链接
	case "connect_tcp":
		connect.New().RunTcp()
	// 执行任务服务
	case "task":
		task.New().Run()
	// api服务（前端提供接口）
	case "api":
		api.New().Run()
	// 前端服务
	case "site":
		site.New().Run()
	default:
		fmt.Println("exiting,module param error!")
		return
	}
	fmt.Println(fmt.Sprintf("run %s module done!", module))
	// 创建一个无缓冲通道， 类型为： os.Signal
	quit := make(chan os.Signal)
	/*
		在 Go 语言中，os/signal 包允许你处理系统信号（例如 SIGINT、SIGTERM 等），以便在程序运行时优雅地处理终止信号。
		signal.Notify 函数用于将指定的信号发送到一个通道。通常情况下，这个通道用于接收并处理操作系统发送的信号。
		package main
		import (
			"fmt"
			"os"
			"os/signal"
			"syscall"
		)
		func main() {
			// 创建一个通道来接收信号
			quit := make(chan os.Signal, 1)
			// 将指定的信号发送到 quit 通道
			signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
			// 等待接收信号
			sig := <-quit
			fmt.Println("Received signal:", sig)
		}
		在这个示例中：
		signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		将指定的信号（SIGHUP、SIGINT、SIGTERM、SIGQUIT）发送到 quit 通道中。
		<-quit 用于等待并接收来自 quit 通道的信号。一旦接收到信号，程序会打印出所接收到的信号。
		通过这种方式，你可以在 Go 程序中设置信号监听器，并根据接收到的信号执行相应的操作。
		通常，这种方法可以用于优雅地关闭服务或执行一些清理操作，以便程序能够在接收到终止信号时进行正确的处理。
	*/
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
	fmt.Println("Server exiting")
}
