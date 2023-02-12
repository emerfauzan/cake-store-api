package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func (test *Test) AuthLogin(t *testing.T) {
	var jsonStr = []byte(`{"username": "cake_shop","password": "P@ssword123"}`)

	req, err := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	r := httptest.NewRecorder()
	handler := http.HandlerFunc(test.handler.AuthLoginHandler)
	handler.ServeHTTP(r, req)
	if status := r.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	type loginUser struct {
		Token string `json:"token"`
	}

	type loginData struct {
		Data loginUser `json:"data"`
	}

	var result loginData

	_ = json.NewDecoder(r.Body).Decode(&result)

	test.token = result.Data.Token
	fmt.Println("== AuthLogin PASS")

}
