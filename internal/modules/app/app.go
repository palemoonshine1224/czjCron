package app

import (
	"os"
	"path/filepath"

	"fmt"
	"github.com/ouqiang/goutil"
	"github.com/palemoonshine1224/czjCron/internal/modules/logger"
	"github.com/palemoonshine1224/czjCron/internal/modules/setting"
	"github.com/palemoonshine1224/czjCron/internal/modules/utils"
)

var (
	// AppDir 应用根目录
	AppDir string // 应用根目录
	// ConfDir 配置文件目录
	ConfDir string // 配置目录
	// LogDir 日志目录
	LogDir string // 日志目录
	// AppConfig 配置文件
	AppConfig  string // 应用配置文件
	Registered bool   // 是否注册过
	// Setting 应用配置
	Setting *setting.Setting // 应用配置

)

// InitEnv 初始化
func InitEnv() {
	logger.InitLogger()
	var err error
	AppDir, err = goutil.WorkDir()
	if err != nil {
		logger.Fatal(err)
	}
	ConfDir = filepath.Join(AppDir, "/conf")
	LogDir = filepath.Join(AppDir, "/log")
	AppConfig = filepath.Join(ConfDir, "/app.ini")
	createDirIfNotExists(ConfDir, LogDir)
	Registered = IsRegistered()

}

// IsInstalled 判断用户是否已注册
func IsRegistered() bool {
	_, err := os.Stat(filepath.Join(ConfDir, "/install.lock"))
	if os.IsNotExist(err) {
		return false
	}

	return true
}

// CreateInstallLock 创建安装锁文件
func CreateRegisterLock() error {
	_, err := os.Create(filepath.Join(ConfDir, "/install.lock"))
	if err != nil {
		logger.Error("创建安装锁文件conf/install.lock失败")
	}

	return err
}

// 检测目录是否存在
func createDirIfNotExists(path ...string) {
	for _, value := range path {
		if utils.FileExist(value) {
			continue
		}
		err := os.Mkdir(value, 0755)
		if err != nil {
			logger.Fatal(fmt.Sprintf("创建目录失败:%s", err.Error()))
		}
	}
}
