package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepositoryTemplateVo2 struct {
	// 模板类型

	TemplateType string `json:"template_type"`
	// 代码模板名称

	CodeTitle *string `json:"code_title,omitempty"`
	// 创建者名称

	CreatorName *string `json:"creator_name,omitempty"`
	// 代码模板描述

	CodeDescription *string `json:"code_description,omitempty"`
	// 模板语言

	Languages *[]string `json:"languages,omitempty"`
	// 模板平台

	Plateform *[]string `json:"plateform,omitempty"`
	// 模板类型

	Entertype *[]string `json:"entertype,omitempty"`
}

func (o RepositoryTemplateVo2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryTemplateVo2 struct{}"
	}

	return strings.Join([]string{"RepositoryTemplateVo2", string(data)}, " ")
}
