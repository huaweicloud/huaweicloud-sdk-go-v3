package model

import (
	"encoding/json"

	"strings"
)

type DomainItemLocationDetails struct {
	// 数据起始时间戳，可能与请求时间不一致，可能不返回

	StartTime *int64 `json:"start_time,omitempty"`
	// 数据结束时间戳，可能与请求时间不一致，可能不返回

	EndTime *int64 `json:"end_time,omitempty"`
	// 指标类型，可能不返回

	StatType *string `json:"stat_type,omitempty"`
	// 域名详情数据列表

	Domains *[]DomainRegion `json:"domains,omitempty"`
}

func (o DomainItemLocationDetails) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DomainItemLocationDetails struct{}"
	}

	return strings.Join([]string{"DomainItemLocationDetails", string(data)}, " ")
}
