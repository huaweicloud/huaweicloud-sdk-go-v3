package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SwitchSlowlogDesensitizationRequestBody struct {

	// 实例慢日志脱敏开关开启状态，取值： - off 关闭
	DesensitizationStatus string `json:"desensitization_status"`
}

func (o SwitchSlowlogDesensitizationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchSlowlogDesensitizationRequestBody struct{}"
	}

	return strings.Join([]string{"SwitchSlowlogDesensitizationRequestBody", string(data)}, " ")
}
