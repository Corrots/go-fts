package gofts

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Document struct {
	ID int
	Title string `xml:"title"`
	URL string `xml:"url"`
	Text string `xml:"abstract"`
}

func LoadDocuments(path string) ([]Document, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("file open err: %v\n", err)
	}
	defer f.Close()
	decoder := xml.NewDecoder(f)
	dump := struct {
		Documents []Document `xml:"doc"`
	}{}
	if err := decoder.Decode(&dump); err != nil {
		return nil, err
	}
	docs := dump.Documents
	for i := range docs {
		docs[i].ID = i
	}
	return docs, nil
}