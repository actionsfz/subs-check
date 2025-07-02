package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/beck-8/subs-check/app"
)

var (
	Version       = "dev"
	CurrentCommit = "none"
)

func main() {
	// 创建 app 实例
	application := app.New(fmt.Sprintf("%s-%s", Version, CurrentCommit))
	slog.Info(fmt.Sprintf("当前版本: %s-%s", Version, CurrentCommit))

	// 初始化（加载配置等）
	if err := application.Initialize(); err != nil {
		slog.Error(fmt.Sprintf("初始化失败: %v", err))
		os.Exit(1)
	}

	// ✅ 核心改动：只执行一次 CheckSub
	application.CheckSub()

	slog.Info("订阅检测已完成，一次性任务退出。")
}
