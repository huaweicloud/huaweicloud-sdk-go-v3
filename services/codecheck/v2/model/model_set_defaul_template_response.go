package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SetDefaulTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetDefaulTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetDefaulTemplateResponse struct{}"
	}

	return strings.Join([]string{"SetDefaulTemplateResponse", string(data)}, " ")
}
