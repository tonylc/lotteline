package lotteline

import (
  "html/template"
  "log"
  "net/http"
  "path"
  "strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
  if (strings.HasPrefix(r.URL.Path, "/templates")) {
    staticHtmlHandler(w, r.URL.Path)
  } else {
    if r.URL.Path == "/" {
      goTemplateHandler(w, "redirect.tmpl")
    } else {
      goTemplateHandler(w, r.URL.Path + ".tmpl")
    }
  }
}

// path is /templates/test.html
func staticHtmlHandler(w http.ResponseWriter, basePath string) {
  // strip leading '/'
  fp := basePath[1:]
  tmpl, _ := template.New("").Delims("<<", ">>").ParseFiles(fp)
  tmpl.ExecuteTemplate(w, "body", nil)
}

func goTemplateHandler(w http.ResponseWriter, basePath string) {
  lp := path.Join("templates", "layout.tmpl")
  fp := path.Join("templates", basePath)
  tmpl, _ := template.New("").Delims("<<", ">>").ParseFiles(lp, fp)
  tmpl.ExecuteTemplate(w, "layout", nil)
}

func init() {
  http.HandleFunc("/", handler)
  log.Println("Listening...")
  http.ListenAndServe(":8080", nil)
}