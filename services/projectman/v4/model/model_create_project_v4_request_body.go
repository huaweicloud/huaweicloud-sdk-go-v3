package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateProjectV4RequestBody struct {
	// 项目名称

	ProjectName string `json:"project_name"`
	// 项目描述

	Description *string `json:"description,omitempty"`
	// 项目来源

	Source *string `json:"source,omitempty"`
	// 项目类型 scrum, xboard(看板项目), basic, phoenix(凤凰项目)

	ProjectType string `json:"project_type"`
	// 项目要绑定的企业项目ID

	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// 用户创建的项目模板id

	TemplateId *int32 `json:"template_id,omitempty"`
}

func (o CreateProjectV4RequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProjectV4RequestBody struct{}"
	}

	return strings.Join([]string{"CreateProjectV4RequestBody", string(data)}, " ")
}
