package main

import (
	"code.google.com/p/go.net/html"
	"fmt"
	"net/http"
)

func main() {
	request, _ := http.Get("http://www.digitalocean.com/")
	defer request.Body.Close()

	document, err := html.Parse(request.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var descent func(*html.Node)
	descent = func(node *html.Node) {
		if node.Type == html.ElementNode && (node.Data == "a" || node.Data == "link") {
			for _, attr := range(node.Attr) {
				if attr.Key == "href" {
					fmt.Println(attr.Val)
				}
			}
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			descent(child)
		}
	}

	descent(document)
}
