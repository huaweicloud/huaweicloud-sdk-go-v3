package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateOrganizationConformancePackResponse Response Object
type UpdateOrganizationConformancePackResponse struct {

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

	// 合规规则包参数。
	VarsStructure *[]VarsStructure `json:"vars_structure,omitempty"`

	// 组织合规规则包创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 组织合规规则包更新时间。
	UpdatedAt      *string `json:"updated_at,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateOrganizationConformancePackResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOrganizationConformancePackResponse struct{}"
	}

	return strings.Join([]string{"UpdateOrganizationConformancePackResponse", string(data)}, " ")
}
