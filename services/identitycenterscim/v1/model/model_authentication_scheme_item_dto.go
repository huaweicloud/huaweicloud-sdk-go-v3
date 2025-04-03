package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AuthenticationSchemeItemDto struct {

	// 认证类型，指定鉴权的方式
	Type *string `json:"type,omitempty"`

	// 认证概要名称
	Name *string `json:"name,omitempty"`

	// 认证概要的描述信息
	Description *string `json:"description,omitempty"`

	// 规范链接
	SpecUri *string `json:"specUri,omitempty"`

	// 帮助文档链接
	DocumentationUri *string `json:"documentationUri,omitempty"`

	// 是否为主要的认证方式
	Primary *bool `json:"primary,omitempty"`
}

func (o AuthenticationSchemeItemDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthenticationSchemeItemDto struct{}"
	}

	return strings.Join([]string{"AuthenticationSchemeItemDto", string(data)}, " ")
}
