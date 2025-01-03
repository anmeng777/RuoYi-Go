// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package logger

import (
	"RuoYi-Go/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

// NewZapLogger
func NewZapLogger(c config.AppConfig) *zap.Logger {
	// 当前时间
	now := time.Now()

	// 格式化为 "2006-01-02" 的形式
	formatted := now.Format("2006-01-02")

	// lumberjack配置
	lumberjackLogger := &lumberjack.Logger{
		Filename:   c.Log.LogPath + "/" + c.App.AppName + "_" + formatted + ".log", // 日志文件路径
		MaxSize:    100,                                                            // 单个日志文件最大大小（单位：MB）
		MaxBackups: 3,                                                              // 保留旧文件的最大数量
		MaxAge:     28,                                                             // 旧文件保留最大天数
		Compress:   true,                                                           // 是否压缩旧文件
	}

	// 自定义zap的encoder配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// 构建zap.Core
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // 使用JSON格式编码日志
		zapcore.NewMultiWriteSyncer( // 同时写入多个地方：控制台和文件
			zapcore.AddSync(os.Stdout),        // 输出到控制台
			zapcore.AddSync(lumberjackLogger), // 输出到文件，使用lumberjack进行日志分割管理
		),
		zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= c.Log.LogLevel // 设置日志级别
		}),
	)

	var zaplogger *zap.Logger
	switch c.Log.LogLevel {
	case zap.ErrorLevel:
		zaplogger = zap.New(core, zap.Development(), zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
		break
	case zap.DebugLevel:
		zaplogger = zap.New(core, zap.Development(), zap.AddCaller(), zap.AddStacktrace(zap.DebugLevel))
		break
	default:
		zaplogger = zap.New(core, zap.AddCaller())
	}

	return zaplogger
}
