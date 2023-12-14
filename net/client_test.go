package net

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet(t *testing.T) {
	t.Run("Get() : Test expected output", func(t *testing.T) {
		expectedOutput := []uint8("Hello, World!")
		handler := func(w http.ResponseWriter, r *http.Request) {
			_, er := w.Write(expectedOutput)
			if er != nil {
				t.Errorf("Get() error: %v", er)
			}
		}

		httpTestServer := httptest.NewServer(http.HandlerFunc(handler))
		defer httpTestServer.Close()

		actualOutput, err := Get(httpTestServer.URL)
		if err != nil {
			t.Errorf("Get() returned an error: %v", err)
		}
		if string(actualOutput) != string(expectedOutput) {
			t.Errorf("Get() returned unexpected output: got %v want %v", actualOutput, expectedOutput)
		}
	})
}
