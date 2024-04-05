package dappconnect

import (
	_ "embed"
	"net/http"

	"github.com/reiver/go-httprequestpath"
	"github.com/reiver/go-path"
)

//go:embed webpage.html
var webpage []byte

var statusTextInternalServerError string = http.StatusText(http.StatusInternalServerError)
var statusTextNotFound            string = http.StatusText(http.StatusNotFound)

func HTTPHandler(httpRequestPath string) http.Handler {
	httpRequestPath = path.Canonical(httpRequestPath)

	return internalHTTPHandler{
		path:httpRequestPath,
	}
}

type internalHTTPHandler struct {
	path string
}

func (receiver internalHTTPHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	if nil == responseWriter {
		return
	}
	if nil == request {
		http.Error(responseWriter, statusTextInternalServerError, http.StatusInternalServerError)
		return
	}

	var httpRequestPath string
	{
		var err error

		httpRequestPath, err = httprequestpath.HTTPRequestPath(request)
		if nil != err {
			http.Error(responseWriter, statusTextInternalServerError, http.StatusInternalServerError)
			return
		}

		httpRequestPath = path.Canonical(httpRequestPath)
	}

	{
		if httpRequestPath != receiver.path {
			http.Error(responseWriter, statusTextNotFound, http.StatusNotFound)
			return
		}
	}

	{
		if len(webpage) <= 0 {
			http.Error(responseWriter, statusTextInternalServerError, http.StatusInternalServerError)
			return
		}

		responseWriter.Header().Add("Cache-Control", "no-cache")
		responseWriter.Write(webpage)
	}
}
