package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWafAttackEventRequest Request Object
type ListWafAttackEventRequest struct {

	// 不传时代表全部域名
	Domains *string `json:"domains,omitempty"`

	// 开始时间（毫秒时间戳）
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间（毫秒时间戳）
	EndTime *string `json:"end_time,omitempty"`

	// 枚举值：yesterday,today,3days,1week,1month 与开始结束时间不同时为空
	Recent *string `json:"recent,omitempty"`

	// 实例类型，0-大陆，1-海外
	OverseasType *int32 `json:"overseas_type,omitempty"`

	// 攻击源IP
	Sip *string `json:"sip,omitempty"`

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// offset
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListWafAttackEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWafAttackEventRequest struct{}"
	}

	return strings.Join([]string{"ListWafAttackEventRequest", string(data)}, " ")
}
