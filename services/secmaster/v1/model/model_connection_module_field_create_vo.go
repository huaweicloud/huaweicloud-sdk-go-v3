package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConnectionModuleFieldCreateVo struct {

	// 名字
	Name *string `json:"name,omitempty"`

	// 其他信息
	Other *string `json:"other,omitempty"`

	// UUID
	TemplateFieldId *string `json:"template_field_id,omitempty"`

	// 标题
	Title *string `json:"title,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 相关值
	Value *string `json:"value,omitempty"`
}

func (o ConnectionModuleFieldCreateVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectionModuleFieldCreateVo struct{}"
	}

	return strings.Join([]string{"ConnectionModuleFieldCreateVo", string(data)}, " ")
}
