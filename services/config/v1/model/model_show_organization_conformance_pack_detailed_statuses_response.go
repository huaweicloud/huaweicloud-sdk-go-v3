package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOrganizationConformancePackDetailedStatusesResponse Response Object
type ShowOrganizationConformancePackDetailedStatusesResponse struct {

	// 组织合规规则包查询列表。
	Statuses *[]OrgConformancePackDetailedStatus `json:"statuses,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowOrganizationConformancePackDetailedStatusesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOrganizationConformancePackDetailedStatusesResponse struct{}"
	}

	return strings.Join([]string{"ShowOrganizationConformancePackDetailedStatusesResponse", string(data)}, " ")
}
