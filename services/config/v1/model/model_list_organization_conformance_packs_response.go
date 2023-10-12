package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOrganizationConformancePacksResponse Response Object
type ListOrganizationConformancePacksResponse struct {

	// 组织合规规则包查询列表。
	OrganizationConformancePacks *[]OrgConformancePackResponse `json:"organization_conformance_packs,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListOrganizationConformancePacksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOrganizationConformancePacksResponse struct{}"
	}

	return strings.Join([]string{"ListOrganizationConformancePacksResponse", string(data)}, " ")
}
