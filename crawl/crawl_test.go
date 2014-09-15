package crawl

import (
	"bufio"
	"os"
	"testing"
)

func TestCrawl(t *testing.T) {
	expected := map[string]int{
		"https://ajax.googleapis.com/ajax/libs/angularjs/1.3.0-rc.1/angular.min.js": ASSET,
		"/pricing/":          PAGE,
		"/features/":         PAGE,
		"/help/":             PAGE,
		"/stylesheet.css":    ASSET,
		"/static/footer.css": ASSET,
		"/static/header.css": ASSET,
		"/community/":        PAGE,
		"/image.png":         ASSET,
		"/image2.jpg":        ASSET,
		"horse.ogg":          ASSET,
		"horse.mp3":          ASSET,
		"movie.mp4":          ASSET,
		"movie.ogg":          ASSET,
	}

	file, err := os.Open("crawl_test.html")
	if err != nil {
		t.Fatal(err)
		t.FailNow()
	}
	defer file.Close()

	actual, err := Tokenize(bufio.NewReader(file), PageRule, ScriptAssetRule, LinkAssetRule, ImageAssetRule, SourceAssetRule)

	if err != nil {
		t.Fatalf("Should not have failed: ", err)
		t.FailNow()
	}

	if len(expected) != len(actual) {
		t.Fatalf("Expected same number of result")
		t.FailNow()
	}

	for path, code := range actual {
		if expected[path] == 0 {
			t.Fatalf("Could not find: %s", path)
			t.FailNow()
		}

		if expected[path] != code {
			t.Fatalf("Expected code for %s: %d but found %d", path, expected[path], code)
		}
	}
}
