package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScreenRecordsTrafficLimitConfigRequestBody 录屏记录配置结果请求体。
type UpdateScreenRecordsTrafficLimitConfigRequestBody struct {

	// 录屏记录。
	Configs *[]UpdateScreenRecordsTrafficLimitConfigRequestBodyConfigs `json:"configs,omitempty"`
}

func (o UpdateScreenRecordsTrafficLimitConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScreenRecordsTrafficLimitConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateScreenRecordsTrafficLimitConfigRequestBody", string(data)}, " ")
}
