package model

import (
	"encoding/json"

	"strings"
)

// 查询配额接口响应结构体
type ResourcesResp struct {
	// 资源列表

	Resources *[]Quotas `json:"resources,omitempty"`
}

func (o ResourcesResp) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ResourcesResp struct{}"
	}

	return strings.Join([]string{"ResourcesResp", string(data)}, " ")
}
