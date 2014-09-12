package crawl

import (
	"code.google.com/p/go.net/html"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Outputs results of Tokenize() from the given url string
// To be used for debugging purpose
func OutputURL(url string) {
	request, err := http.Get(url)
	if err != nil {
		log.Println("Failed to open stream: ", err)
		return
	}
	defer request.Body.Close()

	tokens, err := Tokenize(request.Body, PageRule, LinkAssetRule, ImageAssetRule, ScriptAssetRule)

	for path, code := range tokens {
		fmt.Println(code, path)
	}
}

// Tokenize the HTML structure from the Reader r and generates a set of token based on the rules provided
// Each token(key) will denote an integer identifying the token type.
// If an error occurs, it will mostly happen from the reader stream
func Tokenize(r io.Reader, rules ...Rule) (token_set map[string]int, err error) {
	document, err := html.Parse(r)
	if err != nil {
		log.Fatalln(err)
		return
	}

	token_set = make(map[string]int) // Avoid duplicate tokens
	var descent func(*html.Node)
	descent = func(node *html.Node) {
		if node.Type == html.ElementNode {
			for _, task := range rules { // Process the node based on each rule given
				if code, value := task(node); code != FAIL {
					token_set[value] = code // If the rule given is valid, apply it to the map
					break
				}
			}
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			descent(child)
		}
	}

	descent(document)
	return
}
