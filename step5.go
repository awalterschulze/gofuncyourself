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
			Return a function from a method
		</body>
	</html>`

type state struct {
}

func (this *state) newHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, text)
	}
}

func main() {
	s := &state{}
	http.HandleFunc("/", s.newHandler())
	http.ListenAndServe(":3000", nil)
}
