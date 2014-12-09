package main

import (
	"fmt"
    "net/http"
    "log"
    "html/template"
    "encoding/base64"

	"code.google.com/p/rsc/qr"
)

type Page struct {
    Img string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        p := &Page{Img: ""}
        txt := r.FormValue("text")
        if txt != "" {
            log.Printf("txt: %s \n", txt)
        }

        c, _ := qr.Encode(txt, qr.Q)
        png := c.PNG()
        p.Img = base64.StdEncoding.EncodeToString(png)

        t, _ := template.ParseFiles("templates/index.tpl")
        t.Execute(w, p)
    })

    fmt.Println("Running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
