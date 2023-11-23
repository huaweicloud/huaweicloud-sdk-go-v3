package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetParamGroupTemplateRequest Request Object
type ResetParamGroupTemplateRequest struct {

	// 参数模板ID
	ConfigId string `json:"config_id"`
}

func (o ResetParamGroupTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetParamGroupTemplateRequest struct{}"
	}

	return strings.Join([]string{"ResetParamGroupTemplateRequest", string(data)}, " ")
}
