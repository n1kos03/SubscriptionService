package pkg

import (
	"log/slog"
	"os"
	"time"
)

func ParseDate(t string) (time.Time, error) {
	date, err := time.Parse("01-2006", t)
	if err != nil {
		return time.Time{}, err
	}

	return date, nil
}

func SetupLogger() *slog.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	slog.SetDefault(logger)

	return logger
}
