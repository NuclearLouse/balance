package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer_HandleRegisterJSON(t *testing.T) {

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":           "user@example.org",
				"password":        "password",
				"repeat_password": "password",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid params",
			payload: map[string]string{
				"email": "invalid",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "not equal passwords",
			payload: map[string]string{
				"email":           "user@example.org",
				"password":        "password",
				"repeat_password": "passworf",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	s := testServer(t)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/register", b)
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_HandleRegisterHTTP(t *testing.T) {

	testCases := []struct {
		name         string
		urlformvalue string
		expectedCode int
	}{
		{
			name:         "valid",
			urlformvalue: "email=user@example.org&password=password&passRepeat=password",
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid form value",
			urlformvalue: "email=user@example.org&password=password&passRepeat=passworf",
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	s := testServer(t)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b := strings.NewReader(tc.urlformvalue)
			req, _ := http.NewRequest(http.MethodPost, "/register", b)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			rec := httptest.NewRecorder()
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_HandleLoginHTTP(t *testing.T) {
	
}
