package dapphttp

import (
	_ "embed"
	"net/http"

	"github.com/reiver/go-httprequestpath"
	"github.com/reiver/go-path"
)

//go:embed connect-webpage.html
var webpage []byte

func ConnecHTTPHandler(httpRequestPath string) http.Handler {
	httpRequestPath = path.Canonical(httpRequestPath)

	return internalConnectHTTPHandler{
		path:httpRequestPath,
	}
}

type internalConnectHTTPHandler struct {
	path string
}

func (receiver internalConnectHTTPHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
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
