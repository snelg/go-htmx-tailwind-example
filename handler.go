package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

type HtmlHandler func(r *http.Request) *Response

func (h HtmlHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h(r).Write(rw)
}

func HTML(template string, data interface{}) *Response {

	//render template to buffer
	var buf bytes.Buffer
	if err := templates.ExecuteTemplate(&buf, template, data); err != nil {
		log.Println(err)
		return &Response{Status: http.StatusInternalServerError}
	}
	return &Response{
		Status:      http.StatusOK,
		ContentType: "text/html",
		Content:     &buf,
	}
}

func (response *Response) Write(rw http.ResponseWriter) {
	if response != nil {
		if response.ContentType != "" {
			rw.Header().Set("Content-Type", response.ContentType)
		}
		rw.WriteHeader(response.Status)
		_, err := io.Copy(rw, response.Content)

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		rw.WriteHeader(http.StatusOK)
	}
}

// Response is a generic response object for our handlers
type Response struct {
	// StatusCode
	Status int
	// Content Type to writer
	ContentType string
	// Content to be written to the response writer
	Content io.Reader
}

type Headers map[string]string
