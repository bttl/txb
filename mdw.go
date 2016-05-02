package txb

import (
	"net/http"
	"github.com/ivanrave/apphandler"
)

// Compile the expression once, usually at init time.
// Use raw strings to avoid having to quote the backslashes.
//var validJsonKey = regexp.MustCompile(`^[a-zA-Z]\w+$`)
// ^[a-z]+\[[0-9]+\]$

// beautyMap extract first values from params generated from url.
// Usually url.form returns {val:[123]} instead {val: 123}
// prm1=qwert&prm2=zxcvb to {"prm1":["qwert"],"prm2":["zxcvb"]}
// prm1=awegaw&prm1=ikmikm to {"prm1":["awegaw","ikmikm"]}
//
// In our app forbidden send arrays through url.
// Use underscore-delimited params instead arrays, like {ids: 123_345}
// 
// If a user send {id:123,id:345} - first param will be allowed.
//
// If a user send {id:} - param is ommited
func beautyMap(qwe map[string][]string) (map[string]string) {
	res := make(map[string]string)

	for k, v := range qwe {
	    // if validJsonKey.MatchString(k) != true {
	    //   return nil, fmt.Errorf("json key is unsupported: %v", k)    
	    // }
	    if (len(v) > 0) {
			if v[0] != "" {
			    res[k] = v[0]
			}
	    }
	}

	return res
}


func mdw(ah appHandlerType,
	w http.ResponseWriter,
	r *http.Request) (int16,
		string, // url
		string, // err_key
		string){
	// access only using proxy from nginx
	// headers moved to nginx
	//addAccessControl(w, r)
	
	// Stop here if its Preflighted OPTIONS request
	// move to nginx
	//if r.Method == "OPTIONS" {
	//	return
	//}

	// 2. PARSE middleware
	parseErr := r.ParseForm()
	if parseErr != nil {
		return apphandler.HandleServerError(r, w, parseErr)
	}
	
	// body or url params - body in priority
	inParams := beautyMap(r.Form)
	
	// Returned format stores in url parameter
	// (not in Accept header)
	//acceptTypes := r.Header["Accept"]

	
	// Log a request

	// 3. AUTH middleware
	// - check auth-token
	// - transform it to permScope (if exists)
	// - send permScope to next middleware
	// This MDW doesn't sends AuthToken to next MDW (only perms)
	
	// TODO: #33! Convert r.Form to normal map (without arrays)
	// all arrays will be sended, using _ divider in one param
	// like favarr=123_234_2345&otherparam=123	
	
	// apiKey := hpr.CalcBearerKey(r.Header.Get("Authorization"));
	// authKey := 

	// user id: calc from apiKey
	var uid int32 = 0
	var perms int32 = 0
	// // Check apiKey for all requests (even for non-authed)
	// // a client sends a token only for authed requests
	// if apiKey != "" {

	// 	tkn, err := jauth.CheckToken(apiKey)
	// 		//jwt.Parse(apiKey, cbkJwtParse)

	// 	// Check expired time: automatically inside jwt library
	// 	// https://github.com/dgrijalva/jwt-go/blob/master/jwt.go#L140
	// 	if err != nil {
	// 		if err.Error() == "token is expired" {
	// 			return apphandler.HandleNonAuth(r, w,
	// 				"authTokenIsExpired", apiKey)
	// 		}
			
	// 		return apphandler.HandleServerError(r, w, err)
	// 	}


	// 	if tkn.Valid == false {
	// 		// handle 401 response
	// 		return apphandler.HandleNonAuth(r, w,
	// 			"authTokenIsInvalid", apiKey)
	// 	}

	// 	if uidFloat, isUid := tkn.Claims["uid"].(float64);
	// 	isUid == false {	
	// 		return apphandler.HandleNonAuth(r, w, "authTokenUidIsEmpty", apiKey)
	// 	} else {
	// 		// only int32 supported for UID
	// 		uid = int32(uidFloat)
	// 	}

	// 	if permsFloat, isPerms := tkn.Claims["perms"].(float64);
	// 	isPerms == false {
	// 		return apphandler.HandleNonAuth(r, w,
	// 			"authTokenPermsIsEmpty", apiKey)
	// 	} else {
	// 		// only int32 supported for PERMS
	// 		perms = int32(permsFloat)
	// 	}
		
	// 	//parts := strings.Split(tkn.Raw, ".")

	// 	//dcd, _ := DecodeSegment(parts[1])


	// 	// exp: int64 unixtimestamp
	// }
	
	// translate apiKey to userId + perms (roles)
	// from JWT or TableSession
	// define perms from DB or JWT?
	// Are perms can be different per sessions?
	// To get perms from DB - DB perms required (cycle)
	
	
	// Execute required function
	// If errors occured - return error to client
	// w, r - convert to data for controllers
	// 4. MAIN middleware
	// - check permScope, if required
	// - check inParams
	// - execute main methods
	// ? nil or int32 = 0


	// 5. RESULT middleware
	if 	rdata, errApp := ah(inParams, uid, perms);
	errApp != nil {
		if clerrApp, ok := apphandler.ToClerr(errApp); ok {
			//errApp.(*clerr); ok {
			if clerrApp.ErrKey == "permissionError" {
				// 401
				return apphandler.HandleNonAuth(r, w, "notEnoughPermissions", "")
			} else if clerrApp.ErrKey == "notFoundError" {
				// 404
				return apphandler.HandleNotFound(r, w)
			} else {
				// 422
				return apphandler.HandleClientError(r, w, clerrApp)
			}
		}else {
			// 500
			return apphandler.HandleServerError(r, w, errApp)
		}
	} else {
		// 200, 204
		return apphandler.HandleSuccess(r, w, rdata)
	}
}
