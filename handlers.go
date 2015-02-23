package main

import (
  "encoding/json"
  "encoding/xml"
  "fmt"
  "net/http"
  "log"
  "html/template"
  "path"
)

func RootHandler() http.HandlerFunc{
  return func(w http.ResponseWriter, req *http.Request) {

    fp := path.Join("templates", "index.html")
    tmpl, err := template.ParseFiles(fp)

    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    if err := tmpl.Execute(w, ""); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  }
}

func HealthcheckHandler() http.HandlerFunc{
  return func(w http.ResponseWriter, req *http.Request) {
    if req.Method != "GET" {
      http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
      w.Header().Set("Allow", "GET")
      return
    }

    w.Header().Set("Content-Type", "application/json")

    var responseCode int

    responseCode = http.StatusOK

    w.WriteHeader(responseCode)

    encoder := json.NewEncoder(w)

    res2D := &Response1{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}

    log.Println(res2D)
    err := encoder.Encode(res2D)

    if err != nil {
      http.Error(w, fmt.Sprintf("Cannot encode healthcheck response: %v", err), http.StatusInternalServerError)
    }
  }
}

func ProfileHandler() http.HandlerFunc{
  return func(w http.ResponseWriter, req *http.Request) {
    if req.Method != "GET" {
      http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
      w.Header().Set("Allow", "GET")
      return
    }

    w.Header().Set("Content-Type", "application/json")

    var responseCode int

    responseCode = http.StatusOK

    w.WriteHeader(responseCode)

    profile := Profile{"Alex", []string{"snowboarding", "programming"}}
    js, err := json.Marshal(profile)

    if err != nil {
      http.Error(w, fmt.Sprintf("Cannot encode healthcheck response: %v", err), http.StatusInternalServerError)
    }

    w.Write(js)
  }
}

func XmlProfileHandler() http.HandlerFunc{
  return func(w http.ResponseWriter, req *http.Request) {
    profile := XmlProfile{"Alex", []string{"snowboarding", "programming"}}

    x, err := xml.MarshalIndent(profile, "", "  ")
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    w.Header().Set("Content-Type", "application/xml")
    w.Write(x)
  }
}

func FileHandler() http.HandlerFunc{
  return func(w http.ResponseWriter, req *http.Request){
    fp := path.Join("templates/images", "tile.png")
    http.ServeFile(w, req, fp)
  }
}


