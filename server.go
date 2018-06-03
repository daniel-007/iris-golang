
package main

import (
    "net/http"
    "fmt"
    "path/filepath"
    "os"
    str "strings"
    "github.com/labstack/echo"
 //   "github.com/labstack/gommon/color"
)

type Document struct {
    Name     string `json:"name"`
    Location string `json:"location"`
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
                Location: filepath.ToSlash(str.TrimPrefix(path, root)),
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
            func(f string) bool { return str.HasSuffix(str.ToLower(f), ".md") })

    if err != nil {
        
        fmt.Println("Error occurred")
        fmt.Println(err)

        return
    }

    for _, doc := range *docs {
        fmt.Printf("%v\n", doc)
    }

    e := echo.New()

    e.GET("/", func(c echo.Context) error {
        return c.JSON(http.StatusOK, docs)
    })

    e.Logger.Fatal(e.Start(":1323"))
}
