package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOrganizationConformancePackRequest Request Object
type CreateOrganizationConformancePackRequest struct {

	// 组织ID。
	OrganizationId string `json:"organization_id"`

	Body *OrgConformancePackRequestBody `json:"body,omitempty"`
}

func (o CreateOrganizationConformancePackRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrganizationConformancePackRequest struct{}"
	}

	return strings.Join([]string{"CreateOrganizationConformancePackRequest", string(data)}, " ")
}
