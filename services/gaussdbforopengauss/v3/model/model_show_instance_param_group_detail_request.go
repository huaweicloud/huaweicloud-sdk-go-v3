package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceParamGroupDetailRequest Request Object
type ShowInstanceParamGroupDetailRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowInstanceParamGroupDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceParamGroupDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceParamGroupDetailRequest", string(data)}, " ")
}
