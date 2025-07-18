// Copyright (c) CattleCloud LLC
// SPDX-License-Identifier: BSD-3-Clause

// Package ulog provides a lightweight and efficient leveled logging library for Go.
package ulog

import (
	"fmt"
	"os"
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
	short bool

	E Output // error logs
	W Output // warn logs
	I Output // info logs
	D Output // debug logs
	T Output // trace logs
}

type Output interface {
	Fmt(string, ...any)
}

type format struct {
	prefix string
	name   string
	short  bool
}

func (s *format) Fmt(msg string, args ...any) {
	complete := s.prefix + " [" + s.name + "] " + fmt.Sprintf(msg, args...)
	if !s.short {
		now := time.Now().Format("01/02 15:04:05")
		complete = now + " " + complete
	}
	fmt.Println(complete)
}

type null struct{}

func (null) Fmt(string, ...any) {
	// do nothing
}

func output(name string, short bool, minimum, level int) Output {
	if level > minimum {
		return null{}
	}

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

	return &format{
		prefix: prefix,
		name:   name,
		short:  shorten(short),
	}
}

func shorten(force bool) bool {
	env := os.Getenv("ULOG_NOTIMESTAMP")
	v, _ := strconv.ParseBool(env)
	return v || force
}

func New(name string, opts ...Option) *Log {
	log := &Log{
		name:  name,
		level: Debug,
		short: false,
	}

	for _, opt := range opts {
		opt(log)
	}

	log.E = output(name, log.short, log.level, Error)
	log.W = output(name, log.short, log.level, Warn)
	log.I = output(name, log.short, log.level, Info)
	log.D = output(name, log.short, log.level, Debug)
	log.T = output(name, log.short, log.level, Trace)

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

func SetNoTimestamp() Option {
	return func(l *Log) { l.short = true }
}
