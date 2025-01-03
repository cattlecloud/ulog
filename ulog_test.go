// Copyright (c) CattleCloud LLC
// SPDX-License-Identifier: BSD-3-Clause

package ulog

import "testing"

func Test_New(t *testing.T) {
	t.Parallel()

	log := New("test")

	log.E.Fmt("this is %s", "error")
	log.W.Fmt("this is %s", "warn")
	log.I.Fmt("this is %s", "info")
	log.D.Fmt("this is %s", "debug")
	log.T.Fmt("should not see %s", "trace")

	t.Log("complete")
}

func Test_SetLevel(t *testing.T) {
	t.Parallel()

	t.Run("error", func(_ *testing.T) {
		log := New("test-error", SetLevel(Error))
		log.E.Fmt("this is error")
		log.W.Fmt("should not see warn")
		log.I.Fmt("should not see info")
		log.D.Fmt("should not see debug")
		log.T.Fmt("should not see trace")
	})

	t.Run("warn", func(_ *testing.T) {
		log := New("test-warn", SetLevel(Warn))
		log.E.Fmt("this is error")
		log.W.Fmt("this is warn")
		log.I.Fmt("should not see info")
		log.D.Fmt("should not see debug")
		log.T.Fmt("should not see trace")
	})

	t.Run("info", func(_ *testing.T) {
		log := New("test-info", SetLevel(Info))
		log.E.Fmt("this is error")
		log.W.Fmt("this is warn")
		log.I.Fmt("this is info")
		log.D.Fmt("should not see debug")
		log.T.Fmt("should not see trace")
	})

	t.Run("debug", func(_ *testing.T) {
		log := New("test-debug", SetLevel(Debug))
		log.E.Fmt("this is error")
		log.W.Fmt("this is warn")
		log.I.Fmt("this is info")
		log.D.Fmt("this is debug")
		log.T.Fmt("should not see trace")
	})

	t.Run("trace", func(_ *testing.T) {
		log := New("test-trace", SetLevel(Trace))
		log.E.Fmt("this is error")
		log.W.Fmt("this is warn")
		log.I.Fmt("this is info")
		log.D.Fmt("this is debug")
		log.T.Fmt("this is trace")
	})

	t.Log("complete")
}
