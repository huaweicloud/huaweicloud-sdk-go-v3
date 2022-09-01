package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepositoryTemplateVo struct {

	// 模板类型
	TemplateType string `json:"templateType" xml:"templateType"`

	// 代码模板名称
	CodeTitle *string `json:"codeTitle,omitempty" xml:"codeTitle"`

	// 创建者名称
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName"`

	// 代码模板描述
	CodeDescription *string `json:"codeDescription,omitempty" xml:"codeDescription"`

	// 模板语言
	Languages *[]string `json:"languages,omitempty" xml:"languages"`

	// 模板平台
	Plateform *[]string `json:"plateform,omitempty" xml:"plateform"`

	// 模板类型
	Entertype *[]string `json:"entertype,omitempty" xml:"entertype"`
}

func (o RepositoryTemplateVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryTemplateVo struct{}"
	}

	return strings.Join([]string{"RepositoryTemplateVo", string(data)}, " ")
}
