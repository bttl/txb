package txb

import (
	"github.com/ivanrave/apido"
)

func init() {
	ctrlDscr := "Get API docs: Swagger spec" +
		"\n\nReturns object: see Swagger2.0 specification"

	ctrlParams := []apido.InParam{}

	ctrlResp := map[string]apido.ApiResponse{}

	ctrlFunc := func(reqParams map[string]string,
		uid int32,
		perms int32) (
		interface{},
		error) {
		
		return swago, nil
	}

	regPath("1", "api", "docs", "get",
		ctrlDscr,
		ctrlParams,
		ctrlResp,
		ctrlFunc)
}
