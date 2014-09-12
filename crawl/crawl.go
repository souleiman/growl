package crawl

import (
	"code.google.com/p/go.net/html"
	"fmt"
	"io"
	"log"
	"net/http"
)

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

func Tokenize(r io.Reader, rules ...rule) (token_set map[string]int, err error) {
	document, err := html.Parse(r)
	if err != nil {
		log.Fatalln(err)
		return
	}

	token_set = make(map[string]int) // Avoid duplicate tokens
	var descent func(*html.Node)
	descent = func(node *html.Node) {
		if node.Type == html.ElementNode {
			for _, task := range rules {
				if code, value := task(node); code != FAIL {
					token_set[value] = code
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
