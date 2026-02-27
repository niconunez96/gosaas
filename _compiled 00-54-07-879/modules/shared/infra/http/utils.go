package sharedHttp

import "net/http"

type StatusCodeRecorder struct {
	http.ResponseWriter
	StatusCode int
}

func NewStatusCodeRecorder(w http.ResponseWriter) *StatusCodeRecorder {
	return &StatusCodeRecorder{w, http.StatusOK}
}

func (s *StatusCodeRecorder) WriteHeader(code int) {
	s.StatusCode = code
	s.ResponseWriter.WriteHeader(code)
}
