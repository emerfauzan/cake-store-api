package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func (test *Test) GetCakes(t *testing.T) {
	req, err := http.NewRequest("GET", "/cakes", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", test.token))
	r := httptest.NewRecorder()
	handler := http.HandlerFunc(test.handler.GetCakesHandler)
	handler.ServeHTTP(r, req)
	if status := r.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
