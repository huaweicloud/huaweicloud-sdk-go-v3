package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSlowlogSensitiveSwitchRequest Request Object
type UpdateSlowlogSensitiveSwitchRequest struct {

	// 内容类型。  取值：application/json。
	ContentType string `json:"Content-Type"`

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *UpdateSlowlogSensitiveSwitchRequestBody `json:"body,omitempty"`
}

func (o UpdateSlowlogSensitiveSwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSlowlogSensitiveSwitchRequest struct{}"
	}

	return strings.Join([]string{"UpdateSlowlogSensitiveSwitchRequest", string(data)}, " ")
}
