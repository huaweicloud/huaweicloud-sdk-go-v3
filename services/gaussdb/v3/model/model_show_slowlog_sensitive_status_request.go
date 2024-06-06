package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSlowlogSensitiveStatusRequest Request Object
type ShowSlowlogSensitiveStatusRequest struct {

	// 内容类型。  取值：application/json。
	ContentType string `json:"Content-Type"`

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowSlowlogSensitiveStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSlowlogSensitiveStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowSlowlogSensitiveStatusRequest", string(data)}, " ")
}
