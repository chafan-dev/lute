// Lute - A structured markdown engine.
// Copyright (C) 2019-present, b3log.org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package lute

import "github.com/b3log/gulu"

// Parse 会将 text 指定的 Markdown 原始文本解析为一颗语法树。
func Parse(name string, text []byte) (t *Tree, err error) {
	defer gulu.Panic.Recover(&err)

	t = &Tree{Name: name, text: text, context: &Context{}}
	t.lex = lex(t.text)
	t.Root = &Document{&BaseNode{typ: NodeRoot}}
	t.parseBlocks()
	t.parseInlines()
	t.lex = nil

	return
}

// Context 用于维护解析过程中使用到的公共数据。
type Context struct {
	linkRefDef map[string]*Link // 链接引用定义集

	// 以下变量用于块级解析阶段

	tip                                                      Node
	oldtip                                                   Node
	currentLine                                              items
	currentLineLen                                           int
	offset, column, nextNonspace, nextNonspaceColumn, indent int
	indented, blank, partiallyConsumedTab, allClosed         bool
	lastMatchedContainer                                     Node

	// 以下变量用于行级解析阶段

	pos        int        // 当前 Token 位置
	delimiters *delimiter // 分隔符栈，用于强调解析
	brackets   *delimiter // 括号栈，用于图片和链接解析
}

func (context *Context) advanceOffset(count int, columns bool) {
	var currentLine = context.currentLine
	var charsToTab, charsToAdvance int
	var c byte
	for ; count > 0; {
		c = currentLine[context.offset]
		if itemTab == c {
			charsToTab = 4 - (context.column % 4)
			if columns {
				context.partiallyConsumedTab = charsToTab > count
				if charsToTab > count {
					charsToAdvance = count
				} else {
					charsToAdvance = charsToTab
				}
				context.column += charsToAdvance
				if !context.partiallyConsumedTab {
					context.offset += 1
				}
				count -= charsToAdvance
			} else {
				context.partiallyConsumedTab = false
				context.column += charsToTab
				context.offset += 1
				count -= 1
			}
		} else {
			context.partiallyConsumedTab = false
			context.offset += 1
			context.column += 1 // assume ascii; block starts are ascii
			count -= 1
		}
	}
}

// advanceNextNonspace 用于预移动到下一个非空字符位置。
func (context *Context) advanceNextNonspace() {
	context.offset = context.nextNonspace
	context.column = context.nextNonspaceColumn
	context.partiallyConsumedTab = false
}

// findNextNonspace 用于查找下一个非空字符。
func (context *Context) findNextNonspace() {
	i := context.offset
	cols := context.column

	var token byte
	for {
		token = context.currentLine[i]
		if itemSpace == token {
			i++
			cols++
		} else if itemTab == token {
			i++
			cols += 4 - (cols % 4)
		} else {
			break
		}
	}

	context.blank = itemNewline == token || itemEnd == token
	context.nextNonspace = i
	context.nextNonspaceColumn = cols
	context.indent = context.nextNonspaceColumn - context.column
	context.indented = context.indent >= 4
}

// Finalize and close any unmatched blocks.
func (context *Context) closeUnmatchedBlocks() {
	if !context.allClosed {
		// finalize any blocks not matched
		for context.oldtip != context.lastMatchedContainer {
			parent := context.oldtip.Parent()
			context.finalize(context.oldtip)
			context.oldtip = parent
		}
		context.allClosed = true
	}
}

// finalize 执行 block 的最终化处理。调用该方法会将 context.tip 置为 block 的父节点。
func (context *Context) finalize(block Node) {
	var parent = block.Parent()
	block.Close()
	block.Finalize(context)
	context.tip = parent
}

// addChild 将 child 作为子节点添加到 context.tip 上。如果 tip 节点不能接受子节点（非块级容器不能添加子节点），则最终化该 tip
// 节点并向父节点方向尝试，直到找到一个能接受 child 的节点为止。
func (context *Context) addChild(child Node) {
	for !context.tip.CanContain(child.Type()) {
		context.finalize(context.tip) // 注意调用 finalize 会向父节点方向进行迭代
	}

	context.tip.AppendChild(context.tip, child)
	context.tip = child
}

// Returns true if the two list items are of the same type,
// with the same delimiter and bullet character.  This is used
// in agglomerating list items into lists.
func (context *Context) listsMatch(list_data, item_data *listData) bool {
	return list_data.typ == item_data.typ &&
		list_data.delimiter == item_data.delimiter &&
		list_data.bulletChar.equal(item_data.bulletChar)
}

// Tree 描述了 Markdown 抽象语法树结构。
type Tree struct {
	Name string    // 名称，可以为空
	Root *Document // 根节点

	text    []byte   // 原始的 Markdown 文本
	lex     *lexer   // 词法分析器
	context *Context // 语法解析上下文
}

// Render 使用 renderer 进行语法树渲染，渲染结果以 output 返回。
func (t *Tree) Render(renderer *Renderer) (output string, err error) {
	err = renderer.Render(t.Root)
	if nil != err {
		return "", err
	}
	output = renderer.writer.String()

	return
}
