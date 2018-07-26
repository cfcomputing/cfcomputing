package api

import (
	"io"
	"net/http"
)

// hat tip to https://blog.rapid7.com/2016/07/13/quick-security-wins-in-golang/

// API is a thing
type API struct{}

// ServeHTTP is good enough
func (a *API) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "api\n")
}
