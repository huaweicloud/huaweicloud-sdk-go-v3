package model

import (
	"encoding/json"

	"strings"
)

// 配额信息
type QuotaResource struct {
	// 配额类型信息

	Type *string `json:"type,omitempty"`
	// 配额最小取值

	Min *int32 `json:"min,omitempty"`
	// 配额最大取值

	Max *int32 `json:"max,omitempty"`
	// 用户配额的实际值

	Quota *int32 `json:"quota,omitempty"`
	// 已使用的配额值

	Used *int32 `json:"used,omitempty"`
}

func (o QuotaResource) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QuotaResource struct{}"
	}

	return strings.Join([]string{"QuotaResource", string(data)}, " ")
}
