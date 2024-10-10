package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWafQpsRequest Request Object
type ListWafQpsRequest struct {

	// 不传时代表全部域名
	Domains *string `json:"domains,omitempty"`

	// 用于 QPS、带宽: 平均值 mean、峰值 peak;用于 响应状态码: 源站返回值 source 、高防返回值 proxy
	ValueType string `json:"value_type"`

	// 开始时间（毫秒时间戳）
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间（毫秒时间戳）
	EndTime *string `json:"end_time,omitempty"`

	// 枚举值：yesterday,today,3days,1week,1month 与开始结束时间不同时为空
	Recent *string `json:"recent,omitempty"`

	// 实例类型，0-大陆，1-海外
	OverseasType *int32 `json:"overseas_type,omitempty"`
}

func (o ListWafQpsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWafQpsRequest struct{}"
	}

	return strings.Join([]string{"ListWafQpsRequest", string(data)}, " ")
}
