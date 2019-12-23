// Lute - A structured markdown engine.
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under the Mulan PSL v1.
// You can use this software according to the terms and conditions of the Mulan PSL v1.
// You may obtain a copy of Mulan PSL v1 at:
//     http://license.coscl.org.cn/MulanPSL
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR
// PURPOSE.
// See the Mulan PSL v1 for more details.

// +build javascript

package test

import (
	"testing"

	"github.com/88250/lute"
)

var vditorDOM2MdTests = []parseTest{

	{"53", "<blockquote><p><br></p><p><wbr>foo\n</p></blockquote>", "> foo\n"}, // 在块引用第一个字符前换行
	{"52", "<blockquote><p>foo\n</p><blockquote><p>bar<wbr>\n</p></blockquote></blockquote>", "> foo\n>\n> > bar\n> >\n"},
	{"51", "<blockquote><blockquote><p><wbr>\n</p></blockquote></blockquote>", "\n"},
	{"50", "<blockquote><p><wbr>\n</p></blockquote>", "\n"},
	{"49", "<blockquote><p>f<wbr>\n</p></blockquote>", "> f\n"},
	{"48", "<div class=\"vditor-wysiwyg__block\" data-type=\"math-block\"><pre><code data-code=\"foo\"></code></pre></div>", "$$\nfoo\n$$\n"},
	{"47", "<p><em data-marker=\"*\"><br></em></p><p><em data-marker=\"*\"><wbr>foo</em></p>", "*foo*\n"},
	{"46", "<p><em data-marker=\"*\">foo<wbr></em></p><p><em data-marker=\"*\"></em></p>", "*foo*\n"},
	{"45", "<p><em data-marker=\"*\">foo</em></p><p><em data-marker=\"*\"><wbr><br></em></p>", "*foo*\n"},
	{"44", "<ol><li data-marker=\"1.\"><p>Node.js</p></li><li data-marker=\"2.\"><p>Go<wbr></p></li></ol>", "1. Node.js\n2. Go\n"},
	{"43", "<p>f<i>o</i>o<wbr></p>", "f*o*o\n"},
	{"42", "<ul data-tight=\"true\"><li data-marker=\"*\">foo<br></li><ul><li data-marker=\"*\">b<wbr></li></ul></ul>", "* foo\n  * b\n"},
	{"41", "<div class=\"vditor-wysiwyg__block\" data-type=\"code-block\"><pre><code data-code=\"foo\" class=\"language-go\">foo<br></code></pre></div>", "```go\nfoo\n```\n"},
	{"40", "<p>f<span data-marker=\"*\">o</span>ob<wbr></p>", "foob\n"},
	{"39", "<p><b>foo<wbr></b></p>", "**foo**\n"},
	{"38", "<p>```java</p><p><wbr><br></p>", "```java\n"},
	{"37", "<ul data-tight=\"true\"><li data-marker=\"*\">foo<wbr></li><li data-marker=\"*\"></li><li data-marker=\"*\"><br></li></ul>", "* foo\n*\n*\n"},
	{"36", "<ul data-tight=\"true\"><li data-marker=\"*\">1<em data-marker=\"*\">2</em></li><li data-marker=\"*\"><em data-marker=\"*\"><wbr><br></em></li></ul>", "* 1*2*\n*\n"},
	{"35", "<ul data-tight=\"true\"><li data-marker=\"*\"><wbr><br></li></ul>", "*\n"},
	{"34", "<p>中<wbr>文</p>", "中文\n"},
	{"33", "<ol data-tight=\"true\"><li data-marker=\"1.\">foo</li></ul>", "1. foo\n"},
	{"32", "<ul data-tight=\"true\"><li data-marker=\"*\">foo<wbr></li></ul>", "* foo\n"},
	{"31", "<ul><li data-marker=\"*\">foo<ul><li data-marker=\"*\">bar</li></ul></li></ul>", "* foo\n  * bar\n"},
	{"30", "<ul><li data-marker=\"*\">foo</li><li data-marker=\"*\"><ul><li data-marker=\"*\"><br /></li></ul></li></ul>", "* foo\n* *\n"},
	{"29", "<p><s>del</s></p>", "~~del~~\n"},
	{"29", "<p>[]()</p>", "[]()\n"},
	{"28", ":octocat:", ":octocat:\n"},
	{"27", "<table><thead><tr><th>abc</th><th>def</th></tr></thead></table>\n", "|abc|def|\n|---|---|\n"},
	{"26", "<p><del data-marker=\"~~\">Hi</del> Hello, world!</p>", "~~Hi~~ Hello, world!\n"},
	{"25", "<p><del data-marker=\"~\">Hi</del> Hello, world!</p>", "~Hi~ Hello, world!\n"},
	{"24", "<ul><li data-marker=\"*\" class=\"vditor-task\"><input checked=\"\" disabled=\"\" type=\"checkbox\" /> foo<wbr></li></ul>", "* [X] foo\n"},
	{"23", "<ul><li data-marker=\"*\" class=\"vditor-task\"><input disabled=\"\" type=\"checkbox\" /> foo<wbr></li></ul>", "* [ ] foo\n"},
	{"22", "><wbr>", ">\n"},
	{"21", "<p>> foo<wbr></p>", "> foo\n"},
	{"20", "<p>foo</p><p><wbr><br></p>", "foo\n"},
	{"19", "<ul><li data-marker=\"*\">foo</li></ul>", "* foo\n"},
	{"18", "<p><em data-marker=\"*\">foo<wbr></em></p>", "*foo*\n"},
	{"17", "foo bar", "foo bar\n"},
	{"16", "<p><em><strong>foo</strong></em></p>", "***foo***\n"},
	{"15", "<p><strong data-marker=\"__\">foo</strong></p>", "__foo__\n"},
	{"14", "<p><strong data-marker=\"**\">foo</strong></p>", "**foo**\n"},
	{"13", "<h2>foo</h2><p>para<em>em</em></p>", "## foo\n\npara*em*\n"},
	{"12", "<a href=\"/bar\" title=\"baz\">foo</a>", "[foo](/bar \"baz\")\n"},
	{"11", "<img src=\"/bar\" alt=\"foo\" />", "![foo](/bar)\n"},
	{"10", "<img src=\"/bar\" />", "![](/bar)\n"},
	{"9", "<a href=\"/bar\">foo</a>", "[foo](/bar)\n"},
	{"8", "foo<br />bar", "foo\nbar\n"},
	{"7", "<code>foo</code>", "`foo`\n"},
	{"6", "<div class=\"vditor-wysiwyg__block\" data-type=\"code-block\"><pre><code data-code=\"foo\">foo<br></code></pre></div>", "```\nfoo\n```\n"},
	{"5", "<ul><li data-marker=\"*\">foo</li></ul>", "* foo\n"},
	{"4", "<blockquote>foo</blockquote>", "> foo\n"},
	{"3", "<h2>foo</h2>", "## foo\n"},
	{"2", "<p><strong><em>foo</em></strong></p>", "***foo***\n"},
	{"1", "<p><strong>foo</strong></p>", "**foo**\n"},
	{"0", "<p>foo</p>", "foo\n"},
}

