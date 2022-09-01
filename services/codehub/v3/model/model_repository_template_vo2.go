package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepositoryTemplateVo2 struct {

	// 模板类型
	TemplateType string `json:"template_type" xml:"template_type"`

	// 代码模板名称
	CodeTitle *string `json:"code_title,omitempty" xml:"code_title"`

	// 创建者名称
	CreatorName *string `json:"creator_name,omitempty" xml:"creator_name"`

	// 代码模板描述
	CodeDescription *string `json:"code_description,omitempty" xml:"code_description"`

	// 模板语言
	Languages *[]string `json:"languages,omitempty" xml:"languages"`

	// 模板平台
	Plateform *[]string `json:"plateform,omitempty" xml:"plateform"`

	// 模板类型
	Entertype *[]string `json:"entertype,omitempty" xml:"entertype"`
}

func (o RepositoryTemplateVo2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryTemplateVo2 struct{}"
	}

	return strings.Join([]string{"RepositoryTemplateVo2", string(data)}, " ")
}
