// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dark

import (
	"github.com/yofu/gxui"
	"github.com/yofu/gxui/mixins"
)

func CreateLinearLayout(theme *Theme) gxui.LinearLayout {
	l := &mixins.LinearLayout{}
	l.Init(l, theme)
	return l
}
