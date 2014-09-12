package crawl

import (
	"code.google.com/p/go.net/html"
	"strings"
	"testing"
)

type simple_rule_test struct {
	html           string
	expected_code  int
	expected_value string
	law            rule
}

func (s simple_rule_test) assert(t *testing.T, actual_code int, actual_value string) {
	if s.expected_code != actual_code || s.expected_value != actual_value {
		t.Logf("Expected {%d \u2014 %s} but found {%d \u2014 %s}", s.expected_code, s.expected_value,
			actual_code, actual_value)
		t.FailNow()
	}
}

func TestRules(t *testing.T) {
	rules_test := []simple_rule_test{
		simple_rule_test{"<a href=\"www.example.com\" />", PAGE, "www.example.com", PageRule},
		simple_rule_test{"<img src=\"/example.png\" />", ASSET, "/example.png", ImageAssetRule},
		simple_rule_test{"<link href=\"/example.css\" />", ASSET, "/example.css", LinkAssetRule},
		simple_rule_test{"<script src=\"/example.js\" />", ASSET, "/example.js", ScriptAssetRule}}

	for _, test := range rules_test {
		r := strings.NewReader(test.html)
		document, _ := html.Parse(r)
		code, val := crawl(document, test.law)
		test.assert(t, code, val)
	}

	fail := simple_rule_test{"<p href=\"lol.png\" src=\"nope.gif\">Fail!</p>", FAIL, "", nil}
	r := strings.NewReader(fail.html)
	document, err := html.Parse(r)
	if err != nil {
	}

	for _, law := range []rule{PageRule, ImageAssetRule, LinkAssetRule, ScriptAssetRule} {
		code, val := crawl(document, law)
		fail.assert(t, code, val)
	}
}

func crawl(node *html.Node, r rule) (code int, val string) {
	if node.Type == html.ElementNode {
		if code, val = r(node); code != FAIL {
			return code, val
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if code, val := crawl(child, r); code != FAIL {
			return code, val
		}
	}

	return -1, ""
}
