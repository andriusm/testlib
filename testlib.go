package testlib

import (
	"fmt"
	"log/slog"
)

type TestLib interface {
	DoThings(abc string, num int) string
}

type MyTestLib struct {
	greeting string
	logger   *slog.Logger
}

type Config struct {
	Greeting string
	Logger   *slog.Logger
}

func New(config Config) (*MyTestLib, error) {
	if config.Greeting == "" {
		config.Logger.Debug("TestLib cannot be initialized without greeting")
		return nil, fmt.Errorf("greeting missing in config")
	}

	return &MyTestLib{
		greeting: config.Greeting,
		logger:   config.Logger,
	}, nil
}

func (tlib *MyTestLib) DoThings(abc string, num int) string {
	tlib.logger.Debug("Doing things")

	return fmt.Sprintf("%s, %s! You're #%d", tlib.greeting, abc, num)
}
