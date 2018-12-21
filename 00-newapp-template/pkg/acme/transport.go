package acme

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

var tr = &http.Transport{
	MaxIdleConns:    20,
	IdleConnTimeout: 30 * time.Second,
}

type Transport struct {
	BaseUrl     string
	AccessKey   string
	SecretKey   string
	WorkerCount int
	ThreadSafe  *sync.Mutex
}

func NewTransport(s *Service) (p Transport) {
	p.BaseUrl = s.BaseURL
	p.AccessKey = s.AccessKey
	p.SecretKey = s.SecretKey
	p.ThreadSafe = new(sync.Mutex)
	return
}

// xHeader inserts the AccessKey and SecretKey into the request header.
// AccessKey/SecretKey may be equally lengthed comma separated values that are rotated through each call.
// xHeaderCallCount is thread-safely incremented allowing multiple-requests from multiple-credentials (access/secret keys.)
var xHeaderCallCount int

func (t *Transport) xHeader() (header string) {
	akeys := strings.Split(t.AccessKey, ",")
	skeys := strings.Split(t.SecretKey, ",")

	if len(akeys) != len(skeys) {
		return
	}

	// Ensure incremental non-overalapping count
	t.ThreadSafe.Lock()
	xHeaderCallCount = xHeaderCallCount + 1
	mod := xHeaderCallCount % len(akeys)
	t.ThreadSafe.Unlock()

	header = fmt.Sprintf("AccessKey=%s;SecretKey=%s", akeys[mod], skeys[mod])
	return
}
func (t *Transport) GET(url string) (body []byte, err error) {
	var req *http.Request
	var resp *http.Response

	client := &http.Client{Transport: tr}

	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	req.Header.Add("X-ApiKeys", t.xHeader())

	resp, err = client.Do(req) // <-------HTTPS GET Request!
	if err != nil {
		return
	}

	if resp.StatusCode == 429 {
		err = errors.New("error: we need to slow down")
		return
	}
	if resp.StatusCode == 403 {
		err = errors.New("error: creds no longer authorized")
		return
	}

	if resp.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("error: status code does not appear successful: %d", resp.StatusCode))
		return
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)

	return
}
func (t *Transport) POST(url string, data string, datatype string) (body []byte, err error) {
	var req *http.Request
	var resp *http.Response

	client := &http.Client{Transport: tr}

	req, err = http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return
	}
	req.Header.Add("X-ApiKeys", t.xHeader())
	req.Header.Set("Content-Type", datatype)

	resp, err = client.Do(req) // <-------HTTPS GET Request!
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)

	return
}
func (t *Transport) PUT(url string) (body []byte, err error) {
	var req *http.Request
	var resp *http.Response

	client := &http.Client{Transport: tr}

	req, err = http.NewRequest("PUT", url, nil)
	if err != nil {
		return
	}

	req.Header.Add("X-ApiKeys", t.xHeader())

	resp, err = client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)

	return
}
func (t *Transport) DELETE(url string) (body []byte, err error) {
	var req *http.Request
	var resp *http.Response

	client := &http.Client{Transport: tr}

	req, err = http.NewRequest("DELETE", url, nil)
	if err != nil {
		return
	}

	req.Header.Add("X-ApiKeys", t.xHeader())

	resp, err = client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)

	return
}
