package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateOrganizationConformancePackRequest Request Object
type UpdateOrganizationConformancePackRequest struct {

	// 组织ID。
	OrganizationId string `json:"organization_id"`

	// 合规规则包ID。
	ConformancePackId string `json:"conformance_pack_id"`

	Body *UpdateOrgConformancePackRequestBody `json:"body,omitempty"`
}

func (o UpdateOrganizationConformancePackRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOrganizationConformancePackRequest struct{}"
	}

	return strings.Join([]string{"UpdateOrganizationConformancePackRequest", string(data)}, " ")
}
