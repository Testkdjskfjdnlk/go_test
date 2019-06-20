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
  t.Execute(w, p)
}

func handler(w http.ResponseWriter, r *http.Request) {
  log.Printf("Loading index.html...\n")
  title := "Hello from the web!"

  // Use os/user package to get username from os
  user, err := user.Current()
  if err != nil {
    log.Fatal("Unable to get current user from os")
    panic(err)
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

  // Handle root uri
  http.HandleFunc("/", handler)

  serverPort := 8080

  log.Printf("Started server on port %d\n", serverPort)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
