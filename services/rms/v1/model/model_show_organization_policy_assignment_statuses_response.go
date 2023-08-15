package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOrganizationPolicyAssignmentStatusesResponse Response Object
type ShowOrganizationPolicyAssignmentStatusesResponse struct {

	// 组织合规规则部署状态结果列表。
	Value *[]OrganizationPolicyAssignmentStatusResponse `json:"value,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowOrganizationPolicyAssignmentStatusesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOrganizationPolicyAssignmentStatusesResponse struct{}"
	}

	return strings.Join([]string{"ShowOrganizationPolicyAssignmentStatusesResponse", string(data)}, " ")
}
