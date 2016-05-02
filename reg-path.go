package txb

import (
	"strings"
	"net/http"

	"github.com/ivanrave/apido"
)

func getFirstLine(str string) string {
	return strings.Split(str, "\n")[0]
}

// regPath register a path as
// - an http handler (path + method)
// - a swago doc (path + metadata + input, output params)
func regPath(apiVersion string,
	apiTag string,
	apiOperation string,
	reqType string, // GET, POST only
	description string,
	inParamArr []apido.InParam,
	respMap map[string]apido.ApiResponse,
	ctrlFunc appHandlerType) {

	apiPath := "/" + apiTag + "/" + apiOperation

	// register handler in API
	// http://golang.org/pkg/net/http/#Handle
	// Handle registers the handler for the given pattern in the DefaultServeMux. The documentation for ServeMux explains how patterns are matched
	http.Handle("/v" + apiVersion + apiPath, appHandlerType(ctrlFunc))

	// executes ctrlFunc.ServeHttp

	swago.AppendPath(apiPath,
		reqType,
		getFirstLine(description),
		description,
		[]string {apiTag},
		[]string{"application/x-www-form-urlencoded"},
		[]string{"application/json"},
		inParamArr,
		respMap)
}
