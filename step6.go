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
			Functions can be types
		</body>
	</html>`

type handlerFunc func(w http.ResponseWriter, r *http.Request)

type state struct {
}

func (this *state) newHandler() handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, text)
	}
}

func main() {
	s := &state{}
	http.HandleFunc("/", s.newHandler())
	http.ListenAndServe(":3000", nil)
}
