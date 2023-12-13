/**
 * Created by lock
 * Date: 2019-08-09
 * Time: 15:24
 */
package config

import (
	"github.com/spf13/viper"
	"os"
	"sync"
)

// 使用 once.Do(fun_name) 确保 fun_name 只会被执行一次
var once sync.Once

// Conf 定义一个全局的 Config 指针对象
var Conf *Config

// 定义常量
const (
	SuccessReplyCode      = 0
	FailReplyCode         = 1
	SuccessReplyMsg       = "success"
	QueueName             = "gochat_queue"
	RedisBaseValidTime    = 86400
	RedisPrefix           = "gochat_"
	RedisRoomPrefix       = "gochat_room_"
	RedisRoomOnlinePrefix = "gochat_room_online_count_"
	MsgVersion            = 1
	OpSingleSend          = 2 // single user
	OpRoomSend            = 3 // send to room
	OpRoomCountSend       = 4 // get online user count
	OpRoomInfoSend        = 5 // send info to room
	OpBuildTcpConn        = 6 // build tcp conn
)

// Config 配置结构体
type Config struct {
	// 以下全是结构体组合
	// 公共配置
	Common Common
	// 网络链接相关配置
	Connect ConnectConfig
	// 逻辑层服务配置
	Logic LogicConfig
	// 任务层服务配置
	Task TaskConfig
	// Api层服务配置
	Api ApiConfig
	// 站点层服务配置
	Site SiteConfig
}

// config 包初始化方法
func init() {
	Init()
}

func getCurrentDir() string {
	dir, _ := os.Getwd()
	return dir
}

// Init 实际初始化方法
func Init() {
	// once: sync.Once.Do 保证 该方法只执行一次，类似于单例模式
	// 因为是配置初始化，所以只需要在服务启动的时候初始化一次配置对象即可
	// 不需要每次使用配置对象都初始化一次，浪费资源
	once.Do(func() {
		// 获取运行的环境
		env := GetMode()
		// 获取当前项目根目录路径
		realPath := getCurrentDir()
		// 拼接配置文件路径
		configFilePath := realPath + "/" + env + "/"
		// 设置配置文件类型
		viper.SetConfigType("toml")
		// 设置配置文件名字
		viper.SetConfigName("/connect")
		// 设置配置文件路径
		viper.AddConfigPath(configFilePath)
		// 解析配置文件
		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
		// 可以设置多个配置文件名
		viper.SetConfigName("/common")
		// 合并配置文件
		err = viper.MergeInConfig()
		if err != nil {
			panic(err)
		}
		// 设置多个配置文件
		viper.SetConfigName("/task")
		// 合并配置文件
		err = viper.MergeInConfig()
		if err != nil {
			panic(err)
		}
		viper.SetConfigName("/logic")
		err = viper.MergeInConfig()
		if err != nil {
			panic(err)
		}
		viper.SetConfigName("/api")
		err = viper.MergeInConfig()
		if err != nil {
			panic(err)
		}
		viper.SetConfigName("/site")
		err = viper.MergeInConfig()
		if err != nil {
			panic(err)
		}
		Conf = new(Config)
		viper.Unmarshal(&Conf.Common)
		viper.Unmarshal(&Conf.Connect)
		viper.Unmarshal(&Conf.Task)
		viper.Unmarshal(&Conf.Logic)
		viper.Unmarshal(&Conf.Api)
		viper.Unmarshal(&Conf.Site)
	})
}

// GetMode 通过环境变量来动态切换配置不同环境初始化的配置
func GetMode() string {
	// 通过 RUN_MODE 环境变量来运行不同配置在不同环境
	env := os.Getenv("RUN_MODE")
	if env == "" {
		env = "dev"
	}
	return env
}

func GetGinRunMode() string {
	env := GetMode()
	//gin have debug,test,release mode
	if env == "dev" {
		return "debug"
	}
	if env == "test" {
		return "debug"
	}
	if env == "prod" {
		return "release"
	}
	return "release"
}

