package lotteline

import (
  "html/template"
  "log"
  "net/http"
  "path"
)

func handler(w http.ResponseWriter, r *http.Request) {
  lp := path.Join("templates", "layout.tmpl")
  var baseTemplate = r.URL.Path + ".tmpl"
  if r.URL.Path == "/" {
    baseTemplate = "redirect.tmpl"
  }
  fp := path.Join("templates", baseTemplate)

  tmpl, _ := template.ParseFiles(lp, fp)
  tmpl.ExecuteTemplate(w, "layout", nil)
}

func init() {
  http.HandleFunc("/", handler)
  log.Println("Listening...")
  http.ListenAndServe(":8080", nil)
}