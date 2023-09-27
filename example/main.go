package main

import (
	"errors"

	"github.com/diegostamigni/highlightzap"
	"github.com/highlight/highlight/sdk/highlight-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// create a new Zap logger
	logger, _ := zap.NewProduction()

	// Initialize highlight with your token and optional environment flag
	highlight.SetProjectID("<PROJECT_ID>")
	highlight.Start(
		highlight.WithServiceName("my-app"),
		highlight.WithServiceVersion("git-sha"),
	)
	defer highlight.Stop()

	// create a new core that sends zapcore.ErrorLevel and above messages to Highlight
	highlightCore := highlightzap.NewHighlightCore(zapcore.InfoLevel)

	// Wrap a NewTee to send log messages to both your main logger and to highlight
	logger = logger.WithOptions(zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		return zapcore.NewTee(core, highlightCore)
	}))

	// This message will only go to the main logger
	logger.Info("Highlight Core teed up", zap.String("foo", "bar"))

	// This message will only go to the main logger
	logger.Warn("Warning message with fields", zap.String("foo", "bar"))

	// This error will go to both the main logger and to Highlight. the 'foo' field will appear in highlight as 'custom.foo'
	testError := errors.New("im a test error")
	logger.Error("ran into an error", zap.Error(testError), zap.String("foo", "bar"), zap.Int("some-int", 10))
}
