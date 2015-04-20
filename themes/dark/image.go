// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dark

import (
	"github.com/yofu/gxui"
	"github.com/yofu/gxui/mixins"
)

func CreateImage(theme *Theme) gxui.Image {
	i := &mixins.Image{}
	i.Init(i, theme)
	return i
}
