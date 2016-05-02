// package txb describes API endpoints
package txb

import (
	"github.com/ivanrave/apido"
)

var swago apido.ApiSpec = initApiSpec("1")

// store as global
var APP_ID string = ""
var APP_SKEY string = ""

func InitPackage(appId string, appSkey string){
	APP_ID = appId
	APP_SKEY = appSkey
}
