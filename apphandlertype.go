package txb

import (
	"net/http"

	"fmt"
	// "google.golang.org/appengine"
	// "google.golang.org/appengine/log"
	//	"strconv"
	//lgr	"github.com/Sirupsen/logrus"
)

// appHandler
// func as a paramter
// error as a result
// Result - a map (will be converted to json/othertype to send to user
type appHandlerType func(map[string]string, int32, int32) (
	interface{}, error)

// ServeHTTP implements Handler interface
// Param ah - data of appHandlerType instead Handler type
// ServeHTTP should write reply headers and data to the ResponseWriter and then return. Returning signals that the request is finished;
// it is not valid to use the ResponseWriter or read from the Request.Body after or concurrently with the completion of the ServeHTTP call.
// Depending on the HTTP client software, HTTP protocol version, and any intermediaries between the client and the Go server, it may not be possible to read from the Request.Body after writing to the ResponseWriter. Cautious handlers should read the Request.Body first, and then reply.
// If ServeHTTP panics, the server (the caller of ServeHTTP) assumes that the effect of the panic was isolated to the active request. It recovers the panic, logs a stack trace to the server error log, and hangs up the connection.
func (ah appHandlerType) ServeHTTP(w http.ResponseWriter,
	r *http.Request) {

	statusCode,	url, errKey, msg := mdw(ah, w, r)

	//ctx := appengine.NewContext(r)
	//log.Debugf(ctx, msg)
	//mdw(ah, w, r)

	// lgr.WithFields(lgr.Fields {
	// 	"tag": "qui.r" + strconv.Itoa(int(statusCode)),
	// 	"status_code": statusCode,
	// 	"url": url,
	// 	"err_key": errKey,
	// }).Info(msg)

	fmt.Printf("%v %v %v %v", statusCode, url, errKey, msg)
	//log.Info(msg)
}
