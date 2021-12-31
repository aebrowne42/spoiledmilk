package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/aebrowne42/spoiledmilk/pkg/qa"
)

func questionEndpoint(w http.ResponseWriter, req *http.Request) {
	body, _ := io.ReadAll(req.Body)
	fmt.Println(string(body))

	if string(body) == "" {
		fmt.Fprintf(w, "You need to ask me a question first, bitch")
	} else {
		fmt.Fprintf(w, qa.AnswerQuestion(string(body)))
	}

}

func main() {

	http.HandleFunc("/home", questionEndpoint)

	http.ListenAndServe(":8090", nil)
}
