package main

import (
	"fmt"
	"strings"

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
	//resp, err := http.Get("http://www.google.com")
	//checkErr(err)

	//doc, err := html.Parse(ioutil.ReadAll(resp.Body))
	r := strings.NewReader(exmplHtml)
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
