package crawl

import "code.google.com/p/go.net/html"

const (
	FAIL = -1
	PAGE = 1
	ASSET = 2
)

type rule func(*html.Node) (rule_code int, value string)

func PageRule(node *html.Node) (int, string) {
	return generalized_rule(node, PAGE, "a", "href");
}

func LinkAssetRule(node *html.Node) (int, string) {
	return generalized_rule(node, ASSET, "link", "href")
}

func ImageAssetRule(node *html.Node) (int, string) {
	return generalized_rule(node, ASSET, "img", "src")
}

func ScriptAssetRule(node *html.Node) (int, string) {
	return generalized_rule(node, ASSET, "script", "src")
}

func generalized_rule(node *html.Node, code int, data, key string) (int, string) {
	if node.Data == data {
		for _, attr := range(node.Attr) {
			if attr.Key == key {
				return code, attr.Val
			}
		}
	}
	return FAIL, ""
}
