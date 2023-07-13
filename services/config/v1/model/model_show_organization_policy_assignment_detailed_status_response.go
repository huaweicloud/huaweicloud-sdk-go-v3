package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOrganizationPolicyAssignmentDetailedStatusResponse Response Object
type ShowOrganizationPolicyAssignmentDetailedStatusResponse struct {

	// 组织合规规则部署详细状态结果列表。
	Value *[]OrganizationPolicyAssignmentDetailedStatusResponse `json:"value,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowOrganizationPolicyAssignmentDetailedStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOrganizationPolicyAssignmentDetailedStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowOrganizationPolicyAssignmentDetailedStatusResponse", string(data)}, " ")
}
