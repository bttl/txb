package txb

import (
	"strconv"
	
	"github.com/ivanrave/apido"
	"github.com/bttl/dba"
	"github.com/ivanrave/apphandler"
)

func init() {

	dscr := "Get a text by id"

	params := []apido.InParam{
		apido.InParam{
			In:          "query",
			Name:        "id",	
			Description: "Material id",
			Required:    true,
			SwagType:    "integer",
			SwagFormat:  "int32",
		},
		apido.InParam{
			In:          "query",
			Name:        "vk_id",	
			Description: "VK id",
			Required:    true,
			SwagType:    "integer",
			SwagFormat:  "int32",
		},
		apido.InParam{
			In:          "query",
			Name:        "auth_key",
			Description: "Auth key",
			Required:    true,
			SwagType:    "string",
		},
	}

	resp := map[string]apido.ApiResponse{}

	mthd := func(reqParams map[string]string,
		uid int32,
		perms int32) (interface{}, error) {
			fprm, validCondScope := apido.CheckReq(params,
				reqParams)

			if len(validCondScope) > 0 {
				return nil, apphandler.ErrValidation(validCondScope)
			}

			materialId := fprm["id"].(int32)
			vkId := fprm["vk_id"].(int32)
			authKey := fprm["auth_key"].(string)

			vkStr := strconv.Itoa(int(vkId))

			authStamp := computeChecksum(vkStr, APP_ID, APP_SKEY)

			if authKey != authStamp {
				return nil, apphandler.ErrPerms(0, 1)
			}

			text, errText := dba.GetText(materialId)

			if errText != nil {
				return nil, errText
			}

			return text, nil
		}

	regPath("1", "material", "get-text", "get",
		dscr, params, resp, mthd)
}
