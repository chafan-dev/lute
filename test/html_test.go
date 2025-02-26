// Lute - 一款结构化的 Markdown 引擎，支持 Go 和 JavaScript
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package test

import (
	"strings"
	"testing"

	"github.com/chafan-dev/lute/html"
	"github.com/chafan-dev/lute/html/atom"
)

func TestHTMLParse(t *testing.T) {
	reader := strings.NewReader("<p>foo</p>")
	htmlRoot := &html.Node{Type: html.ElementNode}
	htmlNodes, err := html.ParseFragment(reader, htmlRoot)
	if nil != err || 1 != len(htmlNodes) {
		t.Fatal("test HTML parse failed")
	}

	if atom.P != htmlNodes[0].DataAtom {
		t.Fatal("test HTML parse failed")
	}
}

func TestHTMLParse0(t *testing.T) {
	reader := strings.NewReader("<p>&& &nbsp;</p>")
	htmlRoot := &html.Node{Type: html.ElementNode}
	htmlNodes, err := html.ParseFragment(reader, htmlRoot)
	if nil != err || 1 != len(htmlNodes) {
		t.Fatal("test HTML parse failed")
	}

	if atom.P != htmlNodes[0].DataAtom {
		t.Fatal("test HTML parse failed")
	}
}
