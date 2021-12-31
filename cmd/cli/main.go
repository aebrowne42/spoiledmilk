package main

import (
	"flag"
	"fmt"

	"github.com/aebrowne42/spoiledmilk/pkg/qa"
)

func main() {
	questionFlag := flag.String("question",
		"", "Use this if you would like to ask the 8Ball a question")
	flag.Parse()

	if *questionFlag == "" {
		fmt.Println("Please ask me a question")
		return
	}

	fmt.Println(qa.AnswerQuestion(*questionFlag))
}
