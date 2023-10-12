package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOrganizationConformancePackStatusesResponse Response Object
type ListOrganizationConformancePackStatusesResponse struct {

	// 组织合规规则包查询列表。
	Statuses *[]OrgConformancePackStatus `json:"statuses,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListOrganizationConformancePackStatusesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOrganizationConformancePackStatusesResponse struct{}"
	}

	return strings.Join([]string{"ListOrganizationConformancePackStatusesResponse", string(data)}, " ")
}
