package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSensitiveOperationSwitchRequestBody 敏感操作开关请求体
type UpdateSensitiveOperationSwitchRequestBody struct {

	// 是否开启
	IsOpen bool `json:"is_open"`
}

func (o UpdateSensitiveOperationSwitchRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSensitiveOperationSwitchRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSensitiveOperationSwitchRequestBody", string(data)}, " ")
}
