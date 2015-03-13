package main

import (
  "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
  "net/http/httptest"
  "net/http"
  "testing"
  "strings"
  "log"
)

type htmlTest struct {
  name     string
  shouldBe string
}

func getConnection() *mgo.Session {
  conn, err := mgo.Dial("localhost")
  if err != nil {
    log.Fatal(err)
  }
  return conn
}

func TestSaveNewDoc(t *testing.T) {
  conn := getConnection()
  defer conn.Close()
  c := conn.DB("sources").C("sources")

  source := Source{"nyc", "http://nyc.org/data", "NYC's open data"}
  err := save(&source, nil)
  if err != nil {
    t.Errorf("Save failed, got %v", err)
  }
  result := Source{}
  err = c.Find(bson.M{"Name" : "nyc"}).One(&result)
  if err != nil {
    t.Errorf("Find failed, got %v", err)
  }
  if source != result {
    t.Errorf("Inserted and found not the same")
  }

}

func TestSaveUpdateDoc(t *testing.T) {
  t.Skip("Save update not yet implemented")
}

func TestIndexHandler(t *testing.T) {
  tests := []htmlTest{
    {"Name", "Name:"},
    {"Link", "Link:"},
    {"Desc", "Description:"},
    {"Title", "<title>DataLink</title>"},
  }

  req, err := http.NewRequest("GET", "http://localhost/", nil)
  if err != nil {
    log.Fatal(err)
  }

  w := httptest.NewRecorder()
  index(w, req)
  body := w.Body.String()
  for _, test := range tests {
    if !strings.Contains(body, test.shouldBe) {
      t.Errorf("%v failed, got %v", test.name, body)
    }
  }
}

func TestCreateHandler(t *testing.T) {
  tests := []htmlTest{
    {"Title Input", `<input type="text" name="name"`},
    {"Link Input", `<input type="text" name="link"`},
    {"Description Input", `<textarea name="description"`},
  }

  req, err := http.NewRequest("GET", "http://localhost/new", nil)
  if err != nil {
    log.Fatal(err)
  }

  w := httptest.NewRecorder()
  createHandler(w, req)
  body := w.Body.String()
  for _, test := range tests {
    if !strings.Contains(body, test.shouldBe) {
      t.Errorf("%v failed, got %v", test.name, body)
    }
  }
}

func TestViewHandler(t *testing.T) {
  t.Skip("View handler not yet implemented")
}

func TestEditHandler(t *testing.T)  {
  t.Skip("Edit handler not yet implemented")
}

func TestSaveHandler(t *testing.T)  {
  t.Skip("Save handler not yet implemented")
}

func TestDeleteHandler(t *testing.T)  {
  t.Skip("Delete handler not yet implemented")
}
