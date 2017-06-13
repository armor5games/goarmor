package goarmorapi

import (
	"bytes"
	"io"
	"net/http"

	"fmt"
)

// ResponsePlain requires that the calling party itself set
// `w.Header().Set("Content-type", "text/plain; charset=utf-8")` header
// if the WriteHeader function was called on http.ResponseWriter.
func ResponsePlain(w http.ResponseWriter, r *http.Request, rd io.Reader) error {
	w.Header().Set("Content-type", "text/plain; charset=utf-8")
	b := new(bytes.Buffer)

	_, err := b.ReadFrom(rd)
	if err != nil {
		return fmt.Errorf("bytes.(*Buffer).ReadFrom fn error: %s", err.Error())
	}

	fmt.Fprint(w, b.String())

	return nil
}

// ResponseXML requires that the calling party itself set
// `w.Header().Set("Content-type", "application/xml; charset=utf-8")` header
// if the WriteHeader function was called on http.ResponseWriter.
func ResponseXML(w http.ResponseWriter, r *http.Request, rd io.Reader) error {
	w.Header().Set("Content-type", "application/xml; charset=utf-8")

	b := new(bytes.Buffer)

	_, err := b.ReadFrom(rd)
	if err != nil {
		return fmt.Errorf("bytes.(*Buffer).ReadFrom fn error: %s", err.Error())
	}

	fmt.Fprint(w, b.String())

	return nil
}

// ResponseJSON requires that the calling party itself set
// `w.Header().Set("Content-type", "application/json; charset=utf-8")` header
// if the WriteHeader function was called on http.ResponseWriter.
func ResponseJSON(
	w http.ResponseWriter,
	r *http.Request,
	isSuccess bool,
	responsePayload interface{},
	keyValues KV,
	errs ...*ErrorJSON) error {
	b, err := jsonWithDebug(r.Context(), isSuccess, responsePayload, nil)
	if err != nil {
		return fmt.Errorf("answer.jsonWithDebug fn error: %s", err.Error())
	}

	err = responseRawJSON(w, r, b)
	if err != nil {
		return err
	}

	return nil
}

// ResponseJSONWithDebug requires that the calling party itself set
// `w.Header().Set("Content-type", "application/json; charset=utf-8")` header
// if the WriteHeader function was called on http.ResponseWriter.
func ResponseJSONWithDebug(
	w http.ResponseWriter,
	r *http.Request,
	isSuccess bool,
	responsePayload interface{},
	keyValues KV,
	errs ...*ErrorJSON) error {
	b, err := jsonWithDebug(
		r.Context(), isSuccess, responsePayload, keyValues, errs...)
	if err != nil {
		return fmt.Errorf("answer.jsonWithDebug fn error: %s", err.Error())
	}

	err = responseRawJSON(w, r, b)
	if err != nil {
		return err
	}

	return nil
}

func responseRawJSON(
	w http.ResponseWriter, r *http.Request, rd io.Reader) error {
	w.Header().Set("Content-type", "application/json; charset=utf-8")

	b := new(bytes.Buffer)

	_, err := b.ReadFrom(rd)
	if err != nil {
		return fmt.Errorf("bytes.(*Buffer).ReadFrom fn error: %s", err.Error())
	}

	fmt.Fprint(w, b.String())

	return nil
}