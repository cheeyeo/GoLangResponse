package main

type Response1 struct {
  Page   int
  Fruits []string
}

type Profile struct {
  Name    string
  Hobbies []string
}

type XmlProfile struct {
  Name    string
  Hobbies []string `xml:"Hobbies>Hobby"`
}
