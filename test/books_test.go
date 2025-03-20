package tests

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "Lib-CURD/handlers"
)

func TestGetBooks(t *testing.T) {
    req, err := http.NewRequest("GET", "/books", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handlers.GetBooks(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Expected status 200, got %v", status)
    }
}

