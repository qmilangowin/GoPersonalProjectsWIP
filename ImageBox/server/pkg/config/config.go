// package is used for custom configurations, including logging, text colour, etc.

package config

import (
	"io"

	"github.com/qmilangowin/imagebox/pkg/logging"
)

var LogDebug bool
var LogWriter io.Writer

type textFormat struct {
	Color string
	Reset string
}

type loggingConfig struct {
	Log *logging.Logger
}

type TextFormatConfig interface {
	SetTextFormatting(colour string) *textFormat
}

type LoggingConfig interface {
	SetLogging(w io.Writer, isDebug bool) *loggingConfig
}

func NewTextConfig() TextFormatConfig {
	return &textFormat{}
}

func NewLoggingConfig() LoggingConfig {
	return &loggingConfig{}
}

func (lc *loggingConfig) SetLogging(w io.Writer, isDebug bool) *loggingConfig {

	newLog := logging.New(w, isDebug)
	l := loggingConfig{
		Log: newLog,
	}
	LogDebug = isDebug
	LogWriter = w
	return &l
}

func (t *textFormat) SetTextFormatting(colour string) *textFormat {

	//default is set to white
	tf := textFormat{
		Color: "\033[37m",
		Reset: "\033[0m",
	}

	switch colour {
	case "red":
		tf.Color = "\033[31m"
	case "green":
		tf.Color = "\033[32m"
	case "yellow":
		tf.Color = "\033[33m"
	case "blue":
		tf.Color = "\033[34m"
	case "purple":
		tf.Color = "\033[35m"
	case "cyan":
		tf.Color = "\033[36m"
	case "white":
		tf.Color = "\033[37m"
	default:
		tf.Color = "\033[37m"
	}

	return &tf
}
