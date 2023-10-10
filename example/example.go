// Code generated by go-simpler.org/sloggen. DO NOT EDIT.

package example

import "context"
import "fmt"
import "log/slog"
import "runtime"
import "strings"
import "time"

const LevelInfo = slog.Level(0)
const LevelAlert = slog.Level(1)

const RequestId = "request_id"

func CreatedAt(value time.Time) slog.Attr { return slog.Time("created_at", value) }
func Err(value error) slog.Attr           { return slog.Any("err", value) }
func UserId(value int) slog.Attr          { return slog.Int("user_id", value) }

func ParseLevel(s string) (slog.Level, error) {
	switch strings.ToUpper(s) {
	case "INFO":
		return LevelInfo, nil
	case "ALERT":
		return LevelAlert, nil
	default:
		return 0, fmt.Errorf("slog: level string %q: unknown name", s)
	}
}

func RenameLevels(_ []string, attr slog.Attr) slog.Attr {
	if attr.Key != slog.LevelKey {
		return attr
	}
	switch attr.Value.Any().(slog.Level) {
	case LevelInfo:
		attr.Value = slog.StringValue("INFO")
	case LevelAlert:
		attr.Value = slog.StringValue("ALERT")
	}
	return attr
}

type Logger struct{ handler slog.Handler }

func New(h slog.Handler) *Logger { return &Logger{handler: h} }

func (l *Logger) Handler() slog.Handler { return l.handler }

func (l *Logger) Enabled(ctx context.Context, level slog.Level) bool {
	return l.handler.Enabled(ctx, level)
}

func (l *Logger) With(attrs ...slog.Attr) *Logger {
	if len(attrs) == 0 {
		return l
	}
	return &Logger{handler: l.handler.WithAttrs(attrs)}
}

func (l *Logger) WithGroup(name string) *Logger {
	if name == "" {
		return l
	}
	return &Logger{handler: l.handler.WithGroup(name)}
}

func (l *Logger) Log(ctx context.Context, level slog.Level, msg string, attrs ...slog.Attr) {
	l.log(ctx, level, msg, attrs)
}

func (l *Logger) Info(ctx context.Context, msg string, attrs ...slog.Attr) {
	l.log(ctx, LevelInfo, msg, attrs)
}

func (l *Logger) Alert(ctx context.Context, msg string, attrs ...slog.Attr) {
	l.log(ctx, LevelAlert, msg, attrs)
}

func (l *Logger) log(ctx context.Context, level slog.Level, msg string, attrs []slog.Attr) {
	if !l.handler.Enabled(ctx, level) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(3, pcs[:])
	r := slog.NewRecord(time.Now(), level, msg, pcs[0])
	r.AddAttrs(attrs...)
	_ = l.handler.Handle(ctx, r)
}
