package httpclient

import (
	"io"
	"log"
	"net/http"
	"strings"
)

type Header struct {
	Key   string
	Value string
}

type RequestDetails struct {
	URL     string
	Method  string
	Headers []Header
	Query   map[string]string
	Body    io.Reader
}

func Send(rd RequestDetails) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(rd.Method, rd.URL, rd.Body)
	q := req.URL.Query()
	for key, value := range rd.Query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()
	if err != nil {
		log.Println("[Error] Create request to "+rd.URL+" failed, err:", err)
		return nil, err
	}
	for _, header := range rd.Headers {
		req.Header.Add(header.Key, header.Value)
	}
	return client.Do(req)
}

// Extract authorization value from request header
func ExtractAuthHeader(r *http.Request, authType string) string {
	authHeader := r.Header.Get("Authorization")
	auth := strings.Split(authHeader, " ")
	if len(auth) == 2 && auth[0] == authType {
		return auth[1]
	}
	return ""
}
