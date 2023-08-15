package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepositoryTemplateVo struct {

	// 模板类型
	TemplateType string `json:"templateType"`

	// 代码模板名称
	CodeTitle *string `json:"codeTitle,omitempty"`

	// 创建者名称
	CreatorName *string `json:"creatorName,omitempty"`

	// 代码模板描述
	CodeDescription *string `json:"codeDescription,omitempty"`

	// 模板语言
	Languages *[]string `json:"languages,omitempty"`

	// 模板平台
	Plateform *[]string `json:"plateform,omitempty"`

	// 模板类型
	Entertype *[]string `json:"entertype,omitempty"`
}

func (o RepositoryTemplateVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryTemplateVo struct{}"
	}

	return strings.Join([]string{"RepositoryTemplateVo", string(data)}, " ")
}
