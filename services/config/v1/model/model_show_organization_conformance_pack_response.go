package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOrganizationConformancePackResponse Response Object
type ShowOrganizationConformancePackResponse struct {

	// 组织合规规则包ID。
	OrgConformancePackId *string `json:"org_conformance_pack_id,omitempty"`

	// 组织合规规则包名称。
	OrgConformancePackName *string `json:"org_conformance_pack_name,omitempty"`

	// 组织合规规则包创建者。
	OwnerId *string `json:"owner_id,omitempty"`

	// 组织ID
	OrganizationId *string `json:"organization_id,omitempty"`

	// 组织合规规则包资源唯一标识。
	OrgConformancePackUrn *string `json:"org_conformance_pack_urn,omitempty"`

	// 排除配置合规包的帐号。
	ExcludedAccounts *[]string `json:"excluded_accounts,omitempty"`

	// 合规规则包参数。
	VarsStructure *[]VarsStructure `json:"vars_structure,omitempty"`

	// 组织合规规则包创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 组织合规规则包更新时间。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 预定义合规规则包模板名称。
	TemplateKey *string `json:"template_key,omitempty"`

	// 合规规则包模板OBS地址
	TemplateUri    *string `json:"template_uri,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowOrganizationConformancePackResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOrganizationConformancePackResponse struct{}"
	}

	return strings.Join([]string{"ShowOrganizationConformancePackResponse", string(data)}, " ")
}
