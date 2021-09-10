package model

import (
	"encoding/json"

	"strings"
)

type DomainItemDetail struct {
	// 数据起始时间戳，可能与请求时间不一致，可能不返回

	StartTime *int64 `json:"start_time,omitempty"`
	// 数据结束时间戳，可能与请求时间不一致，可能不返回

	EndTime *int64 `json:"end_time,omitempty"`
	// 指标类型，可能不返回

	StatType *string `json:"stat_type,omitempty"`
	// 数据结束时间戳，可能与请求时间不一致，可能不返回

	Domains *[]map[string]interface{} `json:"domains,omitempty"`
}

func (o DomainItemDetail) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DomainItemDetail struct{}"
	}

	return strings.Join([]string{"DomainItemDetail", string(data)}, " ")
}
