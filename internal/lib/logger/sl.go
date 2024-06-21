package sl

import "log/slog"

// красивый вывод ошибки для логгера
func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
