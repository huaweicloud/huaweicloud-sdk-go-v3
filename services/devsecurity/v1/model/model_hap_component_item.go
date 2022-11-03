package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HapComponentItem struct {

	// 组件名称
	Name *string `json:"name,omitempty"`

	// 元能力可见
	Visible *bool `json:"visible,omitempty"`

	// 动作和实体列表
	ActionsWithEntities *[]ActionWithEntities `json:"actions_with_entities,omitempty"`
}

func (o HapComponentItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HapComponentItem struct{}"
	}

	return strings.Join([]string{"HapComponentItem", string(data)}, " ")
}
