package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net"
	"os"
	"path"
)

type logLevel string

const (
	localLogLevel logLevel = "local"
	stageLogLevel logLevel = "stage"
	prodLogLevel  logLevel = "prod"
)

type Logger interface {
	DebugContext(ctx context.Context, msg string, args ...any)
	InfoContext(ctx context.Context, msg string, args ...any)
	WarnContext(ctx context.Context, msg string, args ...any)
	ErrorContext(ctx context.Context, msg string, args ...any)
}

type Config struct {
	Path    string      `yaml:"path"`
	Level   logLevel    `yaml:"level"`
	Source  bool        `yaml:"source"`
	Graylog GraylogConf `yaml:"graylog"`
}

type GraylogConf struct {
	Use      bool   `yaml:"use"`
	ConnType string `yaml:"conn_type"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}

func NewLogger(cfg *Config) (Logger, error) {
	handlerOptions := &slog.HandlerOptions{
		AddSource: cfg.Source,
	}

	var (
		output io.Writer
		err    error
	)

	switch cfg.Level {
	case localLogLevel:
		handlerOptions.Level = slog.LevelDebug
		output = os.Stdout
	case stageLogLevel:
		handlerOptions.Level = slog.LevelInfo
		output, err = loadFileOutput(cfg)
	case prodLogLevel:
		handlerOptions.Level = slog.LevelError
		output, err = loadFileOutput(cfg)
	}

	if err != nil {
		panic(err)
	}

	if cfg.Graylog.Use {
		output = connectToGraylog(cfg)
	}

	log := slog.New(slog.NewJSONHandler(output, handlerOptions))
	return log, err
}

func loadFileOutput(cfg *Config) (logFile *os.File, err error) {
	if err = os.MkdirAll(path.Dir(cfg.Path), 0755); err != nil {
		return
	}

	logFile, err = os.OpenFile(cfg.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	return
}

func connectToGraylog(cfg *Config) net.Conn {
	conn, err := net.Dial(cfg.Graylog.ConnType, fmt.Sprintf("%s:%s", cfg.Graylog.Host, cfg.Graylog.Port))

	if err != nil {
		panic(err)
	}

	return conn
}
