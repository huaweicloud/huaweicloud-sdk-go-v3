package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Resource struct {
	// 资源类型，包含instance（服务实例数）

	Type *string `json:"type,omitempty"`
	// 单位

	Unit *string `json:"unit,omitempty"`
	// 最小值

	Min *int64 `json:"min,omitempty"`
	// 最大值

	Max *int64 `json:"max,omitempty"`
	// 配额

	Quota *int64 `json:"quota,omitempty"`
	// 已使用配额

	Used *int64 `json:"used,omitempty"`
	// 剩余配额

	Free *int64 `json:"free,omitempty"`
}

func (o Resource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resource struct{}"
	}

	return strings.Join([]string{"Resource", string(data)}, " ")
}
