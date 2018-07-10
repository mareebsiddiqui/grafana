package plugins

import (
	"log"

	glog "github.com/mareebsiddiqui/grafana/pkg/log"
	hclog "github.com/hashicorp/go-hclog"
)

type LogWrapper struct {
	Logger glog.Logger
}

func (lw LogWrapper) Trace(msg string, args ...interface{}) {
	lw.Logger.Debug(msg, args...)
}
func (lw LogWrapper) Debug(msg string, args ...interface{}) {
	lw.Logger.Debug(msg, args...)
}
func (lw LogWrapper) Info(msg string, args ...interface{}) {
	lw.Logger.Info(msg, args...)
}
func (lw LogWrapper) Warn(msg string, args ...interface{}) {
	lw.Logger.Warn(msg, args...)
}
func (lw LogWrapper) Error(msg string, args ...interface{}) {
	lw.Logger.Error(msg, args...)
}

func (lw LogWrapper) IsTrace() bool { return true }
func (lw LogWrapper) IsDebug() bool { return true }
func (lw LogWrapper) IsInfo() bool  { return true }
func (lw LogWrapper) IsWarn() bool  { return true }
func (lw LogWrapper) IsError() bool { return true }

func (lw LogWrapper) With(args ...interface{}) hclog.Logger {
	return LogWrapper{Logger: lw.Logger.New(args...)}
}
func (lw LogWrapper) Named(name string) hclog.Logger {
	return LogWrapper{Logger: lw.Logger.New()}
}

func (lw LogWrapper) ResetNamed(name string) hclog.Logger {
	return LogWrapper{Logger: lw.Logger.New()}
}

func (lw LogWrapper) StandardLogger(ops *hclog.StandardLoggerOptions) *log.Logger {
	return nil
}
