// Package ulog provides a lightweight and efficient leveled logging library for Go.
package ulog

import (
	"fmt"
	"strconv"
	"time"
)

const (
	Error = iota
	Warn
	Info
	Debug
	Trace
)

type Log struct {
	name  string
	level int

	E Output // error logs
	W Output // warn logs
	I Output // info logs
	D Output // debug logs
	T Output // trace logs
}

type Output interface {
	Fmt(string, ...any)
}

type standard struct {
	prefix string
	name   string
}

func (s *standard) Fmt(msg string, args ...any) {
	now := time.Now().Format("01/02 15:04:05")
	complete := now + " " + s.prefix + " [" + s.name + "] " + fmt.Sprintf(msg, args...)
	fmt.Println(complete)
}

type null struct{}

func (null) Fmt(string, ...any) {
	// do nothing
}

func output(name string, minimum, level int) Output {
	var prefix string
	switch level {
	case Error:
		prefix = "ERROR"
	case Warn:
		prefix = "WARN "
	case Info:
		prefix = "INFO "
	case Debug:
		prefix = "DEBUG"
	case Trace:
		prefix = "TRACE"
	}

	if level <= minimum {
		return &standard{
			prefix: prefix,
			name:   name,
		}
	}

	return null{}
}

func New(name string, opts ...Option) *Log {
	log := &Log{
		name:  name,
		level: Debug,
	}

	for _, opt := range opts {
		opt(log)
	}

	log.E = output(name, log.level, Error)
	log.W = output(name, log.level, Warn)
	log.I = output(name, log.level, Info)
	log.D = output(name, log.level, Debug)
	log.T = output(name, log.level, Trace)

	return log
}

type Option func(l *Log)

func SetLevel(value int) Option {
	switch value {
	case Error:
	case Warn:
	case Info:
	case Debug:
	case Trace:
	default:
		s := strconv.Itoa(value)
		panic("ulog: not an acceptable level " + s)
	}
	return func(l *Log) { l.level = value }
}
