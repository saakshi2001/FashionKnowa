package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func Test_ShowUser(t *testing.T) {
    req, err := http.NewRequest("POST", "/users", nil)
    if err != nil {
        t.Fatal(err)
    }

    res := httptest.NewRecorder()
    ShowUser(res, req)

    exp := "Hello World"
    act := res.Body.String()
    if exp != act {
        t.Fatalf("Expected %s got %s", exp, act)
    }
}
func Test_ShowPost(t *testing.T) {
    req, err := http.NewRequest("POST", "/posts", nil)
    if err != nil {
        t.Fatal(err)
    }

    res := httptest.NewRecorder()
    ShowPost(res, req)

    exp := "Hello World"
    act := res.Body.String()
    if exp != act {
        t.Fatalf("Expected %s got %s", exp, act)
    }
}