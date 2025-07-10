package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OptionEntity 字段选项
type OptionEntity struct {

	// 选项id
	Id *string `json:"id,omitempty"`

	// 选项code值
	Code *string `json:"code,omitempty"`

	// 选项显示名称
	DisplayValue *string `json:"display_value,omitempty"`

	// 选项唯一标识
	Value *string `json:"value,omitempty"`

	// 选项层级。用于区分层级字段的层级
	Level *int32 `json:"level,omitempty"`

	// 选项顺序
	Sequence *int32 `json:"sequence,omitempty"`

	// 父选项id
	ParentId *string `json:"parent_id,omitempty"`
}

func (o OptionEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OptionEntity struct{}"
	}

	return strings.Join([]string{"OptionEntity", string(data)}, " ")
}
