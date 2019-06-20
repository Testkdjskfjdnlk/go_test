package model

import "time"

type Page struct {
  Title string
  Name  string
  Time  string
}

func LoadPage(title string, name string) (*Page, error) {
  time := time.Now().Format(time.Stamp)
  return &Page{Title: title, Name: name, Time: time}, nil
}

