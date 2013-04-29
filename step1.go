package main

import (
	"fmt"
	"net/http"
)

var text = `
	<html>
		<head>
			<title>gofuncyourself</title>
		</head>
		<body>
			You have seen this <a href="https://code.google.com/p/nogotovogo/source/browse/step2.go">before</a>
		</body>
	</html>`

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, text)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
