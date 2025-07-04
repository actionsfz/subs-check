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
	application := app.New(fmt.Sprintf("%s-%s", Version, CurrentCommit))
	slog.Info(fmt.Sprintf("当前版本: %s-%s", Version, CurrentCommit))

	if err := application.Initialize(); err != nil {
		slog.Error(fmt.Sprintf("初始化失败: %v", err))
		os.Exit(1)
	}

	application.OnceCheck()

	slog.Info("订阅检测已完成，一次性任务退出。")
}
