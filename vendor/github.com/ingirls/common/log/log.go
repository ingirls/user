package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	// the local logger
	logger *zap.SugaredLogger
)

// New New
func New(filePath string) {
	syncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    128, //MB
		MaxAge:     30,
		MaxBackups: 100,
		LocalTime:  true,
		Compress:   true,
	})
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(zap.DebugLevel))
	log := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	logger = log.Sugar()
}

// Debug Debug
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Debugf Debugf
func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args...)
}

// Info Info
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Infof Infof
func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...)
}

// Warn Warn
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Warnf Warnf
func Warnf(template string, args ...interface{}) {
	logger.Warnf(template, args...)
}

// Error Error
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Errorf Errorf
func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args...)
}

// DPanic DPanic
func DPanic(args ...interface{}) {
	logger.DPanic(args...)
}

// DPanicf DPanicf
func DPanicf(template string, args ...interface{}) {
	logger.DPanicf(template, args...)
}

// Panic Panic
func Panic(args ...interface{}) {
	logger.Panic(args...)
}

// Panicf Panicf
func Panicf(template string, args ...interface{}) {
	logger.Panicf(template, args...)
}

// Fatal Fatal
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Fatalf Fatalf
func Fatalf(template string, args ...interface{}) {
	logger.Fatalf(template, args...)
}
