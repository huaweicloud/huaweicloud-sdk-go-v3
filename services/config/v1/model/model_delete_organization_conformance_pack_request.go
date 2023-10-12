package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteOrganizationConformancePackRequest Request Object
type DeleteOrganizationConformancePackRequest struct {

	// 组织ID。
	OrganizationId string `json:"organization_id"`

	// 合规规则包ID。
	ConformancePackId string `json:"conformance_pack_id"`
}

func (o DeleteOrganizationConformancePackRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteOrganizationConformancePackRequest struct{}"
	}

	return strings.Join([]string{"DeleteOrganizationConformancePackRequest", string(data)}, " ")
}
