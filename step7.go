package main

import (
	"fmt"
	"net/http"
)

var top = `
	<html>
		<head>
			<title>gofuncyourself</title>
		</head>
		<body>
		The `
var bottom = `
		spider is getting tired	
		</body>
	</html>`

type handlerFunc func(w http.ResponseWriter, r *http.Request)

type state struct {
	i int
}

func (this *state) newHandler() handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s %d %s", top, this.i, bottom)
	}
}

func main() {
	s := &state{0}
	http.HandleFunc("/", s.newHandler())
	s.i = 1
	http.ListenAndServe(":3000", nil)
}
