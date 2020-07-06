package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer_HandleCreateClientHTTP(t *testing.T) {

	s := testServer(t)

	testCases := []struct {
		name         string
		urlformvalue string
		expectedCode int
	}{
		{
			name:         "valid",
			urlformvalue: "type=1&name=test_client&markup=1.55&comment=test+client",
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid form value",
			urlformvalue: "type=&name=test_client&markup=1.55&comment=test+client",
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b := strings.NewReader(tc.urlformvalue)
			req, _ := http.NewRequest(http.MethodPost, "/u/create_client", b)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			rec := httptest.NewRecorder()
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
