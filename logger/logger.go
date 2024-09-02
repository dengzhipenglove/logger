package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var errorLogger *zap.SugaredLogger

type Config struct {
	FileName    string
	LogLevel    string
	ServiceName string
}

func NewLogger(cfg *Config) *zap.Logger {
	hook := lumberjack.Logger{
		Filename:   cfg.FileName, // 日志文件路径，默认 os.TempDir()
		MaxSize:    20,           // 每个日志文件保存20M，默认 100M
		MaxBackups: 30,           // 保留30个备份，默认不限
		MaxAge:     7,            // 保留7天，默认不限
		Compress:   false,        // 是否压缩，默认不压缩
	}
	write := zapcore.AddSync(&hook)
	var level zapcore.Level
	switch cfg.LogLevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)

	core := zapcore.NewCore(
		// zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewJSONEncoder(encoderConfig),
		// zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&write)), // print to consol and file
		write,
		level,
	)

	var logger *zap.Logger
	fields := zap.Fields(zap.String("serviceName", cfg.ServiceName))
	if len(cfg.ServiceName) > 0 {
		logger = zap.New(core, zap.AddCaller(), zap.Development(), fields)
	} else {
		logger = zap.New(core, zap.AddCaller(), zap.Development(), fields)
	}

	return logger
}

func init() {
	log := NewLogger(&Config{
		FileName: "./all.log",
		LogLevel: "debug",
	})
	errorLogger = log.Sugar()
}

func Sync() {
	errorLogger.Sync()
}

// msg
func Debug(args ...interface{}) {
	errorLogger.Debug(args...)
}

func Info(args ...interface{}) {
	errorLogger.Info(args...)
}

func Warn(args ...interface{}) {
	errorLogger.Warn(args...)
}

func Error(args ...interface{}) {
	errorLogger.Error(args...)
}

// format
func Debugf(template string, args ...interface{}) {
	errorLogger.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	errorLogger.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	errorLogger.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	errorLogger.Errorf(template, args...)
}

// key : value
func Debugw(msg string, keysAndValues ...interface{}) {
	errorLogger.Debugw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	errorLogger.Infow(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	errorLogger.Warnw(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	errorLogger.Errorw(msg, keysAndValues...)
}
