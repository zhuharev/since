// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package since

// Translator represents interface of translator
type Translator interface {
	Tr(key string, args ...interface{})
}

var (
	// NowKey key for locale
	NowKey = []string{"now"}
	// SecondKey key for locale
	SecondKey = []string{"second", "seconds"}
	// MinuteKey key for locale
	MinuteKey = []string{"minute", "minutes"}
	// HourKey key for locale
	HourKey = []string{"hour", "hours"}
	// DayKey key for locale
	DayKey = []string{"day", "days"}
)
