package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceParamGroupRequest Request Object
type ShowInstanceParamGroupRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowInstanceParamGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceParamGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceParamGroupRequest", string(data)}, " ")
}
