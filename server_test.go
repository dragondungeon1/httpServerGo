package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func ListenAndServe(addr string, handler http.Handler)error

type Handler interface {
	ServeHTTP(http.ResponseWriter, http.Request)
}

func PlayerServer(resWriter http.ResponseWriter, request *http.Request) {

}
func TestGETPlayers(t *testing.T) {
	t.Run("returns pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		expected := response.Body.String()
		got := "20"

		if expected != got {
			t.Errorf("expected %q but got %q", expected, got)
		}
		})
}