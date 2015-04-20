// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dark

import (
	"github.com/yofu/gxui"
	"github.com/yofu/gxui/math"
	"github.com/yofu/gxui/mixins"
)

type CodeEditor struct {
	mixins.CodeEditor
	theme *Theme
}

func CreateCodeEditor(theme *Theme) gxui.CodeEditor {
	t := &CodeEditor{}
	t.theme = theme
	t.Init(t, theme.driver, theme, theme.defaultMonospaceFont)
	t.SetTextColor(theme.TextBoxDefaultStyle.FontColor)
	t.SetMargin(math.Spacing{L: 3, T: 3, R: 3, B: 3})
	t.SetPadding(math.Spacing{L: 3, T: 3, R: 3, B: 3})
	t.SetBorderPen(gxui.TransparentPen)

	return t
}

// mixins.CodeEditor overrides
func (t *CodeEditor) Paint(c gxui.Canvas) {
	t.CodeEditor.Paint(c)

	if t.HasFocus() {
		r := t.Size().Rect()
		c.DrawRoundedRect(r, 3, 3, 3, 3, t.theme.FocusedStyle.Pen, t.theme.FocusedStyle.Brush)
	}
}

func (t *CodeEditor) CreateSuggestionList() gxui.List {
	l := t.theme.CreateList()
	l.SetBackgroundBrush(t.theme.CodeSuggestionListStyle.Brush)
	l.SetBorderPen(t.theme.CodeSuggestionListStyle.Pen)
	return l
}
