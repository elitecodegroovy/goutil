package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"sync"

	"os"
	"path/filepath"
)

var (
	inited = false
	l      *Logger

	sp                             = string(filepath.Separator)
	errWS, warnWS, infoWS, debugWS zapcore.WriteSyncer       // IO输出
	debugConsoleWS                 = zapcore.Lock(os.Stdout) // 控制台标准输出
	errorConsoleWS                 = zapcore.Lock(os.Stderr)

	defaultDevelopmentEnv   = "dev"
	defaultOutputPaths      = "stdout"
	defaultErrorOutputPaths = "stderr"
)

type Logger struct {
	*zap.Logger
	sync.RWMutex
	Opts      *Options `json:"opts"`
	zapConfig zap.Config
	once      sync.Once
}

func GetLogger() *Logger {
	if !inited {
		l = &Logger{
			Opts: &Options{},
		}

		l.Lock()
		defer l.Unlock()

		l.once.Do(func() {
			l.loadCfg()
			l.init()
			l.Info("[initLogger] zap plugin initializing completed")
			inited = true
		})
	}

	return l
}

func (l *Logger) loadCfg() {

	if os.Getenv("LOGGER_ENV") == defaultDevelopmentEnv {
		l.zapConfig = zap.NewDevelopmentConfig()
	} else {
		l.zapConfig = zap.NewProductionConfig()
	}

	// application log output path
	if len(os.Getenv("LOGGER_OUTPUT_PATHS")) != 0 {
		l.zapConfig.OutputPaths = []string{os.Getenv("LOGGER_OUTPUT_PATHS")}
	} else {
		l.zapConfig.OutputPaths = []string{defaultOutputPaths}
	}

	//  error of zap-self log
	if len(os.Getenv("LOGGER_ERROR_OUTPUT_PATHS")) != 0 {
		l.zapConfig.OutputPaths = []string{os.Getenv("LOGGER_ERROR_OUTPUT_PATHS")}
	} else {
		l.zapConfig.OutputPaths = []string{defaultErrorOutputPaths}
	}

	// 默认输出到程序运行目录的logs子目录
	if len(os.Getenv("LOGGER_LOG_FILE_DIR")) == 0 {
		l.Opts.LogFileDir, _ = filepath.Abs(filepath.Dir(filepath.Join(".")))
		l.Opts.LogFileDir += sp + "logs" + sp
	}

	if l.Opts.AppName == "" {
		l.Opts.AppName = "app"
	}

	if l.Opts.ErrorFileName == "" {
		l.Opts.ErrorFileName = "error.log"
	}

	if l.Opts.WarnFileName == "" {
		l.Opts.WarnFileName = "warn.log"
	}

	if l.Opts.InfoFileName == "" {
		l.Opts.InfoFileName = "info.log"
	}

	if l.Opts.DebugFileName == "" {
		l.Opts.DebugFileName = "debug.log"
	}

	if l.Opts.MaxSize == 0 {
		l.Opts.MaxSize = 50
	}
	if l.Opts.MaxBackups == 0 {
		l.Opts.MaxBackups = 3
	}
	if l.Opts.MaxAge == 0 {
		l.Opts.MaxAge = 30
	}
}

func (l *Logger) init() {

	l.setSyncers()
	var err error

	l.Logger, err = l.zapConfig.Build(l.cores())
	if err != nil {
		panic(err)
	}

	defer l.Logger.Sync()
}

func (l *Logger) setSyncers() {
	f := func(fN string) zapcore.WriteSyncer {
		return zapcore.AddSync(&lumberjack.Logger{
			Filename:   l.Opts.LogFileDir + sp + l.Opts.AppName + "-" + fN,
			MaxSize:    l.Opts.MaxSize,
			MaxBackups: l.Opts.MaxBackups,
			MaxAge:     l.Opts.MaxAge,
			Compress:   true,
			LocalTime:  true,
		})
	}

	errWS = f(l.Opts.ErrorFileName)
	warnWS = f(l.Opts.WarnFileName)
	infoWS = f(l.Opts.InfoFileName)
	debugWS = f(l.Opts.DebugFileName)

	return
}

func (l *Logger) cores() zap.Option {
	fileEncoder := getEncoder(false)
	consoleEncoder := getEncoder(true)

	errPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl > zapcore.WarnLevel && zapcore.WarnLevel-l.zapConfig.Level.Level() > -1
	})
	warnPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.WarnLevel && zapcore.WarnLevel-l.zapConfig.Level.Level() > -1
	})
	infoPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel && zapcore.InfoLevel-l.zapConfig.Level.Level() > -1
	})
	debugPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.DebugLevel && zapcore.DebugLevel-l.zapConfig.Level.Level() > -1
	})

	cores := []zapcore.Core{
		// region 日志文件
		// error 及以上
		zapcore.NewCore(fileEncoder, errWS, errPriority),
		// warn
		zapcore.NewCore(fileEncoder, warnWS, warnPriority),
		// info
		zapcore.NewCore(fileEncoder, infoWS, infoPriority),
		// debug
		zapcore.NewCore(fileEncoder, debugWS, debugPriority),

		// endregion
		// region 控制台

		// 错误及以上
		zapcore.NewCore(consoleEncoder, errorConsoleWS, errPriority),
		// 警告
		zapcore.NewCore(consoleEncoder, debugConsoleWS, warnPriority),
		// info
		zapcore.NewCore(consoleEncoder, debugConsoleWS, infoPriority),
		// debug
		zapcore.NewCore(consoleEncoder, debugConsoleWS, debugPriority),
		// endregion
	}
	return zap.WrapCore(func(c zapcore.Core) zapcore.Core {
		return zapcore.NewTee(cores...)
	})
}

func getEncoder(console bool) zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	//encoderConfig := zap.NewProductionEncoderConfig()
	//encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	if console {
		return zapcore.NewConsoleEncoder(encoderConfig)
	}
	return zapcore.NewJSONEncoder(encoderConfig)

}
