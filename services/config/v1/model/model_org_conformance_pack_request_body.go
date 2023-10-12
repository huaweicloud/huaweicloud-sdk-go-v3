package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OrgConformancePackRequestBody 组织合规规则包请求体。
type OrgConformancePackRequestBody struct {

	// 组织合规规则包名称。
	Name string `json:"name"`

	// 排除配置合规包的帐号。
	ExcludedAccounts *[]string `json:"excluded_accounts,omitempty"`

	// 预定义合规规则包模板名称。
	TemplateKey *string `json:"template_key,omitempty"`

	// 自定义合规规则包内容。
	TemplateBody *string `json:"template_body,omitempty"`

	// 合规规则包模板OBS地址。
	TemplateUri *string `json:"template_uri,omitempty"`

	// 合规规则包参数。
	VarsStructure *[]VarsStructure `json:"vars_structure,omitempty"`
}

func (o OrgConformancePackRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrgConformancePackRequestBody struct{}"
	}

	return strings.Join([]string{"OrgConformancePackRequestBody", string(data)}, " ")
}
