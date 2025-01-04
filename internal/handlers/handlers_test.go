package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		name     string
		method   string
		body     interface{}
		expected int
		response interface{}
	}{
		{
			name:     "Valid expression",
			method:   http.MethodPost,
			body:     Request{Expression: "2+3*4"},
			expected: http.StatusOK,
			response: Response{Result: 14},
		},
		{
			name:     "Invalid method",
			method:   http.MethodGet,
			body:     nil,
			expected: http.StatusMethodNotAllowed,
			response: ErrorResponse{Error: "Method not allowed"},
		},
		{
			name:     "Empty expression",
			method:   http.MethodPost,
			body:     Request{Expression: ""},
			expected: http.StatusUnprocessableEntity,
			response: ErrorResponse{Error: "Expression is not valid"},
		},
		{
			name:     "Invalid expression",
			method:   http.MethodPost,
			body:     Request{Expression: "+2+2"},
			expected: http.StatusUnprocessableEntity,
			response: ErrorResponse{Error: "Expression is not valid"},
		},
		{
			name:     "Division by zero",
			method:   http.MethodPost,
			body:     Request{Expression: "2/0"},
			expected: http.StatusInternalServerError,
			response: ErrorResponse{Error: "Internal server error"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var body []byte
			if tt.body != nil {
				body, _ = json.Marshal(tt.body)
			}

			req := httptest.NewRequest(tt.method, "/", bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			Handler(w, req)

			resp := w.Result()
			if resp.StatusCode != tt.expected {
				t.Errorf("expected status %d, got %d", tt.expected, resp.StatusCode)
			}

			if tt.response != nil {
				switch expected := tt.response.(type) {
				case Response:
					var actual Response
					if err := json.NewDecoder(resp.Body).Decode(&actual); err != nil {
						t.Errorf("failed to decode response body: %v", err)
					}
					if actual != expected {
						t.Errorf("expected response %v, got %v", expected, actual)
					}
				case ErrorResponse:
					var actual ErrorResponse
					if err := json.NewDecoder(resp.Body).Decode(&actual); err != nil {
						t.Errorf("failed to decode response body: %v", err)
					}
					if actual != expected {
						t.Errorf("expected response %v, got %v", expected, actual)
					}
				default:
					t.Errorf("unsupported response type")
				}
			}
		})
	}
}
