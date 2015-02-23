package main

import (
  "fmt"
  "log"
  "net/http"
  "os"
  "github.com/alext/tablecloth"
)

var (
  eventPort  = getenvDefault("PORT", "3097")
)

func getenvDefault(key string, defaultVal string) string {
  val := os.Getenv(key)
  if val == "" {
    val = defaultVal
  }

  return val
}

func main() {
  publicMux := http.NewServeMux()

  publicMux.HandleFunc("/", RootHandler())
  publicMux.HandleFunc("/healthcheck", HealthcheckHandler())
  publicMux.HandleFunc("/profile", ProfileHandler())
  publicMux.HandleFunc("/profile/xml", XmlProfileHandler())
  publicMux.HandleFunc("/image", FileHandler())

  log.Println("httptestapp: listening for events on " + eventPort)

  err := tablecloth.ListenAndServe(fmt.Sprintf(":%v", eventPort), publicMux, "reports")

  if err != nil {
    log.Fatal(err)
  }
}
