package main

import (
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
  t.Skip("Create handler not yet implemented")
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
