txb
===

API endpoints for bttl

Using
---

```
package main

import (
	"log"
	"net/http"
	"github.com/bttl/txb"
)

func main(){
	txb.InitPackage("app_id", "app_skey")

	port := ":56789"
	err := http.ListenAndServe(port, nil)
	if err != nil {
	    log.Fatal(err)
	}	
}
```
