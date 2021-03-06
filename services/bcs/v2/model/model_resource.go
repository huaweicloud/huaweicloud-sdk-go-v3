package model

import (
	"encoding/json"

	"strings"
)

type Resource struct {
	// 资源类型，包含instance（服务实例数）和peers（总peer数）两种

	Type *string `json:"type,omitempty"`
	// 单位

	Unit *string `json:"unit,omitempty"`
	// 最小值

	Min *int32 `json:"min,omitempty"`
	// 最大值

	Max *int32 `json:"max,omitempty"`
	// 配额

	Quota *int32 `json:"quota,omitempty"`
	// 已使用配额

	Used *int32 `json:"used,omitempty"`
	// 剩余配额

	Free *int32 `json:"free,omitempty"`
}

func (o Resource) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Resource struct{}"
	}

	return strings.Join([]string{"Resource", string(data)}, " ")
}
