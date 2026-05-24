package logger

type LogFormat string

const (
	CONSOLE LogFormat = "console"
	JSON    LogFormat = "json"
)

type LogLevel string

const (
	DEBUG LogLevel = "debug"
	INFO  LogLevel = "info"
	WARN  LogLevel = "warn"
	ERROR LogLevel = "error"
)

type Option struct {
	Format LogFormat     // 日志格式
	Level  LogLevel      // 日志级别
	Output *OutputConfig // 输出配置
}

type OutputConfig struct {
	EnableFile bool
	FilePath   string
	MaxAge     int
}

func defaultOption() *Option {
	return &Option{
		Format: CONSOLE,
		Level:  "debug",
		Output: &OutputConfig{
			EnableFile: false,
		},
	}
}

func withDefault(opt *Option) *Option {
	def := defaultOption()

	if opt == nil {
		return def
	}

	if opt.Format != "" {
		def.Format = opt.Format
	}

	if opt.Level != "" {
		def.Level = opt.Level
	}

	if opt.Output != nil {
		def.Output.EnableFile = opt.Output.EnableFile
	}

	return def
}
