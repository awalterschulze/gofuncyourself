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
			This is a little new
		</body>
	</html>`

type state struct {
}

func (this *state) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, text)
}

func main() {
	s := &state{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s.handler(w, r)
	})
	http.ListenAndServe(":3000", nil)
}
