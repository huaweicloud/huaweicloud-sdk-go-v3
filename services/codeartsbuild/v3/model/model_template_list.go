package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TemplateList 模板列表
type TemplateList struct {

	// 模版支持的语言
	Language *string `json:"language,omitempty"`

	// 是否收藏模板
	Favourite *bool `json:"favourite,omitempty"`

	// uuid
	Uuid *string `json:"uuid,omitempty"`

	// 模板类别
	Type *string `json:"type,omitempty"`

	// 模板命名
	Name *string `json:"name,omitempty"`

	// 权重
	Weight *float64 `json:"weight,omitempty"`

	// 模板范围，自定义模板默认为：custom,官方模版为：official
	Scope *string `json:"scope,omitempty"`

	// 模板说明
	Description *string `json:"description,omitempty"`

	// 昵称
	NickName *string `json:"nick_name,omitempty"`

	// 模板是否公开
	Public *bool `json:"public,omitempty"`

	// 构建工具类型，mono、npm、maven等
	ToolType *string `json:"tool_type,omitempty"`

	Template *QueryTemplate `json:"template,omitempty"`

	// 构建执行参数列表
	Parameters *[]Parameter `json:"parameters,omitempty"`
}

func (o TemplateList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateList struct{}"
	}

	return strings.Join([]string{"TemplateList", string(data)}, " ")
}
