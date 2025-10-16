package main

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "net/url"
    "testing"
)

func TestHandleHelloWorld(t *testing.T) {
    tests := []struct {
        name       string
        query      string
        wantStatus int
        wantMsg    string
        wantErr    string
    }{
        {"Valid name starting with A", "Alice", http.StatusOK, "Hello Alice", ""},
        {"Valid name starting with M", "Mary", http.StatusOK, "Hello Mary", ""},
        {"Invalid name starting with N", "Nancy", http.StatusBadRequest, "", "Invalid Input"},
        {"Invalid name starting with Z", "Zane", http.StatusBadRequest, "", "Invalid Input"},
        {"Empty name", "", http.StatusBadRequest, "", "Invalid Input"},
        {"Missing name param", " ", http.StatusBadRequest, "", "Invalid Input"},
    }

    for _, tt := range tests {
        req := httptest.NewRequest(http.MethodGet, "/hello-world?name="+url.QueryEscape(tt.query), nil)
        w := httptest.NewRecorder()

        handleHelloWorld(w, req)

        res := w.Result()
        defer res.Body.Close()

        var resp Response
        json.NewDecoder(res.Body).Decode(&resp)

        if res.StatusCode != tt.wantStatus {
            t.Errorf("[%s] expected status %d, got %d", tt.name, tt.wantStatus, res.StatusCode)
        }
        if resp.Message != tt.wantMsg {
            t.Errorf("[%s] expected message %q, got %q", tt.name, tt.wantMsg, resp.Message)
        }
        if resp.Error != tt.wantErr {
            t.Errorf("[%s] expected error %q, got %q", tt.name, tt.wantErr, resp.Error)
        }
    }
}
