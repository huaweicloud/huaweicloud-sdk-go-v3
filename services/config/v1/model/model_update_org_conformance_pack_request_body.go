package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateOrgConformancePackRequestBody 更新组织合规规则包请求体。
type UpdateOrgConformancePackRequestBody struct {

	// 组织合规规则包名称。
	Name string `json:"name"`

	// 排除配置合规包的帐号。
	ExcludedAccounts *[]string `json:"excluded_accounts,omitempty"`

	// 合规规则包参数。
	VarsStructure *[]VarsStructure `json:"vars_structure,omitempty"`
}

func (o UpdateOrgConformancePackRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOrgConformancePackRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateOrgConformancePackRequestBody", string(data)}, " ")
}
