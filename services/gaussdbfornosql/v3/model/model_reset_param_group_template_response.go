package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetParamGroupTemplateResponse Response Object
type ResetParamGroupTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResetParamGroupTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetParamGroupTemplateResponse struct{}"
	}

	return strings.Join([]string{"ResetParamGroupTemplateResponse", string(data)}, " ")
}
