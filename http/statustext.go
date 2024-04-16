package dapphttp

import (
	"net/http"
)

var statusTextInternalServerError string = http.StatusText(http.StatusInternalServerError)
var statusTextNotFound            string = http.StatusText(http.StatusNotFound)
