package txb

import (
	"github.com/ivanrave/apido"
)

func initApiSpec(curVersion string) apido.ApiSpec {
	return apido.ApiSpec{
		Swagger: "2.0",
		//If the host is not included, the host serving the documentation is to be used (including the port)
		Host: "",
		Info: apido.ApiInfo{
			Title:          "API for bttl",
			Description:    "Methods to get and put information",
			TermsOfService: "http://example.com",
			Contact: map[string]string{
				"name": "Master team",
				"url":  "http://example.com",
			},
			License: map[string]string{
				"name": "Apache License, Version 2.0",
				"url":  "http://www.apache.org/licenses/LICENSE-2.0",
			},
			Version: curVersion,
		},
		BasePath: "/v" + curVersion,
		// "/master/get_by_rubric?id=1&sdf=123"
		Paths:    apido.ApiPaths{},
		// If the schemes is not included, the default scheme to be used is the one used to access the Swagger definition itself.
		// conf.APP_PROTO
		Schemes:  []string{},
		Consumes: []string{"application/x-www-form-urlencoded", "multipart/form-data"},
		Produces: []string{"application/json"},
		// https://github.com/swagger-api/swagger-spec/blob/master/versions/2.0.md#definitions
		Definitions: map[string]apido.ApiDefinition{
			// "serv_group": apido.ApiDefinition{
			// 	Title: "A group of rubrics",
			// 	//Required: []string{"id", "name"},
			// 	Properties: apido.ToSwag(mddt.ServGroup{}),
			// },
		},
		SecurityDefinitions: map[string]apido.SecurityScheme{
			"api_key": apido.SecurityScheme{
				ScrType: "apiKey",
				// The name of the header or query parameter to be used.
				Name: "api_key",
				In:   "header",
			},
		},
		Security: []apido.ScrRequirement{
			apido.ScrRequirement{
				"api_key": []string{},
			},
		},
	}
}
