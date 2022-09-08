package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Property struct {

	// 属性名称
	Name *string `json:"name,omitempty"`

	// 属性值类型
	Type *string `json:"type,omitempty"`

	// 属性描述
	Description *string `json:"description,omitempty"`

	// 属性单位
	Unit *string `json:"unit,omitempty"`
}

func (o Property) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Property struct{}"
	}

	return strings.Join([]string{"Property", string(data)}, " ")
}
