package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateFunctionAppRequestBody struct {

	// 应用名称
	Name *string `json:"name,omitempty"`

	// 应用使用的模板ID
	TemplateId *string `json:"template_id,omitempty"`

	// 应用描述
	Description *string `json:"description,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 委托名称
	AgencyName *string `json:"agency_name,omitempty"`

	// 模板参数
	Params map[string]string `json:"params,omitempty"`
}

func (o CreateFunctionAppRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFunctionAppRequestBody struct{}"
	}

	return strings.Join([]string{"CreateFunctionAppRequestBody", string(data)}, " ")
}
