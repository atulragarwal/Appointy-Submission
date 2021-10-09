package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserCheck(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/?id=10", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CheckUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Error-Wrong Code: %v", status)
	}
}

func TestPostCheck(t *testing.T) {
	req, err := http.NewRequest("GET", "/posts/?id=1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CheckPost)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Error-Wrong Code: %v", status)
	}
}

func TestUserPostCheck(t *testing.T) {
	req, err := http.NewRequest("GET", "/posts/users/?id=1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUserPosts)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Error-Wrong Code: %v", status)
	}
}

func TestCreatePostCheck(t *testing.T) {
	var jsonStr = []byte(`{"ID":4,"PostId":15,"Caption":"Test Image","ImageUrl":"/test/test"}`)
	req, err := http.NewRequest("POST", "/posts/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(MakePost)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Error-Wrong Code: %v", status)
	}
}

func TestCreateUserCheck(t *testing.T) {
	var jsonStr = []byte(`{"ID":4,"Name":"test","Email":"test@gmail.com","Password":"tester"}`)
	req, err := http.NewRequest("POST", "/users/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(MakeUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Error-Wrong Code: %v", status)
	}
}
