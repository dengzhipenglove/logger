package main

import (
	"logtest/logger"
	"time"

	"go.uber.org/zap"
)

func main() {
	// 历史记录日志名字为：all.log，服务重新启动，日志会追加，不会删除
	log := logger.NewLogger(&logger.Config{
		FileName: "./all.log",
		LogLevel: "debug",
	})

	log.Debug("dfds")
	// 强结构形式
	log.Info("test",
		zap.String("string", "string"),
		zap.Int("int", 3),
		zap.Duration("time", time.Second),
	)
	// 必须 key-value 结构形式 性能下降一点
	log.Sugar().Infow("test-",
		"string", "string",
		"int", 1,
		"time", time.Second,
	)
	logger.Debug("----fdfdsfd")
}

// =========================================================================
