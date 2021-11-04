package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var exmplHtml = `<html>
<head>
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
</head>
<body>
  <h1>Social stuffs</h1>
  <div>
    <a href="https://www.twitter.com/joncalhoun">
      Check me out on twitter
      <i class="fa fa-twitter" aria-hidden="true"></i>
    </a>
    <a href="https://github.com/gophercises">
      Gophercises is on <strong>Github</strong>!
    </a>
  </div>
</body>
</html>`

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fileFlag := flag.String("f", "ex2.html", ".html file to parse")
	flag.Parse()

	r, err := os.Open(*fileFlag)
	checkErr(err)

	doc, err := html.Parse(r)
	checkErr(err)

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				fmt.Println(a.Val)
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)
}
