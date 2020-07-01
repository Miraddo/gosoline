package env

import (
	"github.com/applike/gosoline/pkg/cfg"
	"github.com/applike/gosoline/pkg/mon"
)

type Option func(env *Environment)
type ConfigOption func(config cfg.GosoConf) error
type LoggerOption func(config cfg.GosoConf, logger mon.GosoLog) error

type loggerSettings struct {
	Level           string `cfg:"level" default:"info" validate:"required"`
	Format          string `cfg:"format" default:"console" validate:"required"`
	TimestampFormat string `cfg:"timestamp_format" default:"15:04:05.000" validate:"required"`
}

func WithConfigFile(file string) Option {
	return func(env *Environment) {
		env.addConfigOption(func(config cfg.GosoConf) error {
			return config.Option(cfg.WithConfigFile(file, "yml"))
		})
	}
}

func WithLoggerLevel(level string) Option {
	return func(env *Environment) {
		env.addLoggerOption(func(_ cfg.GosoConf, logger mon.GosoLog) error {
			return logger.Option(mon.WithLevel(level))
		})
	}
}

func WithLoggerSettingsFromConfig(env *Environment) {
	env.addLoggerOption(func(config cfg.GosoConf, logger mon.GosoLog) error {
		settings := &loggerSettings{}
		config.UnmarshalKey("test.logger", settings)

		loggerOptions := []mon.LoggerOption{
			mon.WithLevel(settings.Level),
			mon.WithFormat(settings.Format),
			mon.WithTimestampFormat(settings.TimestampFormat),
		}

		return logger.Option(loggerOptions...)
	})
}