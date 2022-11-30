package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowSlowLogDesensitizationResponse struct {

	// 实例慢日志脱敏开启状态，取值：  - on 开启  - off 关闭
	DesensitizationStatus *string `json:"desensitization_status,omitempty"`
	HttpStatusCode        int     `json:"-"`
}

func (o ShowSlowLogDesensitizationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSlowLogDesensitizationResponse struct{}"
	}

	return strings.Join([]string{"ShowSlowLogDesensitizationResponse", string(data)}, " ")
}
