package logger

import (
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)

func Logger() *slog.Logger {
	w := os.Stderr
	logger := slog.New(tint.NewHandler(w, nil))
	return logger
}
