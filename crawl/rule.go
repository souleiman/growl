package crawl

import "code.google.com/p/go.net/html"

const (
	FAIL  = -1 // Failure code
	IGN   = 0  // Reserved
	PAGE  = 1  // Denotes that the code of the item is a page (usually used with a tag)
	ASSET = 2  // Denotes that the code of the item is an asset (used with img, link, script)
)

// Template of the rule function which will return the code constant and string of the extracted value
type Rule func(*html.Node) (int, string)

// Checks if the given node is a tag or hyperlink, if so return the result from 'href' and PAGE constant.
// Otherwise, return a FAIL code
func PageRule(node *html.Node) (int, string) {
	return generalized_rule(node, PAGE, "a", "href")
}

// Checks if the given node is a link tag, if so return the result from 'href' and ASSET constant.
// Otherwise, return a FAIL code
func LinkAssetRule(node *html.Node) (int, string) {
	return generalized_rule(node, ASSET, "link", "href")
}

// Checks if the given node is a img, if so return the result from 'src' and ASSET constant.
// Otherwise, return a FAIL code
func ImageAssetRule(node *html.Node) (int, string) {
	return generalized_rule(node, ASSET, "img", "src")
}

// Checks if the given node is a script, if so return the result from 'src' and ASSET constant.
// Otherwise, return a FAIL code
func ScriptAssetRule(node *html.Node) (int, string) {
	return generalized_rule(node, ASSET, "script", "src")
}

// Checks if the given node is a source tag which is used in multimedia format. Return any result from src attribute and ASSET constant.
// Otherwise, return a FAIL code.
func SourceAssetRule(node *html.Node) (int, string) {
	return generalized_rule(node, ASSET, "source", "src")
}

// Helper function that allows the functions to provide the code value and expected test data to check for
// Returns the appropriate result of the data found from the attribute of the given key
// Otherwise, returns FAIL code.
func generalized_rule(node *html.Node, code int, data, key string) (int, string) {
	if node.Data == data {
		for _, attr := range node.Attr {
			if attr.Key == key {
				return code, attr.Val
			}
		}
	}
	return FAIL, ""
}
