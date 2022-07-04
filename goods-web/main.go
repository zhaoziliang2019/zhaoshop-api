package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/inner/uuid"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"zhaoshop-api/commons"
	"zhaoshop-api/goods-web/global"
	"zhaoshop-api/goods-web/initialize"
	"zhaoshop-api/goods-web/utils"
	"zhaoshop-api/goods-web/utils/register/consul"
)

func main() {
	//1.初始化logger
	initialize.InitLogger()
	//2、初始化config
	initialize.InitConfig()
	//2.1自动获取ip地址
	global.ServerConfig.Host, _ = commons.GetHost()
	//3、初始化routers
	router := initialize.Routers()
	//4.初始化翻译
	err := initialize.InitTrans("zh")
	if err != nil {
		return
	}
	//5、注册用户服务consul
	initialize.InitSrvConn()
	//6、初始化sentinel
	initialize.InitSentinel()
	//自动获取服务端口号
	debug := viper.GetBool("ZHAOSHOP_DEBUG")
	if debug {
		port, err := utils.GetFreePort()
		if err == nil {
			global.ServerConfig.Port = port
		}
	}
	/*
		1、S()可以获取一个全局的sugar，可以让我们自己设置一个全局的logger
		2、日志是分级别的，debug,info,warn,error,fetal
		3、S函数和L函数很有用，提供一个全局的安全访问logger的途径
	*/
	register_client := consul.NewRegistry(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	serviceId, _ := uuid.NewV4()
	serviceIdStr := fmt.Sprintf("%s", serviceId)
	err = register_client.Register(global.ServerConfig.Host, global.ServerConfig.Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceIdStr)
	if err != nil {
		zap.S().Panic("服务注册失败：", err.Error())
	}
	zap.S().Infof("启动服务器，端口：%d", global.ServerConfig.Port)
	go func() {
		if err := router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
			zap.S().Panic("启动失败：", err.Error())
		}
	}()
	//接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	err = register_client.DeRegister(serviceIdStr)
	if err != nil {
		zap.S().Panic("注销失败：", err.Error())
	} else {
		zap.S().Info("注销成功：")
	}
}
