package main

import (
  "html/template"
  "log"
  "net/http"

  "github.com/joaonsantos/hello-go-web/utils"
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


func rootHandler(w http.ResponseWriter, r *http.Request) {
  valid := validatePath(w, r); if !valid {
    return
  }

  userName := utils.GetUsernameFromOS()
  title := "Hello from the web!"
  pageStruct, _ := model.LoadPage(title, userName)

  renderTemplate(w, "index.html", pageStruct)
}

func logRequest(handler http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
    handler.ServeHTTP(w, r)
  })
}

func main() {
  // Register static directory
  http.Handle("/static/",
    http.StripPrefix("/static/",
      http.FileServer(http.Dir("static"))))

  // Register handlers
  http.HandleFunc("/", rootHandler)

  serverPort := 8080

  log.Printf("Started server on port %d\n", serverPort)
  log.Fatal(http.ListenAndServe("0.0.0.0:8080", logRequest(http.DefaultServeMux)))
}
