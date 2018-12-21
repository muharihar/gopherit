package acme

import (
	"bytes"
	"errors"
	"fmt"
	"text/template"
	"time"
)

func (s *Service) SleepBeforeRetry(attempt int) (shouldReRun bool) {
	if attempt < len(s.RetryIntervals) {
		time.Sleep(time.Duration(s.RetryIntervals[attempt]) * time.Millisecond)
		shouldReRun = true
	}
	return
}

func (s *Service) Get(name string, p map[string]string) (error, []byte) {
	url, err := s.ToURL(name, p)
	if err != nil {
		return err, nil
	}
	t := NewTransport(s)
	body, err := t.GET(url)
	if err != nil {
		return err, nil
	}
	return err, body
}
func (s *Service) Delete(name string, p map[string]string) (error, []byte) {
	url, err := s.ToURL(name, p)
	if err != nil {
		return err, nil
	}
	t := NewTransport(s)
	body, err := t.DELETE(url)
	if err != nil {
		return err, nil
	}

	return err, body
}
func (s *Service) Update(name string, p map[string]string) (error, []byte) {
	url, err := s.ToURL(name, p)
	if err != nil {
		return err, nil
	}
	json, err := s.ToJSON(name, p)
	if err != nil {
		return err, nil
	}

	t := NewTransport(s)
	body, err := t.POST(url, json, "application/json")
	if err != nil {
		return err, nil
	}

	return err, body
}

func (s *Service) ToURL(name string, p map[string]string) (url string, err error) {
	if p == nil {
		p = make(map[string]string)
	}
	p["BaseURL"] = s.BaseURL
	return s.ToTemplate(name, p, URLServiceTmpl)
}
func (s *Service) ToJSON(name string, p map[string]string) (url string, err error) {
	return s.ToTemplate(name, p, JSONBodyTmpl)
}

func (s *Service) ToTemplate(name string, p map[string]string, tmap map[string]string) (url string, err error) {
	var rawURL bytes.Buffer
	t, terr := template.New(name).Parse(tmap[name])
	if terr != nil {
		err = errors.New(fmt.Sprintf("error: failed to parse template for %s: %v", name, err))
		return
	}
	err = t.Execute(&rawURL, p)
	if err != nil {
		return
	}

	url = rawURL.String()

	return
}
