package crawl

import (
	"code.google.com/p/go.net/html"
	"fmt"
	"net/http"
	"io"
	"log"
)

func OutputURL(url string) {
	request, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to open stream.")
		return
	}
	defer request.Body.Close()

	tokens, _ := Tokenize(request.Body, PageRule, LinkAssetRule, ImageAssetRule, ScriptAssetRule)

	for _, token := range(tokens) {
		fmt.Println(token)
	}
}

func Tokenize(r io.Reader, rules ...rule) (tokens []Token, err error) {
	document, err := html.Parse(r)
	if err != nil {
		log.Fatalln(err)
		return
	}


	var descent func(*html.Node)
	descent = func(node *html.Node) {
		if node.Type == html.ElementNode {
			for _, task := range(rules) {
				if code, value := task(node); code != FAIL {
					tokens = append(tokens, Token{code: code, value: value})
				}
			}
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			descent(child)
		}
	}

	descent(document)
	return tokens, nil
}