func TestVditorDOM2Md(t *testing.T) {
	luteEngine := lute.New()

	for _, test := range vditorDOM2MdTests {
		md := luteEngine.VditorDOM2Md(test.from)
		if test.to != md {
			t.Fatalf("test case [%s] failed\nexpected\n\t%q\ngot\n\t%q\noriginal html\n\t%q", test.name, test.to, md, test.from)
		}
	}
}

var spinVditorDOMTests = []*parseTest{

	{"37", "<blockquote><p><wbr>\n</p></blockquote>", ""},
	{"36", "<div class=\"vditor-wysiwyg__block\" data-type=\"code-block\"><pre><code data-code=\"foo\" class=\"language-go\">foo</code></pre></div>", "<div class=\"vditor-wysiwyg__block\" data-type=\"code-block\"><pre><code data-code=\"foo\" class=\"language-go\"></code></pre></div>"},
	{"35", "<p><em data-marker=\"*\">foo</em></p><p><em data-marker=\"*\"><wbr><br></em></p>", "<p><em data-marker=\"*\">foo</em>\n</p><p><em data-marker=\"*\">\n<wbr></em>\n</p>"},
	{"34", "<p><span class=\"vditor-wysiwyg__block\" data-type=\"math-inline\"><code data-type=\"math-inline\">a1a</code></span></p>", "<p> <span class=\"vditor-wysiwyg__block\" data-type=\"math-inline\"><code data-type=\"math-inline\">a1a</code></span> \n</p>"},
	{"33", "<p><code>foo</code><wbr></p>", "<p> <code data-code=\"foo\"></code> <wbr>\n</p>"},
	{"32", "<p>```<wbr></p>", "<div class=\"vditor-wysiwyg__block\" data-type=\"code-block\"><pre><code data-code=\"\"><wbr></code></pre></div>"},
	{"31", "<div class=\"vditor-wysiwyg__block\" data-type=\"pre\"><pre><code><span style=\"color:#000080;font-weight:bold;\">package1<wbr></span>\n</code></pre><div class=\"vditor-wysiwyg__preview\" contenteditable=\"false\" data-render=\"false\"></div></div>", "<div class=\"vditor-wysiwyg__block\" data-type=\"html-block\"><pre><code data-type=\"html-block\"><div class=\"vditor-wysiwyg__block\" data-type=\"pre\"><pre><code><span style=\"color:#000080;font-weight:bold;\">package1<wbr></span>\n</code></pre><div class=\"vditor-wysiwyg__preview\" contenteditable=\"false\" data-render=\"false\"></div></div></code></pre></div>"},
	{"30", "<p>1. Node.js</p><p>2. Go<wbr></p>", "<ol><li data-marker=\"1.\"><p>Node.js\n</p></li><li data-marker=\"2.\"><p>Go<wbr>\n</p></li></ol>"},
	{"29", "<p><wbr><br></p>", "<p><wbr>\n</p>"},
	{"28", "<p>foo</p>\n<div class=\"vditor-wysiwyg__block\" data-type=\"html-block\"><textarea class=\"vditor-reset\" data-type=\"html-block\"><audio controls=\"controls\" src=\"http://localhost:8080/upload/file/2019/11/1440573175609-96444c00.mp3\"></audio></textarea></div>\n<p>bar</p>", "<p>foo\n</p><div class=\"vditor-wysiwyg__block\" data-type=\"html-block\"><pre><code data-type=\"html-block\"><textarea class=\"vditor-reset\" data-type=\"html-block\">&lt;audio controls=&#34;controls&#34; src=&#34;http://localhost:8080/upload/file/2019/11/1440573175609-96444c00.mp3&#34;&gt;&lt;/audio&gt;</textarea></code></pre></div><p>bar\n</p>"},
	{"27", "<p><wbr></p>", "<p><wbr>\n</p>"},
	{"26", "<p>![alt](src \"title\")</p>", "<p><img src=\"src\" alt=\"alt\" title=\"title\" />\n</p>"},
	{"25", "<pre><code class=\"language-java\"><wbr>\n</code></pre>", "<div class=\"vditor-wysiwyg__block\" data-type=\"code-block\"><pre><code data-code=\"\" class=\"language-java\"></code></pre></div>"},
	{"24", "<ul data-tight=\"true\"><li data-marker=\"*\"><wbr><br></li></ul>", "<ul data-tight=\"true\"><li data-marker=\"*\"><wbr></li></ul>"},
	{"23", "<ol><li data-marker=\"1.\">foo</li></ol>", "<ol data-tight=\"true\"><li data-marker=\"1.\">foo</li></ol>"},
	{"22", "<ul><li data-marker=\"*\">foo</li><li data-marker=\"*\"><ul><li data-marker=\"*\">bar</li></ul></li></ul>", "<ul data-tight=\"true\"><li data-marker=\"*\">foo</li><li data-marker=\"*\"><ul data-tight=\"true\"><li data-marker=\"*\">bar</li></ul></li></ul>"},
	{"21", "<p>[foo](/bar \"baz\")</p>", "<p><a href=\"/bar\" title=\"baz\">foo</a>\n</p>"},
	{"20", "<p>[foo](/bar)</p>", "<p><a href=\"/bar\">foo</a>\n</p>"},
	{"19", "<p>[foo]()</p>", "<p>[foo]()\n</p>"},
	{"18", "<p>[](/bar)</p>", "<p>[](/bar)\n</p>"},
	{"17", "<p>[]()</p>", "<p>[]()\n</p>"},
	{"16", "<p>[](</p>", "<p>[](\n</p>"},
	{"15", "<p><img alt=\"octocat\" class=\"emoji\" src=\"https://cdn.jsdelivr.net/npm/vditor/dist/images/emoji/octocat.png\" title=\"octocat\" /></p>", "<p><img alt=\"octocat\" class=\"emoji\" src=\"https://cdn.jsdelivr.net/npm/vditor/dist/images/emoji/octocat.png\" title=\"octocat\" />\n</p>"},
	{"14", ":octocat:", "<p><img alt=\"octocat\" class=\"emoji\" src=\"https://cdn.jsdelivr.net/npm/vditor/dist/images/emoji/octocat.png\" title=\"octocat\" />\n</p>"},
	{"13", "<div class=\"vditor-block\"><table><thead><tr><th>abc</th><th>def</th></tr></thead></table></div>", "<div class=\"vditor-wysiwyg__block\" data-type=\"html-block\"><pre><code data-type=\"html-block\"><div class=\"vditor-block\"><table><thead><tr><th>abc</th><th>def</th></tr></thead></table></div></code></pre></div>"},
	{"12", "<p><s data-marker=\"~~\">Hi</s> Hello, world!</p>", "<p><s data-marker=\"~~\">Hi</s> Hello, world!\n</p>"},
	{"11", "<p><del data-marker=\"~\">Hi</del> Hello, world!</p>", "<p><s data-marker=\"~\">Hi</s> Hello, world!\n</p>"},
	{"10", "<ul><li data-marker=\"*\" class=\"vditor-task\"><input checked=\"\" type=\"checkbox\" /> foo<wbr></li></ul>", "<ul data-tight=\"true\"><li data-marker=\"*\" class=\"vditor-task\"><input checked=\"\" type=\"checkbox\" /> foo<wbr></li></ul>"},
	{"9", "<ul><li data-marker=\"*\" class=\"vditor-task\"><input type=\"checkbox\" /> foo<wbr></li></ul>", "<ul data-tight=\"true\"><li data-marker=\"*\" class=\"vditor-task\"><input type=\"checkbox\" /> foo<wbr></li></ul>"},
	{"8", "> <wbr>", "<p>> <wbr>\n</p>"},
	{"7", "><wbr>", "<p>><wbr>\n</p>"},
	{"6", "<p>> foo<wbr></p>", "<blockquote><p>foo<wbr>\n</p></blockquote>"},
	{"5", "<p>foo</p><p><wbr><br></p>", "<p>foo\n</p><p><wbr>\n</p>"},
	{"4", "<ul data-tight=\"true\"><li data-marker=\"*\">foo<wbr></li></ul>", "<ul data-tight=\"true\"><li data-marker=\"*\">foo<wbr></li></ul>"},
	{"3", "<p><em data-marker=\"*\">foo<wbr></em></p>", "<p><em data-marker=\"*\">foo<wbr></em>\n</p>"},
	{"2", "<p>foo<wbr></p>", "<p>foo<wbr>\n</p>"},
	{"1", "<p><strong data-marker=\"**\">foo</strong></p>", "<p><strong data-marker=\"**\">foo</strong>\n</p>"},
	{"0", "<p>foo</p>", "<p>foo\n</p>"},
}

func TestSpinVditorDOM(t *testing.T) {
	luteEngine := lute.New()

	for _, test := range spinVditorDOMTests {
		html := luteEngine.SpinVditorDOM(test.from)
		if test.to != html {
			t.Fatalf("test case [%s] failed\nexpected\n\t%q\ngot\n\t%q\noriginal html\n\t%q", test.name, test.to, html, test.from)
		}
	}
}
