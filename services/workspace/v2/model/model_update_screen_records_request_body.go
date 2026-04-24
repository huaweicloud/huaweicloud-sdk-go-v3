package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScreenRecordsRequestBody 更新录屏审计请求体
type UpdateScreenRecordsRequestBody struct {

	// 录屏限速类型
	TrafficLimitType string `json:"traffic_limit_type"`
}

func (o UpdateScreenRecordsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScreenRecordsRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateScreenRecordsRequestBody", string(data)}, " ")
}
