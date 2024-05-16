package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BusinessEnvRequest 获取Region环境入参。
type BusinessEnvRequest struct {

	// 应用id。
	BusinessId int64 `json:"business_id"`

	// region英文名称。
	Region *string `json:"region,omitempty"`

	// 开始时间。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间。
	EndTime *string `json:"end_time,omitempty"`
}

func (o BusinessEnvRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BusinessEnvRequest struct{}"
	}

	return strings.Join([]string{"BusinessEnvRequest", string(data)}, " ")
}
