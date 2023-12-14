/**
 * Created by lock
 * Date: 2019-08-09
 * Time: 18:25
 */
package logic

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gochat/config"
	"runtime"
)

// Logic 逻辑层结构体
type Logic struct {
	// 服务ID
	ServerId string
}

// New 逻辑层服务指针对象
func New() *Logic {
	return new(Logic)
}

// Run 启动逻辑层服务
// Run 通过实现一个指针对象方法来启动服务
func (logic *Logic) Run() {
	// 加载 logic 逻辑层配置
	logicConfig := config.Conf.Logic
	// 设置程序可以同时执行的CPU最大核心数
	runtime.GOMAXPROCS(logicConfig.LogicBase.CpuNum)
	// 设置程序的服务ID
	logic.ServerId = fmt.Sprintf("logic-%s", uuid.New().String())
	// 初始化 redis 连接对象
	if err := logic.InitPublishRedisClient(); err != nil {
		logrus.Panicf("logic init publishRedisClient fail,err:%s", err.Error())
	}

	//init rpc server
	if err := logic.InitRpcServer(); err != nil {
		logrus.Panicf("logic init rpc server fail")
	}
}
