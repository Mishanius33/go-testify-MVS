package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreCorrectRequest(t *testing.T) {
	req := httptest.NewRequest("GET", "/?count=4&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusOK)
	assert.NotEmpty(t, responseRecorder.Body.String())

}

func TestMainHandlerWhenCountMoreWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=UnExistsCity", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	city := responseRecorder.Body.String()

	require.Equal(t, responseRecorder.Code, http.StatusBadRequest)
	assert.Equal(t, "wrong city value", city)

}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/?count=5&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusOK)
	assert.Len(t, strings.Split(responseRecorder.Body.String(), ","), totalCount)

}
