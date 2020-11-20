// Lute - 一款对中文语境优化的 Markdown 引擎，支持 Go 和 JavaScript
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package parse

import (
	"bytes"

	"github.com/88250/lute/ast"
	"github.com/88250/lute/lex"
	"github.com/88250/lute/util"
)

func SuperBlockContinue(superBlock *ast.Node, context *Context) int {
	if ok := context.isSuperBlockClose(context.currentLine[context.nextNonspace:]); ok {
		context.finalize(context.Tip) // 最终化所有子块
		return 2
	}
	return 0
}

func (context *Context) superBlockFinalize(superBlock *ast.Node) {
	superBlock.Tokens = bytes.TrimSuffix(superBlock.Tokens, []byte("}}}"))
}

func (t *Tree) parseSuperBlock() (ok bool, layout []byte) {
	marker := t.Context.currentLine[t.Context.nextNonspace]
	if lex.ItemOpenBrace != marker {
		return
	}

	fenceChar := marker
	var fenceLen int
	for i := t.Context.nextNonspace; i < t.Context.currentLineLen && fenceChar == t.Context.currentLine[i]; i++ {
		fenceLen++
	}

	if 3 != fenceLen {
		return
	}

	layout = t.Context.currentLine[t.Context.nextNonspace+fenceLen:]
	layout = lex.TrimWhitespace(layout)
	if !bytes.EqualFold(layout, nil) && !bytes.EqualFold(layout, []byte("row")) && !bytes.EqualFold(layout, []byte("col")) {
		return
	}
	return true, layout
}

func (context *Context) isSuperBlockClose(tokens []byte) (ok bool) {
	if context.Option.KramdownIAL && len("{: id=\"") < len(tokens) {
		// 判断 IAL 打断
		if ial := context.parseKramdownIAL(tokens); 0 < len(ial) {
			context.Tip.KramdownIAL = ial
			context.Tip.InsertAfter(&ast.Node{Type: ast.NodeKramdownBlockIAL, Tokens: tokens})
			return true
		}
	}

	tokens = lex.TrimWhitespace(tokens)
	if !bytes.Equal([]byte("}}}"), tokens) {
		return
	}

	endCaret := bytes.HasSuffix(tokens, util.CaretTokens)
	if context.Option.VditorWYSIWYG || context.Option.VditorIR || context.Option.VditorSV {
		tokens = bytes.ReplaceAll(tokens, util.CaretTokens, nil)
		if endCaret {
			context.Tip.Tokens = bytes.TrimSuffix(context.Tip.Tokens, []byte("\n"))
			context.Tip.Tokens = append(context.Tip.Tokens, util.CaretTokens...)
		}
	}
	return true
}
