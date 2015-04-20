// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dark

import (
	"github.com/yofu/gxui"
	"github.com/yofu/gxui/math"
	"github.com/yofu/gxui/mixins"
)

func CreateLabel(theme *Theme) gxui.Label {
	l := &mixins.Label{}
	l.Init(l, theme, theme.defaultFont, theme.LabelStyle.FontColor)
	l.SetMargin(math.Spacing{L: 3, T: 3, R: 3, B: 3})
	return l
}
