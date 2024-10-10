package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWafQpsRequest Request Object
type ShowWafQpsRequest struct {

	// 枚举值：yesterday,today,3days,1week,1month 与开始结束时间不同时为空
	Recent *string `json:"recent,omitempty"`

	// 查询域名
	Domains *string `json:"domains,omitempty"`

	// 开始时间（毫秒时间戳）
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间（毫秒时间戳）
	EndTime *string `json:"end_time,omitempty"`

	// 防护区域，0-大陆，1-海外
	OverseasType *int32 `json:"overseas_type,omitempty"`
}

func (o ShowWafQpsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWafQpsRequest struct{}"
	}

	return strings.Join([]string{"ShowWafQpsRequest", string(data)}, " ")
}
