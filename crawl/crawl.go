package crawl

import (
	"net/http"
	"code.google.com/p/go.net/html"
	"fmt"
)

func OutputURL(url string) {
	request, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to open stream.")
		return
	}
	defer request.Body.Close()

	document, err := html.Parse(request.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var descent func(*html.Node)
	descent = func(node *html.Node) {
		if node.Type == html.ElementNode && (node.Data == "a" || node.Data == "link" || node.Data == "img") {
			for _, attr := range node.Attr {
				if (attr.Key == "href" || attr.Key == "src") {
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
