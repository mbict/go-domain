package domain

import (
	"encoding/json"
	"net/http"
)

type Error interface {
	error
	Domain() string
	Aggregate() string
}

type domainError struct {
	code      int
	message   string
	domain    string
	aggregate string
	metadata  map[string]interface{}
}

func (e *domainError) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.message)
}

func (e *domainError) Error() string {
	return e.message
}

// StatusCode should return the http status code used to determine the error
func (e *domainError) StatusCode() int {
	return e.code
}

func (e *domainError) Domain() string {
	return e.domain
}

func (e *domainError) Aggregate() string {
	return e.aggregate
}

func (e *domainError) Metadata() map[string]interface{} {
	return e.metadata
}

func (e *domainError) With(key string, value interface{}) Error {
	metadata := make(map[string]interface{}, len(e.metadata)+1)
	for k, v := range e.metadata {
		metadata[k] = v
	}
	metadata[key] = value

	return &domainError{
		code:      e.code,
		message:   e.message,
		domain:    e.domain,
		aggregate: e.aggregate,
		metadata:  metadata,
	}
}

func NewError(message string, domain string) Error {
	return NewErrorWithCode(message, domain, http.StatusBadRequest)
}

func NewErrorWithCode(message string, domain string, code int) Error {
	return &domainError{
		code:    code,
		message: message,
		domain:  domain,
	}
}

func NewNotFoundError(message string, domain string) Error {
	return NewErrorWithCode(message, domain, http.StatusNotFound)
}
