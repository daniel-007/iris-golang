
package main

import (
    "fmt"
   // "strconv"
    "path/filepath"
    "os"
    "strings"
)

type Document struct {
    Name string
    Location string
}

type DocumentAccessor interface {
    GetDocuments() *[]Document
}

type LocalDocumentAccessor struct {
}

func (acc *LocalDocumentAccessor) GetDocuments(root string, filter func(string) bool) (*[]Document, error) {

    var docs []Document

    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

        if filter(path) {
            docs = append(docs, Document {
                Name: info.Name(),
                Location: strings.TrimPrefix(path, root),
            })
        }

        return err
    })

    if err != nil {
        return nil, err
    }

    return &docs, nil
}

func main() {

    fmt.Println("Program started")

    j := LocalDocumentAccessor {}

    docs, err :=
        j.GetDocuments(
            "C:\\Users\\Nitzanz\\Dropbox\\Projects",
            func(f string) bool { return strings.HasSuffix(f, ".md") })

    if err != nil {
        
        fmt.Println("Error occurred")
        fmt.Println(err)

        return
    }

    fmt.Println("hello, world")
    fmt.Println(docs)
}
