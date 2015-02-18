package lotteline

import (
  // "fmt"
  "html/template"
  "log"
  "net/http"
  "path"
)

func handler(w http.ResponseWriter, r *http.Request) {
  // fmt.Fprintf(w, "request path", r.URL, r.URL.Path)
  // t, _ := template.ParseFiles("templates/home.html")
  //t, _ := template.ParseFiles("templates/redirect.html")
  // t.Execute(w, nil)

  lp := path.Join("templates", "layout.tmpl")
  fp := path.Join("templates", r.URL.Path + ".tmpl")

  tmpl, _ := template.ParseFiles(lp, fp)
  tmpl.ExecuteTemplate(w, "layout", nil)
}

func init() {
  http.HandleFunc("/", handler)
  log.Println("Listening...")
  http.ListenAndServe(":8080", nil)
}