type CommonEtcd struct {
	Host              string `mapstructure:"host"`
	BasePath          string `mapstructure:"basePath"`
	ServerPathLogic   string `mapstructure:"serverPathLogic"`
	ServerPathConnect string `mapstructure:"serverPathConnect"`
	UserName          string `mapstructure:"userName"`
	Password          string `mapstructure:"password"`
	ConnectionTimeout int    `mapstructure:"connectionTimeout"`
}

type CommonRedis struct {
	RedisAddress  string `mapstructure:"redisAddress"`
	RedisPassword string `mapstructure:"redisPassword"`
	Db            int    `mapstructure:"db"`
}

type Common struct {
	CommonEtcd  CommonEtcd  `mapstructure:"common-etcd"`
	CommonRedis CommonRedis `mapstructure:"common-redis"`
}

type ConnectBase struct {
	CertPath string `mapstructure:"certPath"`
	KeyPath  string `mapstructure:"keyPath"`
}

type ConnectRpcAddressWebsockts struct {
	Address string `mapstructure:"address"`
}

type ConnectRpcAddressTcp struct {
	Address string `mapstructure:"address"`
}

type ConnectBucket struct {
	CpuNum        int    `mapstructure:"cpuNum"`
	Channel       int    `mapstructure:"channel"`
	Room          int    `mapstructure:"room"`
	SrvProto      int    `mapstructure:"svrProto"`
	RoutineAmount uint64 `mapstructure:"routineAmount"`
	RoutineSize   int    `mapstructure:"routineSize"`
}

type ConnectWebsocket struct {
	ServerId string `mapstructure:"serverId"`
	Bind     string `mapstructure:"bind"`
}

type ConnectTcp struct {
	ServerId      string `mapstructure:"serverId"`
	Bind          string `mapstructure:"bind"`
	SendBuf       int    `mapstructure:"sendbuf"`
	ReceiveBuf    int    `mapstructure:"receivebuf"`
	KeepAlive     bool   `mapstructure:"keepalive"`
	Reader        int    `mapstructure:"reader"`
	ReadBuf       int    `mapstructure:"readBuf"`
	ReadBufSize   int    `mapstructure:"readBufSize"`
	Writer        int    `mapstructure:"writer"`
	WriterBuf     int    `mapstructure:"writerBuf"`
	WriterBufSize int    `mapstructure:"writeBufSize"`
}

type ConnectConfig struct {
	ConnectBase                ConnectBase                `mapstructure:"connect-base"`
	ConnectRpcAddressWebSockts ConnectRpcAddressWebsockts `mapstructure:"connect-rpcAddress-websockts"`
	ConnectRpcAddressTcp       ConnectRpcAddressTcp       `mapstructure:"connect-rpcAddress-tcp"`
	ConnectBucket              ConnectBucket              `mapstructure:"connect-bucket"`
	ConnectWebsocket           ConnectWebsocket           `mapstructure:"connect-websocket"`
	ConnectTcp                 ConnectTcp                 `mapstructure:"connect-tcp"`
}

type LogicBase struct {
	ServerId   string `mapstructure:"serverId"`
	CpuNum     int    `mapstructure:"cpuNum"`
	RpcAddress string `mapstructure:"rpcAddress"`
	CertPath   string `mapstructure:"certPath"`
	KeyPath    string `mapstructure:"keyPath"`
}

type LogicConfig struct {
	LogicBase LogicBase `mapstructure:"logic-base"`
}

type TaskBase struct {
	CpuNum        int    `mapstructure:"cpuNum"`
	RedisAddr     string `mapstructure:"redisAddr"`
	RedisPassword string `mapstructure:"redisPassword"`
	RpcAddress    string `mapstructure:"rpcAddress"`
	PushChan      int    `mapstructure:"pushChan"`
	PushChanSize  int    `mapstructure:"pushChanSize"`
}

type TaskConfig struct {
	TaskBase TaskBase `mapstructure:"task-base"`
}

type ApiBase struct {
	ListenPort int `mapstructure:"listenPort"`
}

type ApiConfig struct {
	ApiBase ApiBase `mapstructure:"api-base"`
}

type SiteBase struct {
	ListenPort int `mapstructure:"listenPort"`
}

type SiteConfig struct {
	SiteBase SiteBase `mapstructure:"site-base"`
}
