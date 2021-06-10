package model

import (
	"encoding/json"

	"strings"
)

// 配额信息
type QueryQuotaInfo struct {
	Resource *QuotaResource `json:"resource,omitempty"`
}

func (o QueryQuotaInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QueryQuotaInfo struct{}"
	}

	return strings.Join([]string{"QueryQuotaInfo", string(data)}, " ")
}
