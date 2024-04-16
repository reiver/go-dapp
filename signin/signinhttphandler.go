package dappsignin

import (
	"crypto/rand"
	"fmt"
	_ "embed"
	"io"
	"net/http"
	"strings"
	"time"

//	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/reiver/go-httprequestpath"
	"github.com/reiver/go-path"
)

var replacemekey string = "REPLACE_ME_CHALLENGE"

//go:embed signin-webpage.html
var webpage string

var statusTextInternalServerError string = http.StatusText(http.StatusInternalServerError)
var statusTextMethodNotAllowed    string = http.StatusText(http.StatusMethodNotAllowed)
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
		switch request.Method {
		case http.MethodGet:
			receiver.ServeGET(responseWriter, request)
			return
		case http.MethodPost:
			receiver.ServePOST(responseWriter, request)
			return
		default:
			http.Error(responseWriter, statusTextMethodNotAllowed, http.StatusMethodNotAllowed)
			return
		}
	}
}

func (receiver internalHTTPHandler) ServeGET(responseWriter http.ResponseWriter, request *http.Request) {
	if nil == responseWriter {
		return
	}
	if nil == request {
		http.Error(responseWriter, statusTextInternalServerError, http.StatusInternalServerError)
		return
	}

	{
		if len(webpage) <= 0 {
			http.Error(responseWriter, statusTextInternalServerError, http.StatusInternalServerError)
			return
		}

		var nonce [25]byte
		_, err :=  rand.Read(nonce[:])
		if nil != err {
			http.Error(responseWriter, statusTextInternalServerError, http.StatusInternalServerError)
			return
		}

		var replacemevalue string = fmt.Sprintf(
				"%s wants you to sign in with your Ethereum account:"+"\n"+
				"%s"+"\n"+
				""+"\n"+
				"I accept the %s Terms of Service: https://%s/tos"+"\n"+
				""+"\n"+
				"URI: https://${domain}"+"\n"+
				"Version: 1"+"\n"+
				"Chain ID: 1"+"\n"+
				"Nonce: %d"+"\n"+
				"Issued At: 2021-09-30T16:25:24.000Z"
			internetDomain,
			address,
			internetDomain,
			internetDomain,
			,
			time.Format("2006-01-02T15:04:05.999Z"),

		replacemevalue = fmt.Sprintf("0x%x", replacemevalue)

		var replaced string = strings.ReplaceAll(webpage, replacemekey, replacemevalue)

		responseWriter.Header().Add("Cache-Control", "no-cache")
		io.WriteString(responseWriter, replaced)
	}
}

func (receiver internalHTTPHandler) ServePOST(responseWriter http.ResponseWriter, request *http.Request) {
	if nil == responseWriter {
		return
	}
	if nil == request {
		http.Error(responseWriter, statusTextInternalServerError, http.StatusInternalServerError)
		return
	}

//	{
//		var verified bool = ethcrypto.VerifySignature()
//	}

			
}
