package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataLayerVo struct {

	// 数仓分层的ID，ID字符串。
	Id *string `json:"id,omitempty"`

	// 层级，从1开始。
	Level int32 `json:"level"`

	// 中文名称。
	Name string `json:"name"`

	Type *LayerType `json:"type"`

	// 数仓分层描述。
	Description *string `json:"description,omitempty"`

	// 是否是不可删除的默认分层，包含SDI\\DWR\\DM。
	IsDefault *bool `json:"is_default,omitempty"`

	// 该分层禁用的自定义项的ID列表，包括表级自定义项和字段级自定义项。
	DisabledCustomizedFieldIds *[]string `json:"disabled_customized_field_ids,omitempty"`
}

func (o DataLayerVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataLayerVo struct{}"
	}

	return strings.Join([]string{"DataLayerVo", string(data)}, " ")
}
