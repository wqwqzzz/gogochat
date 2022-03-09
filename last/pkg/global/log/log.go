package log
import (
   // "github.com/gookit/color"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gochat/internal/kafka"
	"net/url"
)
type Color struct {
}
func (c Color) Sync() error {
    // 因为只是控制台打印, 涉及不到缓存同步问题, 所以简单return就可以
	return nil
}
func (c Color) Close() error {
    // 因为只是控制台打印, 也涉及不到关闭对象的问题, return就好
	return nil
}
func (c Color) Write(p []byte) (n int, err error) {
    // 使用带颜色的控制台输出日志信息
	kafka.SendLog(string(p))
	// 返回写入日志的长度,以及错误
	return len(p), nil
}
func colorSink(url *url.URL) (sink zap.Sink, err error) {
    // 工厂函数中, 定义了必须接收一个*url.URL参数
    // 但是我们的需求比较简单, 暂时用不到, 所以可以直接忽略这个参数的使用
    
    // 实例化一个Color对象, 该对象实现了Sink接口    
	c := Color{}
	return c, nil
}
var Logger *zap.Logger
var Any =zap.Any
var	String  = zap.String
var	Int     = zap.Int
var	Float32 = zap.Float32
func init(){
        // 将colorSink工厂函数注册到zap中, 自定义协议名为 Color
	if err := zap.RegisterSink("Color", colorSink); err != nil {
		return
	}

	logLevel := zap.NewAtomicLevelAt(zapcore.DebugLevel)
	var zc = zap.Config{
		Level:             logLevel,
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "message",
			LevelKey:       "level",
			TimeKey:        "time",
			NameKey:        "name",
			CallerKey:      "caller",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeName:     zapcore.FullNameEncoder,
		},
		// 日志标准输出的定义中, 除了标准的控制台输出, 还增加了一个我们自定义的Color协议输出
		// 这里需要注意的是, 我们的自定义协议中, 固定是接收了一个 *url.URL, 虽然我们没有用到
		// 但是在日志实际配置使用时, 我们仍需要显示传递该参数. 按照http协议的风格, 我们可以
		// 将其定义为 "Color://127.0.0.1", 当然 "Color:127.0.0.1" 和 "Color:" 
		// 这种形式也是可以的. 但是 "Color" 这种是错误的配置形式
		OutputPaths:      []string{"Color://127.0.0.1"},
		ErrorOutputPaths: []string{"stderr"},
		InitialFields:    map[string]interface{}{"app": "apdex"},
	}

	Logger, _= zc.Build()

	defer Logger.Sync()


    
}

