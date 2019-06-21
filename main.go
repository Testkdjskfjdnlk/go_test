package main

import (
    "html/template"
    "log"
    "net/http"
    "os/user"

    "github.com/joaonsantos/hello-go-web/model"
)

func renderTemplate(w http.ResponseWriter, tmpl string, p *model.Page) {

  tmpl = "templates/" + tmpl
  t := template.Must(template.ParseFiles(tmpl))

  err := t.Execute(w, p); if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func validatePath(w http.ResponseWriter, r *http.Request) (bool) {
  valid := true
  m := "/" != r.URL.Path

  if m {
    valid = false
    http.NotFound(w, r)
  }

  return valid
}

func handler(w http.ResponseWriter, r *http.Request) {
  log.Printf("Loading index.html...\n")
  title := "Hello from the web!"

  valid := validatePath(w, r); if !valid {
    return
  }

  // Use os/user package to get username from os
  user, err := user.Current(); if err != nil {
    log.Fatal("Unable to get current user from os")
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  // Load page structure
  name := user.Name
  pageStruct, _ := model.LoadPage(title, name)

  // Render the template
  renderTemplate(w, "index.html", pageStruct)
}

func main() {
  // Register static directory
  http.Handle("/static/",
    http.StripPrefix("/static/",
      http.FileServer(http.Dir("static"))))

  // Register handlers
  http.HandleFunc("/", handler)

  serverPort := 8080

  log.Printf("Started server on port %d\n", serverPort)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
