package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourcePoolVo 实际的数据类型：单个对象，集合 或 NULL
type ResourcePoolVo struct {

	// 资源池ID
	Id *string `json:"id,omitempty"`

	// 资源池名称
	Name *string `json:"name,omitempty"`

	// 资源池类型（VM/DOCKER）
	Type *string `json:"type,omitempty"`

	// 是否选中
	Selected *string `json:"selected,omitempty"`

	// 资源池状态
	ActiveState *string `json:"active_state,omitempty"`
}

func (o ResourcePoolVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourcePoolVo struct{}"
	}

	return strings.Join([]string{"ResourcePoolVo", string(data)}, " ")
}
