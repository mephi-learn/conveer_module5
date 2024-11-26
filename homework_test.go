package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_homeWorkHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	homeWorkHandler(w, req)
	res := w.Result()
	defer func() {
		_ = res.Body.Close()
	}()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(bytes.TrimSpace(data)) != "Привет, я домашнее задание" {
		t.Errorf("unexpected message %v", string(data))
	}
}
