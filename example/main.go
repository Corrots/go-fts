package main

import (
	"fmt"
	"github.com/corrots/gofts"
	"log"
)

const (
	DocFile = "/Users/dwcnmac39/Downloads/201904/enwiki-latest-abstract1.xml"
)

func main() {
	documents, err := gofts.LoadDocuments(DocFile)
	if err != nil {
		log.Fatal(err)
	}
	search := gofts.Search(documents, "cat")
	fmt.Println(search)
}
